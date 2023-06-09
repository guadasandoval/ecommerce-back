package password

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/pkg/libs/random"
	"encoding/hex"
	"hash"
)

var (
	saltLength = 10
)

// Encode - Encode the password
func Encode(rawPwd string) string {
	salt := random.GetRandomString(saltLength, nil)
	return EncodeWithSalt(rawPwd, salt)
}

// EncodeWithSalt - Encode de password with salt
func EncodeWithSalt(rawPwd string, salt string) string {

	pwd := PBKDF2([]byte(rawPwd), []byte(salt), 10000, 50, sha256.New)
	return (salt + hex.EncodeToString(pwd))
}

// CheckValid - Check if a password is valid
func CheckValid(checkPwd string, encodedPwd string) bool {

	salt := encodedPwd[:saltLength]
	newEncodedPwd := EncodeWithSalt(checkPwd, salt)
	return (newEncodedPwd == encodedPwd)
}

// PBKDF2 - Hash
// http://code.google.com/p/go/source/browse/pbkdf2/pbkdf2.go?repo=crypto
func PBKDF2(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {

	prf := hmac.New(h, password)
	hashLen := prf.Size()
	numBlocks := (keyLen + hashLen - 1) / hashLen

	var buf [4]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	U := make([]byte, hashLen)
	for block := 1; block <= numBlocks; block++ {
		// N.B.: || means concatenation, ^ means XOR
		// for each block T_i = U_1 ^ U_2 ^ ... ^ U_iter
		// U_1 = PRF(password, salt || uint(i))
		prf.Reset()
		prf.Write(salt)
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)

		// U_n = PRF(password, U_(n-1))
		for n := 2; n <= iter; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for x := range U {
				T[x] ^= U[x]
			}
		}
	}
	return dk[:keyLen]
}
