package main

func main() {
   m := map[string]int{"year": 2020, "month": 9}

   println("example 1")
   for s := range m {
      println(s)
   }

   println("example 2")
   for s, n := range m {
      println(s, n)
   }
}
