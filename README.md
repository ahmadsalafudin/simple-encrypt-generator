# simple-encrypt-generator

Layanan ini merupakan aplikasi konsol yang menyediakan fitur enkripsi dan dekripsi menggunakan beberapa metode, termasuk Base64, Hash, dan RSA. Pengguna dapat memilih metode yang diinginkan untuk mengenkripsi atau mendekripsi teks.

## Fitur

- Enkripsi teks menggunakan metode:
    - Base64
    - Hash
    - AES
    - RSA
    - DES
    - Blowfish
- Dekripsi teks dengan metode yang sama
- Pencocokan hash untuk memverifikasi apakah teks asli sesuai dengan hash yang disimpan

## Struktur Proyek

```
simple-encrypt-generator/ 
├── config/ 
│ └── config.go 
├── controller/ 
│ └── encrypt_controller.go 
└── service/ 
└── encrypt_service.go
server.go
```

## Instalasi

1. **Clone repositori**:
   ```bash
   git clone https://github.com/ahmadsalafudin/simple-encrypt-generator.git
   cd simple-encrypt-generator
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the application**:
   ```bash
   go run server.go
   ```

## Contoh Penggunaan

**1. Encrypt**
  Pilih operasi yang ingin dilakukan:
  - (1). Encrypt
  - (2). Decrypt
  Masukkan pilihan Anda: 1
  Masukkan teks yang ingin di-encrypt: Hello World
  Pilih metode enkripsi:
    - (1). Base64
    - (2). Hash
    - (3). RSA
    - (4). AES
    - (5). DES
    - (6). Blowfish
  Masukkan pilihan metode enkripsi: 1
  Teks terenkripsi: Basic SGVsbG8gV29ybGQ=

**2. Decrypt**
Pilih operasi yang ingin dilakukan:
  - (1). Encrypt
  - (2). Decrypt
Masukkan pilihan Anda: 2
Masukkan teks yang ingin di-decrypt: Basic SGVsbG8gV29ybGQ=
Pilih metode dekripsi:
  - (1). Base64
  - (2). Hash
  - (3). RSA
  - (4). AES
  - (5). DES
  - (6). Blowfish
Masukkan pilihan metode dekripsi: 1
Teks terdekripsi: Hello World

    