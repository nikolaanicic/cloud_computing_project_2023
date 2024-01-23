package responsemodels

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"time"
)

type Token struct {
	Value string `json:"value"`
}

func randomSalt() []byte {

	rand.Seed(time.Now().UnixMicro())
	charset := "abcdefghijklmonpqrstuvwxyz"
	maxLen := 10

	retval := make([]byte, maxLen)

	for i := 0; i < maxLen; i++ {
		retval[i] = charset[rand.Intn(len(charset))]
	}

	return retval
}

func NewToken(username string) Token {
	hashed := sha512.Sum512(append([]byte(username), randomSalt()...))
	return Token{hex.EncodeToString(hashed[:])}
}

func (t Token) AsJson() []byte {
	data, _ := json.Marshal(t)

	return data
}
