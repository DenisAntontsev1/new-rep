package main

//import "fmt"
import "math"

// Наш тип для комплексных чисел - структура из
// действительной и мнимой части
type cmplx_num struct {
    real float64
    img float64
}

// Сложение
func cmplx_add(n1, n2 cmplx_num) (cmplx_num) {
    n3 := cmplx_num{n1.real + n2.real, n1.img + n2.img}
    return n3
}

// Вычитание
func cmplx_sub(n1, n2 cmplx_num) (cmplx_num) {
    n3 := cmplx_num{n1.real - n2.real, n1.img - n2.img}
    return n3
}

// Умножение
func cmplx_mul(n1, n2 cmplx_num) (cmplx_num) {
    n3 := cmplx_num{n1.real * n2.real - n1.img * n2.img, n1.real*n2.img + n1.img*n2.real}
    return n3
}

// Деление
func cmplx_div(n1, n2 cmplx_num) (cmplx_num) {
    new_real := (n1.real*n2.real + n1.img*n2.img) / (n2.real*n2.real + n2.img*n2.img)
    new_img  := (n1.img*n2.real - n1.real*n2.img) / (n2.real*n2.real + n2.img*n2.img)
    n3 := cmplx_num{new_real, new_img}
    return n3
}

// Abs (модуль комплексного числа)
func cmplx_abs(n1 cmplx_num) (float64) {
    return math.Sqrt(n1.real*n1.real + n1.img*n1.img)
}

// Возведение в степень
func cmplx_pow(n1 cmplx_num, p float64) (cmplx_num) {
    r, phi := cmplx_toPolar(n1)
    return cmplx_num{math.Pow(r, p)*math.Cos(p*phi), math.Pow(r, p)*math.Sin(p*phi)}
}

// Sqrt (корень квадратный)
func cmplx_sqrt(n1 cmplx_num) (cmplx_num) {
    if (n1.img == 0) {
        return cmplx_num{math.Sqrt(n1.real), 0}
    }

    z := cmplx_abs(n1)
    new_real := math.Sqrt((z + n1.real)/2)
    new_img  := n1.img/math.Abs(n1.img)*math.Sqrt((z - n1.real)/2)
    return cmplx_num{new_real, new_img}
}

// Перевод в полярные координаты
func cmplx_toPolar(n1 cmplx_num) (float64, float64) {
    r := cmplx_abs(n1)
    phi := math.Acos(n1.real/r)
    return r, phi
}


/*
func main() {
    n1 := cmplx_num{1, 2}
    n2 := cmplx_num{3, 4}

    res := cmplx_add(n1, n2)
    fmt.Println("n1 / n2 = ", res)
}
*/
