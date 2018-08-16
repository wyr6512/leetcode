func reverse(x int) int {
    y := x;
    if x < 0 {
        y = 0-x;
    }
    s := f(y)
    num, _ := strconv.Atoi(s)
    if float64(num) > math.Pow(2, 31) - 1 {
        return 0
    }
    if x < 0{
        return 0 - num
    }
    return num
}


func f(x int) string{
    n := x / 10
    i := x % 10    
    if n == 0{
        return strconv.Itoa(i)
    }
    return strconv.Itoa(i) + f(n)
}