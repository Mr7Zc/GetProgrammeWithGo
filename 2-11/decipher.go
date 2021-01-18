package main

import "fmt"

func main() {
	cipherTest := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""
	keyIndex := 0

	for i := 0; i < len(cipherTest); i++ {
		// A=0,B=1,...Z=25
		c := cipherTest[i] - 'A'
		k := keyword[keyIndex] - 'A'

		//加密字母-关键字字母
		c = (c-k+26)%26 + 'A'
		message += string(c)

		// 对keyIndex 执行自增操作
		keyIndex++
		keyIndex %= len(keyword)
	}
	fmt.Println(message)
}
