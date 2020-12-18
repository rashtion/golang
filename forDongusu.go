package main

import (
    "fmt"
)

func main() {
    deger := 1
    for deger <= 10 {
        fmt.Println(deger) // Değişkeni yazdır ama:
        deger = deger + 1 // Değişken aynı ama 10'a kadar değişken'e 1 ekle ve yazdır.
    // Çıktı: 1 2 3 4 5 6 7 8 9 10
    
    // Tüm kodlar tarafımca öğrenirken yazılmıştır

    }
}
