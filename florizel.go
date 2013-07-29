package main

import (
  "github.com/NovemberFoxtrot/tango"
  "fmt"
  "flag"
)

func main() {
  var ip = flag.String("http", "000", "A valid HTTP code")

  flag.Parse()

  fmt.Println(tango.FetchMeaning(*ip))
}
