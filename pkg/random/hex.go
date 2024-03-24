package random

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func Hex(n int) string {
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		// A random generator only can't fail, otherwise, should be give a fatal
		// error, because this error on generating random numbers can cause a lot
		// of not previsible bugs.
		log.Fatal(err)
	}

	return hex.EncodeToString(bytes)
}
