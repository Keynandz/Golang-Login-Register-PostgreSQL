# Golang Login Register PostgreSQL
### 1. Setup Database PostgreSQL
Jalankan perintah pada PostgreSQL kalian

**create database sekolah;**

Setelah itu jalankan perinah

**\c sekolah;**

Jika PostgreSQL sudah berubah menjadi "sekolah=#". Maka selanjutnya jalankan perinah

**create table akun (**

**id SERIAL PRIMARY KEY,**

**nama VARCHAR(250) NOT NULL,**

**email VARCHAR(250) NOT NULL,**

**password VARCHAR(250) NOT NULL**

**);**

Setup Database Selesai

### 2. Sesuaikan file .env dengan database kalian
Buka file .env lalu ubah

CONTOH:

**DB_HOST=localhost**

**DB_PORT=5432**

**DB_USER=postgres**

**DB_PASSWORD=123**

**DB_NAME=sekolah**

### 3. Jalankan program
Setelah semua ter setting dengan benar, jalankan perinah "go run ." atau "go run main.go"

Untuk menglihat tampilannya di web, buka http://localhost:5000/login

