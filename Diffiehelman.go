package main
import (
    "fmt"
    "math/rand"
    "time"
)

var generator int
var prime int
var key_length int

generator := 2
key_length := 600
prime := 23

// func power(a ,b , p) int{
//   if b == 1 {
//     return a
//   } else {
//     return math.Pow(a,b)% p
//   }
//
// }

rand.Seed(time.Now().UnixNano())
func generate_private_key (length int) int{
  // _rand := 0
  _bytes := length // 8 + 8
  private_key := int64(rand.Intn(_bytes))
  return private_key
}

func generate_public_key (private_key int) int{
  public_key := math.Pow(generator, private_key, prime)
  return public_key
}

func generate_secret_key (private_key int, public_key int) int{
  secret_key := math.Pow(int64(public_key), int64(private_key), prime)
  return secret_key
}
