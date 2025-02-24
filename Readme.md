# Temperature Converter Module

## 📌 Deskripsi

Go Temperature Converter adalah module Go sederhana untuk melakukan konversi suhu antara Celsius, Fahrenheit, dan Kelvin.

## 🚀 Instalasi

Untuk menggunakan module ini dalam proyek Go, jalankan perintah berikut:

```sh
go get github.com/ren-zi-fa/temperature-converter
```

## 📚 Penggunaan

Import package ke dalam kode Go kamu:

```go
package main

import (
    "fmt"
    "github.com/username/go-temp-converter/temperature"
)

func main() {
    fmt.Println("Celsius ke Fahrenheit:", temperature.CelsiusToFahrenheit(100)) // Output: 212
    fmt.Println("Fahrenheit ke Celsius:", temperature.FahrenheitToCelsius(32))  // Output: 0
    fmt.Println("Celsius ke Kelvin:", temperature.CelsiusToKelvin(25))          // Output: 298.15
    fmt.Println("Kelvin ke Celsius:", temperature.KelvinToCelsius(273.15))      // Output: 0
}
```

## 🔥 Fungsi yang Tersedia

| Fungsi                                   | Deskripsi                           |
| ---------------------------------------- | ----------------------------------- |
| `CelsiusToFahrenheit(c float64) float64` | Konversi dari Celsius ke Fahrenheit |
| `FahrenheitToCelsius(f float64) float64` | Konversi dari Fahrenheit ke Celsius |
| `CelsiusToKelvin(c float64) float64`     | Konversi dari Celsius ke Kelvin     |
| `KelvinToCelsius(k float64) float64`     | Konversi dari Kelvin ke Celsius     |




