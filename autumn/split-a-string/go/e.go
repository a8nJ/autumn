package main

import (
   "strings"
)

func main() {
   s := strings.SplitN("west,north,east", ",", 2)
   fmt.Printf("%q\n", s) // ["west" "north,east"]
}
