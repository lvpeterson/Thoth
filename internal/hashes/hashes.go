package hashes

import "thoth/internal/util"

var hashMap map[int]string
var hashDelim map[int]string

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
	hashMap[1000] = "domain\\username:4422421:aad3b435b51404eeaad3b435b51404ee:1635072b2da4acb1de9e9fd44d835085:::"
	hashMap[3000] = "domain\\username:4422421:aad3b435b51404eeaad3b435b51404ee:1635072b2da4acb1de9e9fd44d835085:::"
	hashMap[5500] = "u4-netntlm::kNS:338d08f8e26de93300000000000000000000000000000000:9526fb8c23a90751cdd619b6cea564742e1e4bf33006ba41:cb8086049ec4736c"
	hashMap[5600] = "admin::N46iSNekpT:08ca45b7d7ea58ee:88dcbe4446168966a153a0064958dac6:5c7830315c7830310000000000000b45c67103d07d7b95acd12ffa11230e0000000052920b85f78d013c31cdb3b92f5d765c783030"
	hashMap[13100] = "$krb5tgs$23$*user$realm$test/spn*$63386d22d359fe42230300d56852c9eb$891ad31d09ab89c6b3b8c5e5de6c06a7f49fd559d7a9a3c32576c8fedf705376cea582ab5938f7fc8bc741acf05c5990741b36ef4311fe3562a41b70a4ec6ecba849905f2385bb3799d92499909658c7287c49160276bca0006c350b0db4fd387adc27c01e9e9ad0c20ed53a7e6356dee2452e35eca2a6a1d1432796fc5c19d068978df74d3d0baf35c77de12456bf1144b6a750d11f55805f5a16ece2975246e2d026dce997fba34ac8757312e9e4e6272de35e20d52fb668c5ed"
	hashMap[18200] = "$krb5asrep$23$user@domain.com:3e156ada591263b8aab0965f5aebd837$007497cb51b6c8116d6407a782ea0e1c5402b17db7afa6b05a6d30ed164a9933c754d720e279c6c573679bd27128fe77e5fea1f72334c1193c8ff0b370fadc6368bf2d49bbfdba4c5dccab95e8c8ebfdc75f438a0797dbfb2f8a1a5f4c423f9bfc1fea483342a11bd56a216f4d5158ccc4b224b52894fadfba3957dfe4b6b8f5f9f9fe422811a314768673e0c924340b8ccb84775ce9defaa3baa0910b676ad0036d13032b0dd94e3b13903cc738a7b6d00b0b3c210d1f972a6c7cae9bd3c959acf7565be528fc179118f28c679f6deeee1456f0781eb8154e18e49cb27b64bf74cd7112a0ebae2102ac"

	hashDelim = make(map[int]string)
	hashDelim[1000] = ":"
	hashDelim[3000] = ":"
	hashDelim[5500] = ":"
	hashDelim[5600] = ":"
	hashDelim[13100] = "$"
	hashDelim[18200] = "$"

}
