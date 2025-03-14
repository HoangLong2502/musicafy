package zingmp3

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

// e is url path request, n is query params string
func GenerateHash(e string, n string) (sig string) {
	fmt.Println("==== chuoi N", strings.ReplaceAll(n, "&", ""))

	// Băm SHA256 để tạo r
	sha256Hasher := sha256.New()
	sha256Hasher.Write([]byte(strings.ReplaceAll(n, "&", "")))
	r := hex.EncodeToString(sha256Hasher.Sum(nil))

	// Tạo khóa HMAC-SHA512 với key
	key := []byte("acOrvUS15XRW2o9JksiK1KgQ6Vbds8ZW")
	hmacSha512 := hmac.New(sha512.New, key)
	hmacSha512.Write([]byte(e + r))

	// Trả về chuỗi mã hóa SHA512 dạng hex
	return hex.EncodeToString(hmacSha512.Sum(nil))
}

// func main() {
// 	e := "/api/v2/song/get/streaming"
// 	id := "Z78BZ0D7"
// 	hashedValue := generateHash(e, id)
// 	fmt.Println("===Generated Hash:", hashedValue)
// }
