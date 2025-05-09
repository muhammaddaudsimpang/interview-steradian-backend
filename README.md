# Jalankan SQL untuk membuat database dan tabel
psql -U postgres -f sql/initSQL.sql

# Buat file environment
.env

# Inisialisasi modul Go
go mod init github.com/username/car-rental-backend

# Install dependencies
go mod tidy

# Jalankan aplikasi
go run .
