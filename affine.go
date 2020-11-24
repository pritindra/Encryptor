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

func affineDecrpyt(string cipher) string {
  msg := ""
  a_inv := 0
  flag := 0
  var i int
  for i:=0; i<26; i++ {
    flag := (a*i) % 26
    if flag = 1 {
      a_inv = i
    }
  }
  for i:=0; i < len(cipher); i++  {
    if cipher[i] != ' '{
      msg := msg + char(((a_inv*((cipher[i]+'A' - b))%26)) + 'A')
    }else {
      msg := msg + cipher[i]
    }
    
    
  }
  return msg
}
