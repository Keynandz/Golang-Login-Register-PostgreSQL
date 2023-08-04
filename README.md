# Golang Login Register PostgreSQL
## Indonesia:
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

Semoga ini dapat membantu dan bermanfaat. Terimakasih

## English
### 1. Setup PostgreSQL Database
Run the command on your PostgreSQL

**create school database;**

After that run the command

**\c school;**

If PostgreSQL has changed to "school=#". Then next run the command

**create account table (**

**id SERIAL PRIMARY KEY,**

**name VARCHAR(250) NOT NULL,**

**email VARCHAR(250) NOT NULL,**

**password VARCHAR(250) NOT NULL**

**);**

Database Setup Complete

### 2. Customize the .env file with your database
Open the .env file and change it

EXAMPLE:

**DB_HOST=localhost**

**DB_PORT=5432**

**DB_USER=postgres**

**DB_PASSWORD=123**

**DB_NAME=school**

### 3. Run the program
After everything is set correctly, run the "go run ." or "go run main.go" command.

To see how it looks on the web, go to http://localhost:5000/login

Hopefully this can help and be useful. Thank you
