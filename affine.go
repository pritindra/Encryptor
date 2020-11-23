package main

func affineEncrypt(string msg) string {
  cipher := ""
  var i int
  for i := 0; i < len(msg); i++ {
    if msg[i] != '' {
      cipher := cipher + char((((a*(msg[i]-'A'))+b)%26) + 'A')

    } else {
      cipher := cipher + msg[i]

    }
  }
  return cipher
}
