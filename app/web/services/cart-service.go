package services

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hiboedi/go-store-backend/app/exceptions"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/repositories"
	"gorm.io/gorm"
)

type CartService interface {
	AddToCart(ctx context.Context, request models.CartItemCreate, userId string) models.CartItemResponse
	CheckoutCart(ctx context.Context, request models.OrderCreate, userID string) models.OrderResponse
}

type cartServiceImpl struct {
	cartRepo    repositories.CartRepository
	orderRepo   repositories.OrderRepository
	productRepo repositories.ProductRepository
	userRepo    repositories.UserRepository
	DB          *gorm.DB
	Validate    *validator.Validate
}

func NewCartService(cartRepo repositories.CartRepository, orderRepo repositories.OrderRepository, productRepo repositories.ProductRepository, userRepo repositories.UserRepository, db *gorm.DB, validate *validator.Validate) CartService {
	return &cartServiceImpl{
		cartRepo:    cartRepo,
		orderRepo:   orderRepo,
		productRepo: productRepo,
		userRepo:    userRepo,
		DB:          db,
		Validate:    validate,
	}
}

func (s *cartServiceImpl) AddToCart(ctx context.Context, request models.CartItemCreate, userID string) models.CartItemResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	cart, err := s.cartRepo.GetCartByUserID(ctx, tx, userID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	product, err := s.productRepo.GetProductById(ctx, tx, request.ProductID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	if request.Quantity > uint32(product.Stock) {
		log.Fatal("Sorry product is out of stock")
	}

	basePrice := helpers.BasePrice(product.Price, request.Quantity)
	discountAmount := helpers.CalculateDiscount(basePrice, product.DiscountPercent)
	priceAfterDiscount := helpers.PriceAfterDiscount(basePrice, discountAmount)
	taxAmount := helpers.TaxAmount(priceAfterDiscount)
	totalPrice := helpers.TotalPrice(basePrice, discountAmount, taxAmount)

	cartItemId := uuid.New().String()
	cartItem := models.CartItem{
		ID:              cartItemId,
		CartID:          cart.ID,
		ProductID:       product.ID,
		Quantity:        request.Quantity,
		BasePrice:       basePrice,
		DiscountPercent: product.DiscountPercent,
		DiscountAmount:  discountAmount,
		TaxPercent:      helpers.TaxPercent,
		TaxAmount:       taxAmount,
		TotalPrice:      totalPrice,
	}

	data, err := s.cartRepo.AddCartItem(ctx, tx, cartItem)
	helpers.PanicIfError(err)

	product.Stock = product.Stock - int64(request.Quantity)
	s.productRepo.UpdateProduct(ctx, tx, product)

	cartItems := []models.CartItem{
		cartItem,
	}
	cart.CartItems = cartItems
	s.cartRepo.UpdateCart(ctx, tx, cart)

	return models.ToCartItemResponse(data)
}

func (s *cartServiceImpl) CheckoutCart(ctx context.Context, request models.OrderCreate, userID string) models.OrderResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	user, err := s.userRepo.GetUserById(ctx, tx, userID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	cart, err := s.cartRepo.GetCartByUserID(ctx, tx, userID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	order := models.Order{
		ID:           uuid.New().String(),
		UserID:       userID,
		StoreID:      request.StoreID,
		IsPaid:       false,
		CustomerName: user.Name,
		Phone:        user.Phone,
		TotalPrice:   cart.CartItems[0].TotalPrice,
		Address:      request.Address,
	}
	order, err = s.orderRepo.CreateOrder(ctx, tx, order)
	helpers.PanicIfError(err)

	var orderItems []models.OrderItem
	for _, cartItem := range cart.CartItems {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: cartItem.ProductID,
			Quantity:  int(cartItem.Quantity),
		}
		orderItem, _ = s.orderRepo.CreateOrderItem(ctx, tx, orderItem)
		orderItems = append(orderItems, orderItem)
	}

	order.OrderItems = orderItems
	order, err = s.orderRepo.UpdateOrder(ctx, tx, order)
	helpers.PanicIfError(err)

	// if err := tx.Delete(&cart).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// if err := tx.Delete(&CartItem{}, "cart_id = ?", cart.ID).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	return models.ToOrderResponse(order)
}
