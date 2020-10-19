package main

// function to generate key in circular way
func generateKey(string str, string key) string {
	x := len(str)
	y := len(key)

	for i := 0; ; i++ {
		if x == i {
			i := 0
		} else if x == y {
			break
		}
		key = append(key, key[i])
	}
	return key
}

//  function to return the encrypted text
func VigenereCipher(string str, string key) string {
	var vigenere_cipher string
	x := len(str)

	for i := 0; i < x; i++ {
		y := (int(str[i]) + int(key[i])) % 26
		y := y + int('A')
		vigenere_cipher = append(vigenere_cipher, x)
	}
	return vigenere_cipher
}

func VigenereDeCipher(string vigenere_cipher, string key) {
	var orig_text string
	x := len(vigenere_cipher)

	for i := 0; i < x; i++ {
		y := (int(vigenere_cipher[i]) - int(key[i]) + 26) % 26
		y := y + int('A')
		orig_text = append(orig_text, y)
	}
	return orig_text
}
