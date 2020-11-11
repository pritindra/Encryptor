package main

import (
	"fmt"
	"math"
)

func encrpyt (string s) string {
	l := len(text)
	b := math.Ceil(math.Sqrt(l))
	a := math.Floor(math.Sqrt(l))
	encrpyted := ""
	if b*a < l {
		if math.Min(b,a) = b {
			b := b + 1
		} else {
			a := a + 1
		}
	}

	var arr []string
	var i,j int
	arr := [['' for i := range(a)] for j in range(b)]
	var k := 0

	for j := range a{
		for i := range b {
			if k<l {
				arr[j][i] = s[k]
			}
			k := k + 1
		}
	}

	for j := range b {
		for i in range a{
			encrpyted := encrpyted + arr[i][j]
		}
	}
	return encrpyted
}


func decrypt (string s) string {
	l := len(text)
	b := math.Ceil(math.Sqrt(l))
	a := math.Floor(math.Sqrt(l))
	decrypted := ""
	

	var arr []string
	var i,j int
	arr := [['' for i := range(a)] for j in range(b)]
	var k := 0

	for j := range a{
		for i := range b {
			if k<l {
				arr[j][i] = s[k]
			}
			k := k + 1
		}
	}

	for j := range b {
		for i in range a{
			decrypted := decrypted + arr[i][j]
		}
	}
	return decrypted
}


