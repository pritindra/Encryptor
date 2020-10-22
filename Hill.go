package main

func getKeyMatrix(string key, int keyMatrix[][3]) {
	
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			keyMatrix[i][j] = key([k]) % 65
			k++
		}
	}
}

func HillEncrypt(int cipherMatrix[][1], int keyMatrix[][3], int messageVector[][1]) {

	var x,i,j int
	for i :=0; i < 3; i++ {
		for j := 0; j < 1; j++ {
			cipherMatrix[i][j] := 0
			
			for x := 0; x < 3; x++ {
				cipherMatrix[i][j] := cipherMatrix[i][j] + keyMatrix[i][x] * messageVector[x][j]
			}
			cipherMatrix[i][j] := cipherMatrix[i][j] % 26
		}
	}
}

func HillCipher(string message, string key) {
	var keyMatrix[3][3] int
	getKeyMatrix(key, keyMatrix)

	var messageVector[3][1]
	for i := 0; i < 3; i++ {
		messageVector[i][0] := (message[i]) % 65
	}

	var cipherMatrix[3][1] int 
	HillEncrypt(cipherMatrix, keyMatrix, messageVector)

	var CipherText string
	for i := 0; i < 3; i++ {
		CipherText := cipherMatrix[i][0] + 65;
	}
  fmt.Printf("Cipher Text: %s", CipherText)
}
}
