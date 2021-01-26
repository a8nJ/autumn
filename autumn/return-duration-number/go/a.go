package main

import (
   "fmt"
   "log"
)

func main() {
   var tests = []string {
      "2020",
      "2020-12",
      "2020-12-31",
      "2020-12-31T01",
      "2020-12-31T01:02",
      "2020-12-31T01:02:31",
   }
   for _, test := range tests {
      n, e := sinceHours(test)
      if e != nil {
         log.Fatal(e)
      }
      fmt.Println(n)
   }
}
