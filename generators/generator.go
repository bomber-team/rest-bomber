package generators

type Generator struct {
	IpGenerator
	MacGenerator
	WordGenerator
}

//Recursively go through body to generate attribute
func (generator Generator) GenerateForJsonBody(parsed map[string]interface{}, resultOut map[string]interface{}) {
	for key, value := range parsed {
		switch value.(type) {
		case string:
			resultOut[key] = generator.generateAttribute(value.(string))
		case int:
			resultOut[key] = generator.generateAttribute(value.(string))
		case interface{}:
			generator.GenerateForJsonBody(value.(map[string]interface{}), resultOut)
		//handle another json object and another and another ....
		default:
			resultOut[key] = generator.generateAttribute(value.(string))
		}
	}
}

const (
	WordGeneratorEnum     = "WordGenerator"
	PasswordGeneratorEnum = "PasswordGenerator"
	IpGeneratorV4Enum     = "IpGeneratorV4"
	IpGeneratorV6Enum     = "IpGeneratorV6"
	MacGeneratorEnum      = "MacGenerator"
)

//Switch between generator and assign value from user
func (generator Generator) generateAttribute(valueName string) string {
	switch valueName {
	case WordGeneratorEnum:
		return generator.generate()
	case PasswordGeneratorEnum:
		return "******^&*%"
	case IpGeneratorV4Enum:
		return generator.generateV4()
	case IpGeneratorV6Enum:
		return generator.generateV6()
	case MacGeneratorEnum:
		return generator.generateMac()
	default:
		return ""
	}
}
