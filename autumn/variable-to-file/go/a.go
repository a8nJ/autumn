package main

import (
   "fmt"
   "os"
)

func main() {
   s := "March"
   o, e := os.Create("a.txt")
   if e != nil {
      os.Exit(1)
   }
   fmt.Fprintln(o, s)
}
