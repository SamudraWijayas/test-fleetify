# Procurement Backend

Backend API untuk sistem procurement yang dibangun dengan Go, Fiber, dan GORM.

## Deskripsi

Aplikasi backend ini menyediakan API untuk sistem procurement yang mencakup manajemen user, supplier, item, dan pembelian. Menggunakan autentikasi JWT dan database MySQL.

## Prerequisites

- Go 1.25.6 atau lebih baru
- MySQL Server

## Instalasi

1. Clone repository ini.
2. Masuk ke direktori backend:
   ```
   cd backend
   ```
3. Download dependencies:
   ```
   go mod download
   ```

## Setup

1. Buat file `.env` di direktori backend dengan konfigurasi database:

   ```
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=procurement
   JWT_SECRET=your_jwt_secret
   ```

2. Jalankan script database:
   ```
   mysql -u your_username -p < db.sql
   ```

## Menjalankan Aplikasi

Jalankan aplikasi dengan perintah:

```
go run main.go
```

Server akan berjalan di `http://localhost:3000`.

## API Endpoints

### Autentikasi

- `POST /register` - Registrasi user baru
- `POST /login` - Login user

User default untuk testing:
Username: carrick
Password: 123456

### Items (Memerlukan autentikasi)

- `GET /items` - Mendapatkan daftar item
- `POST /items` - Membuat item baru

### Suppliers (Memerlukan autentikasi)

- `GET /suppliers` - Mendapatkan daftar supplier
- `POST /suppliers` - Membuat supplier baru

### Purchases (Memerlukan autentikasi)

- `POST /purchases` - Membuat pembelian baru

## Teknologi yang Digunakan

- Go
- Fiber (web framework)
- GORM (ORM)
- MySQL
- JWT (autentikasi)
