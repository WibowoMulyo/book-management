# 📚 Book Management API

Book Management API adalah REST API yang dibangun menggunakan **Go (Gin Framework)**, **PostgreSQL**, dan **JWT Authentication** untuk mengelola data buku dan kategori.

---

## 🚀 Fitur

- 🔑 Autentikasi berbasis JWT
- 📚 CRUD Buku
- 📂 CRUD Kategori
- 🔗 Relasi Buku–Kategori
- 📏 Perhitungan otomatis ketebalan buku (`tipis/tebal`)
- ✅ Validasi input dengan aturan bisnis
- 🗃️ Database migration & seeding

---

## 🛠️ Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin Gonic
- **Database**: PostgreSQL 12+
- **Authentication**: JWT
- **Migration**: sql-migrate
- **Validation**: go-playground/validator
- **Password Hashing**: bcrypt

---

## ⚙️ Setup Project

### 1. Clone Repository
```bash
git clone <your-repository-url>
cd book-management
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Setup Database

**Dengan Docker:**
```bash
docker-compose up -d
```

**Atau manual:**
```sql
CREATE DATABASE book_management;
```

### 4. Konfigurasi Environment
Copy `.env.example` ke `.env` lalu sesuaikan:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=book_management
DB_SSLMODE=disable

JWT_SECRET=your_super_secret_jwt_key_here
JWT_EXPIRE_HOURS=24

PORT=8080
```

### 5. Jalankan Aplikasi
```bash
go run cmd/main.go
```

Server akan berjalan di:
👉 **http://localhost:8080**

---

## 🔑 Default Login

```
Username: admin
Password: admin123
```

---

## 📖 API Documentation

### Base URL
```
http://localhost:8080/api
```

### Authentication
Semua endpoint (kecuali login) memerlukan JWT token:

```
Authorization: Bearer <token>
```

---

## 📋 Endpoints

### 🔐 Authentication
- `POST /users/login` → login & dapatkan JWT token

### 📂 Categories
- `GET /categories` → semua kategori
- `GET /categories/{id}` → detail kategori
- `POST /categories` → tambah kategori
- `PUT /categories/{id}` → update kategori
- `DELETE /categories/{id}` → hapus kategori
- `GET /categories/{id}/books` → daftar buku dalam kategori

### 📚 Books
- `GET /books` → semua buku
- `GET /books/{id}` → detail buku
- `POST /books` → tambah buku
- `PUT /books/{id}` → update buku
- `DELETE /books/{id}` → hapus buku

---

## 📝 Request & Response Examples

### Login
**Request:**
```bash
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

**Response:**
```json
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_at": "2024-01-01T12:00:00Z"
  }
}
```

### Create Book
**Request:**
```bash
curl -X POST http://localhost:8080/api/books \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "title": "Belajar Go Programming",
    "author": "John Doe",
    "pages": 250,
    "category_id": 1
  }'
```

**Response:**
```json
{
  "status": "success",
  "message": "Book created successfully",
  "data": {
    "id": 1,
    "title": "Belajar Go Programming",
    "author": "John Doe",
    "pages": 250,
    "thickness": "tebal",
    "category_id": 1,
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
}
```

---

## 🗂️ Project Structure

```
book-management/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── database.go
│   ├── handlers/
│   │   ├── auth.go
│   │   ├── books.go
│   │   └── categories.go
│   ├── middleware/
│   │   └── jwt.go
│   ├── models/
│   │   ├── book.go
│   │   ├── category.go
│   │   └── user.go
│   └── routes/
│       └── routes.go
├── migrations/
│   ├── 001_create_users_table.sql
│   ├── 002_create_categories_table.sql
│   └── 003_create_books_table.sql
├── docker-compose.yml
├── .env.example
├── go.mod
├── go.sum
└── README.md
```

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👨‍💻 Author

**Wibowo Mulyo**
- GitHub: [@WibowoMulyo](https://github.com/WibowoMulyo)
