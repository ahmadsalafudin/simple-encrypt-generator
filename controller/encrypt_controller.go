package controller

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "github.com/ahmadsalafudin/simple-encrypt-generator/service"
    "github.com/ahmadsalafudin/simple-encrypt-generator/config"
)

func RunConsoleInterface() {
    scanner := bufio.NewScanner(os.Stdin)

    // Display menu
    fmt.Println("Pilih operasi yang ingin dilakukan:")
    fmt.Println("1. Encrypt")
    fmt.Println("2. Decrypt")
    fmt.Print("Masukkan pilihan Anda: ")

    // Read user input
    scanner.Scan()
    choice := scanner.Text()

    switch choice {
    case "1":
        fmt.Print("Masukkan teks yang ingin di-encrypt: ")
        scanner.Scan()
        plainText := scanner.Text()

        fmt.Println("Pilih metode enkripsi:")
        fmt.Println("1. Base64")
        fmt.Println("2. Hash")
        fmt.Println("3. AES")
        fmt.Println("4. RSA")
        fmt.Println("5. DES")
        fmt.Println("6. Blowfish")
        fmt.Print("Masukkan pilihan metode enkripsi: ")
        scanner.Scan()
        method := getEncryptionMethod(scanner.Text())

        encryptedText := service.Encrypt(plainText, method)
        fmt.Printf("Teks terenkripsi: %s\n", encryptedText)

    case "2":
        fmt.Print("Masukkan teks yang ingin di-decrypt (untuk hash, masukkan hash yang ingin dicocokkan): ")
        scanner.Scan()
        cipherText := scanner.Text()

        fmt.Println("Pilih metode dekripsi:")
        fmt.Println("1. Base64")
        fmt.Println("2. Hash")
        fmt.Println("3. AES")
        fmt.Println("4. RSA")
        fmt.Println("5. DES")
        fmt.Println("6. Blowfish")
        fmt.Print("Masukkan pilihan metode dekripsi: ")
        scanner.Scan()
        method := getEncryptionMethod(scanner.Text())

        if method == config.Hash {
            fmt.Print("Masukkan teks asli untuk dibandingkan dengan hash: ")
            scanner.Scan()
            originalText := scanner.Text()
            result := service.Decrypt(cipherText, method, originalText)
            fmt.Printf("Hasil pencocokan hash: %s\n", result)
        } else {
            decryptedText := service.Decrypt(cipherText, method, "")
            fmt.Printf("Teks terdekripsi: %s\n", decryptedText)
        }

    default:
        fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
    }
}

// getEncryptionMethod maps the user input to the corresponding encryption method.
func getEncryptionMethod(choice string) string {
    switch strings.TrimSpace(choice) {
    case "1":
        return config.Base64
    case "2":
        return config.Hash
    case "3":
        return config.AES
    case "4":
        return config.RSA
    case "5":
        return config.DES
    case "6":
        return config.Blowfish
    default:
        return ""
    }
}
