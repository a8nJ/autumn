package main

import (
   "encoding/json"
   "strings"
)

func main() {
   m := Map{}
   o := strings.NewReader(in_s)
   json.NewDecoder(o).Decode(&m)
   out_s := m.M("U2").A("Boy")[0]
   println(out_s == "Twilight")
}
