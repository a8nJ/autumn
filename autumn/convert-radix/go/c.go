package main
import "strconv"

func Encode36(x int) string {
   n := int64(x)
   return strconv.FormatInt(n, 36)
}

func main() {
   n := 1577858399
   s := Encode36(n)
   println(s == "q3ezbz")
}
