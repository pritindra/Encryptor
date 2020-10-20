package main

arr1 := [26]string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
var arr2 int

// for i := 0; i < 26; i++ {
// 	arr2[i] = i
// }

func VernamCipher(string str, string key) string {

	var vernamcipher string
	str_no := []int{}
	key_no := []int{}
	sum_no := []int{}
	key_len := len(key)
	str_len := len(str)

	for j := 0; j < key_len; j++ {
		for k := 0; k < 26; k++ {
			if key[j] == arr1[k] {
				key_no[j] := k
			}
		}
	}

	for j := 0; j < str_len; j++ {
		for k := 0; k < 26; k++ {
			if str[j] == arr1[k] {
				str_no[j] := k
			}
		}
	}

	for j := 0; j < str_len; j++ {
		sum_no[j] := str_no[j] + key_no[j]

		if sum_no[j] > 26 {
			sum_no[j] := sum_no[j] - 26
		}
	}

	for j := 0; j < str_len; j++ {
		vernamcipher[j] := arr1[sum_no[j]]
		vernamcipher = append(vernamcipher, vernamcipher[j])
	}

	return vernamcipher
}
