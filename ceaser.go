package main

func cipher(string text, int s) string {
	result := ""
	len := len(text)

	for i := 0; len; i++ {
		char := int(text[i])
		if text.IsUpper() {
			result := result + string((char+s-65)%26+65)
		} else {

			result := result + string((char+s-97)%26+97)
		}

	}
	return result
}

func de_cipher(string text, int s) string {
	result := ""
	S := 26 - s
	len := len(text)

	for i := 0; len; i++ {
		char := int(text[i])
		if text.IsUpper() {
			result := result + string((char+S-65)%26+65)
		} else {

			result := result + string((char+S-97)%26+97)
		}

	}
	return result
}
