# GoStore

GoStore adalah aplikasi e-commerce yang saya bangun sebagai proyek pribadi. Aplikasi ini menggunakan berbagai teknologi modern untuk menyediakan fitur-fitur e-commerce yang handal dan aman.

## Tech Stack

- **Golang**: Bahasa pemrograman yang digunakan untuk membangun aplikasi ini.
- **GORM**: ORM untuk Golang yang digunakan untuk berinteraksi dengan database.
- **GorillaMux**: Router HTTP yang sangat fleksibel untuk routing.
- **JWT (JSON Web Token)**: Digunakan untuk otentikasi dan otorisasi.
- **Session**: Mengelola sesi pengguna untuk pengalaman pengguna yang lebih baik.
- **Bcrypt**: Algoritma untuk enkripsi password, memastikan keamanan data pengguna.
- **MySQL**: Database relasional yang digunakan untuk menyimpan semua data aplikasi.

## Fitur

- **Manajemen Produk**: Tambahkan, edit, dan hapus produk.
- **Keranjang Belanja**: Tambahkan produk ke keranjang belanja.
- **Checkout**: Proses checkout yang mencakup pembuatan pesanan.
- **Otentikasi**: Registrasi dan login pengguna dengan enkripsi password menggunakan bcrypt.
- **Otorisasi**: Menggunakan JWT untuk mengamankan endpoint API.
- **Manajemen Sesi**: Mengelola sesi pengguna untuk menjaga pengalaman pengguna.
- **Swagger Documentation**: Dokumentasi API interaktif yang dapat diakses di [Swagger UI](http://localhost:8000/api/swagger).

## Instalasi

1. Clone repository ini:

   ```sh
   git clone https://github.com/hiboedi/go-store-backend.git
   cd go-store-backend
   ```

2. Instal dependencies:

   ```sh
   go mod download
   ```

3. Buat file `.env` untuk konfigurasi lingkungan:

   ```env
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_NAME=gostore
   JWT_SECRET=your_jwt_secret
   ```

4. Migrasi database:

   ```sh
   go run main.go
   ```

5. Jalankan aplikasi:
   ```sh
   go run main.go
   ```

## API Endpoint

### Otentikasi

- **Register**: `POST /api/signup`
- **Login**: `POST /api/login`

### Produk

- **Tambah Produk**: `POST /api/products`
- **Lihat Produk**: `GET /api/products`
- **Edit Produk**: `PUT /api/products/:id`
- **Hapus Produk**: `DELETE /api/products/:id`

### Keranjang Belanja

- **Tambah ke Keranjang**: `POST /api/cart`
- **Lihat Keranjang**: `GET /api/cart`

### Checkout

- **Proses Checkout**: `POST /api/checkout`

## Contributing

Jika Anda ingin berkontribusi pada proyek ini, silakan fork repository ini dan buat pull request dengan perubahan Anda.

## Lisensi

Proyek ini dilisensikan di bawah lisensi MIT. Lihat file [LICENSE](LICENSE) untuk informasi lebih lanjut.
