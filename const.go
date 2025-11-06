/*
Qaysi biri yaxshiroq?
Bu vazifaga bog'liq:
✅ const ishlatish kerak:

Qiymat hech qachon o'zgarmasa (API URL, limit qiymatlari)
Konfiguratsiya konstantalari
Magic numberlarni nomlab qo'yish uchun

✅ var ishlatish kerak:

Qiymat dastur davomida o'zgarishi kerak
Complex tiplarda (struct, array, map)
Runtime'da hisoblangan qiymatlar

Umumiy qoida: 
Agar qiymat o'zgarmaydigan bo'lsa va 
compile-time'da ma'lum bo'lsa - const ishlatish yaxshiroq. 
Bu kod xavfsizligini oshiradi va xatolarni kamaytiradi.
*/
package main

import "fmt"

func main(){
	// Yaxshi amaliyot
const (
    StatusActive   = 1
    StatusInactive = 0
    MaxRetries     = 3
)

var (
    currentUser User
    counter int
)
}