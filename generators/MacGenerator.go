package generators


import (
	"crypto/rand"
	"fmt"
)

type MacGenerator struct{}

func (MacGenerator) GenRandomMac() string {
	token := make([]byte, 6)
	_, err := rand.Read(token)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", token[0], token[1], token[2], token[3], token[4], token[5])
}