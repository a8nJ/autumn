package main

import (
   "log"
   "net/url"
)

func main() {
   u, e := url.Parse("https://example.com/one?two=even")
   if e != nil {
      log.Fatal(e)
   }
   v := u.Query()
   log.Print(v)
}
