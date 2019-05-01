package generators

import (
	"encoding/hex"
	"math/rand"
	"rest-bomber/gparams"
)

type (
	/*IPGenerator - main structure for generator ip address*/
	IPGenerator struct {
	}

	/*IIPGenerator - interface for all methods needs for IPGenerator*/
	IIPGenerator interface {
		Generate(tp int) string // generating new ip by version of protoco;
		generateV4() string     // generating new ip of 4 version protocol
		generateV6() string     // generating new ip of 6 version protocol
		randomHex(n int) (string, error)
	}
)

/*Generate - generating new ip*/
func (gen *IPGenerator) Generate(tp int) string {
	switch tp {
	case gparams.IPV4:
		return gen.generateV4()
	case gparams.IPV6:
		return gen.generateV6()
	default:
		return ""
	}
}

//Generate ip in v4 format 192.168.19.2
func (gen *IPGenerator) generateV4() string {
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
func (gen *IPGenerator) generateV6() string {
	sb := ""
	for i := 0; i < 7; i++ {
		generated, _ := gen.randomHex(4)
		sb += string(generated) + ":"
	}

	generated, _ := gen.randomHex(4)
	sb += string(generated)
	return sb
}

//Generate random hex for ipv6
func (gen *IPGenerator) randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
