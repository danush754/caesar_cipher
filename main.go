package main

import (
	"fmt"
	"strings"
)

func main() {

	caesar_cipher("Dhanusz", 5)

}

func caesar_cipher(cipher string, shiftVal int) {

	var maxlimit = 122
	var finalCipher string

	var toLowerCasedCipher = strings.ToLower(cipher)

	for _, cipherwords := range toLowerCasedCipher {

		convertedCipher := cipherwords + rune(shiftVal)

		if convertedCipher <= rune(maxlimit) {
			finalCipher += string(convertedCipher)
		} else {
			convertedCipher = convertedCipher - rune(26)

			finalCipher += string(convertedCipher)
		}
	}

	fmt.Println("Secret Key :==", finalCipher)
}
