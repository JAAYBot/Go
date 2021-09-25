sum := 0
for i := 1; i < 5; i++ {
    sum += i
}
fmt.Println(sum)


n := 1
for n < 5 {
    n *= 2
}
fmt.Println(n)


strings := []string{"hello", "world"}
for i, s := range strings {
    fmt.Println(i, s)
}

var test = "hello" string
var test2 = 1 int
