package generators

import (
	"encoding/hex"
	"math/rand"
)

//Structure for ips generators
type IpGenerator struct {
}

//Generate ip in v4 format 192.168.19.2
func (IpGenerator IpGenerator) generateV4() string{
	sb := ""
	for i := 0; i < 3; i++ {
		generated := rand.Int31n(256)

		sb += string(generated) + "."
	}

	generated := rand.Int31n(256)
	sb += string(generated)

	return sb
}

//Generate ip in v6 format
func (IpGenerator IpGenerator) generateV6() string {
	sb := ""
	for i := 0; i < 7; i++ {
		generated, _ := randomHex(4)
		sb += string(generated) + ":"
	}

	generated, _ := randomHex(4)
	sb += string(generated)
	return sb
}

//Generate random hex for ipv6
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
