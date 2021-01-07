package main

import (
   "io"
   "net/http"
   "os"
)

func HttpCopy(source, dest string) (int64, error) {
   get_o, e := http.Get(source)
   if e != nil {
      return 0, e
   }
   create_o, e := os.Create(dest)
   if e != nil {
      return 0, e
   }
   return io.Copy(create_o, get_o.Body)
}

func main() {
   HttpCopy("http://speedtest.lax.hivelocity.net", "index.html")
}
