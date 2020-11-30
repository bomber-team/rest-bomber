package generators

import (
	"github.com/bomber-team/bomber-proto-contracts/golang/rest_contracts"
	"github.com/lucasjones/reggen"
)

func GenerateByRegexp(config *rest_contracts.GeneratorConfig_RegexpConfig) string {
	str, err := reggen.Generate(config.RegexpConfig.Pattern, 100)
	if err != nil {
		return ""
	}
	return str
}
