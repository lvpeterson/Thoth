package hashes

import "thoth/internal/util"

var hashMap map[int]string

// Check formatting of hashes by hash-mode determiniation
func HashCheck(mode int, hashString string) {
	value, exists := hashMap[mode]
	if exists {
		// check input string against this
		util.Green(value)
	} else {
		util.Red("Hash pairing doesn't exist")
		return
	}

}

func init() {
	hashMap = make(map[int]string)
	hashMap[1000] = "this is a testaaaa"
	hashMap[5600] = "5600 aaa"
}
