package main

import (
   "fmt"
   "io"
   "net/http"
)

type progress struct {
   io.Reader
   total int
}

func (p *progress) Read(b []byte) (int, error) {
   n, e := p.Reader.Read(b)
   if e != nil {
      fmt.Println()
   } else {
      p.total += n
      fmt.Printf("\rRead %9v", p.total)
   }
   return n, e
}

func main() {
   get, e := http.Get("http://speedtest.lax.hivelocity.net/10Mio.dat")
   if e != nil {
      panic(e)
   }
   io.ReadAll(&progress{Reader: get.Body})
}
