package main
// example 1
func f(n, n2 int) bool {
   return n > n2
}
// example 2
var f2 = func (n, n2 int) bool {
   return n > n2
}
// print
func main() {
   b := f(9, 8)
   b2 := f2(9, 8)
   println(b, b2)
}
