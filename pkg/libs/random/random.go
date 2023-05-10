package random

import (
	crypto_rand "crypto/rand"
	math_rand "math/rand"
	"time"
)

func init() {
	math_rand.Seed(time.Now().UnixNano())
}

// GetRandomBytes - Generate random bytes
func GetRandomBytes(n int, alphabets []byte) []byte {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	var randby bool
	num, err := crypto_rand.Read(bytes)
	if num != n || err != nil {
		randby = true
	}
	for i, b := range bytes {
		if len(alphabets) == 0 {
			if randby {
				bytes[i] = alphanum[math_rand.Intn(len(alphanum))]
			} else {
				bytes[i] = alphanum[b%byte(len(alphanum))]
			}
		} else {
			if randby {
				bytes[i] = alphabets[math_rand.Intn(len(alphabets))]
			} else {
				bytes[i] = alphabets[b%byte(len(alphabets))]
			}
		}
	}
	return bytes
}

// GetRandomString - Generate random string
func GetRandomString(n int, alphabets []byte) string {
	return string(GetRandomBytes(n, alphabets))
}

// GetRandomInt - Generate random int
func GetRandomInt(min int, max int) int {
	return math_rand.Intn(max-min) + min
}
