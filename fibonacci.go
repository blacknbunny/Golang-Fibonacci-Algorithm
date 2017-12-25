package main


import (
  "fmt"
  "time"
)
var (
  g int = 0
  v int = 1
  l int
)

func fib() {
  fmt.Println(g, v)
  l = g + v
  g = v
  v = l
  fmt.Println(l)
  fmt.Println("")
}

func main() {
  for {
    fib()
    time.Sleep(time.Second * 1)
  }
}
