package main

import (
    "fmt"
)

func main() {
    var dizi[3]string

    dizi[0] = "C++"
    dizi[1] = "Python"
    dizi[2] = "GoLang" 

    fmt.Printf("Yaşınız nedir? : ")
    var yas int
    fmt.Scanln(&yas)
    fmt.Printf("Terminal için mi, masaüstü için mi kodlarsınız? : ")
    var termdesk string
    fmt.Scanln(&termdesk)
    // Verileri input edelim

    if yas <= 18 {
        fmt.Println("Siz ", dizi[1], " öğrenmelisiniz.")
    } 

    if termdesk == "terminal" {
        fmt.Println("Siz ", dizi[0], " veya ", dizi[1], " öğrenmelisiniz.")
    } 


    // Kod'da ufak hatalar olabilir, bunlar için özür dilerim. <3

}

