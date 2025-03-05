package main
import "fmt"

func kombinasi (n int, r int) int{
    var i, result, j, k, l, m, o int
    m = r
    o = 1
    i = n
    k = 1
    j = n-r
    l = 1
    for i >= 1 {
        l *= i
        i--
    }
    for j >= 1{
        k *= j
        j--
    }
    for m >= 1{
        o *=m
        m--
    }
    result = l/(k*o)
    return result
}
func main(){
    var d1, d2 int
    fmt.Scan(&d1, &d2)
    fmt.Print(kombinasi(d1, d2))
}

