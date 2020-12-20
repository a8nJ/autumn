package main

import (
   "log"
   "os"
)

func main() {
   o, e := os.Stat(`C:\Users`)
   if os.IsNotExist(e) || ! o.Mode().IsDir() {
      log.Fatal(e)
   }
}
