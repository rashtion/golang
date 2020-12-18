package main

import "fmt"

func main() {
    myarr := [5]string{"C++", "Python", "JavaScript", "PHP", "GoLang"}
    // Dizimdeki elemanlar:
    
    // Bu şekilde de yazılabilir.
 /* myarr[0] = "C++"
    myarr[1] = "Python"
    myarr[2] = "JavaScript"
    myarr[3] = "PHP"
    myarr[4] = "GoLang"
*/
    // Ekrana yazdıralım.
    fmt.Println("Benim, sırasıyla öğreneceğim programlama dilleri şunlardır:")
    fmt.Println(myarr)
    fmt.Println("Ama, öncelikle", myarr[2], "öğrenceğim.")

// Tüm kodlar tarafımca öğrenirken yazılmıştır

}
