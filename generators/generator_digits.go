package generators

import (
	"math/rand"

	"github.com/bomber-team/bomber-proto-contracts/golang/rest_contracts"
)

func GenerateDigits(config rest_contracts.GeneratorConfig_DigitGeneratorConfig) int32 {
	return rand.Int31n(config.DigitGeneratorConfig.EndTo-config.DigitGeneratorConfig.StartFrom) + config.DigitGeneratorConfig.StartFrom
}
