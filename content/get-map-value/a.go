package main
func main() {
   m := map[string]int{"year": 2020}
   // example 1
   n, b := m["year"]
   // example 2
   n2 := m["year"]
   // print
   println(n, b, n2)
}
