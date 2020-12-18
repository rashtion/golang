package main

import (
    "fmt"
)

func main() {
    fmt.Printf("Isminizi giriniz: ")
    var isim string
    
    // Kullanıcıdan input alma kısmı
    fmt.Scanln(&isim)
    fmt.Printf("Soyadınızı giriniz: ")
    var soyisim string
    fmt.Scanln(&soyisim)

    // Gelen verileri ekrana yazdıralım.

    fmt.Println("Sizin tam isminiz:", isim + " " + soyisim)

    // Çıktı : Isminizi giriniz: Michael
    // Soyadınızı giriniz: Jackson
    // Sizin tam isminiz: Michael Jackson

    // Tüm kodlar tarafımca öğrenirken yazılmıştır

}
