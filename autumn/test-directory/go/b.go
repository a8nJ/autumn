package main

import (
   "log"
   "os"
)

func IsDir(s string) bool {
   o, e := os.Stat(s)
   return e == nil && o.IsDir()
}

func main() {
   b := IsDir(`C:\Users`)
   println(b)
}
