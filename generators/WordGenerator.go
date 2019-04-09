package generators

import (
	"math/rand"
	"time"
)

type WordGenerator struct {
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func (wordGenerator WordGenerator) generate() string {
	return wordGenerator.generateWithCharset(5, charset)
}

func (wordGenerator WordGenerator) generateWithLength(length int) string {
	return wordGenerator.generateWithCharset(length, charset)
}

func (WordGenerator) generateWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
