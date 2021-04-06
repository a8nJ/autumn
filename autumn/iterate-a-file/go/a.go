package main

import (
   "bufio"
   "os"
)

func main() {
   file, e := os.Open("a.go")
   if e != nil {
      panci(e)
   }
   defer file.Close()
   buf := bufio.NewScanner(file)
   for buf.Scan() {
      println(buf.Text())
   }
}
