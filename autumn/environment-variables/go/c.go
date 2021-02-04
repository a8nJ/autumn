package main

import (
   "log"
   "os"
)

func main() {
   e := os.Setenv("MSYSTEM", "MINGW64")
   if e != nil {
      log.Fatal(e)
   }
}
