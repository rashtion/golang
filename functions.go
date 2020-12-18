package main

import (
    "fmt"
)

func sayilar(a string, b string)string {
    return a + b
} // Burada belirlediğim fonksiyona eğer string gireceksem, tip olarak string, sayı(integer) gireceksem, tip olarak int yazıyorum.

func main() {
    oku := sayilar("Ahmet ", "Mehmet")
    fmt.Println(oku)
} // Oku adında belirlediğim değişkene, fonksiyonun adını belirtip parantez açtım ve içine gelecek değerleri girdim. Sonra ekrana değişkenin adını girerek yazdırdım
