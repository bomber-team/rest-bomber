package generators

type Generator struct {
}

func (generator Generator) generate(parsed map[string]interface{}) {
	for _, value := range parsed {
		switch value.(type) {
		case interface{}:
		//handle another json object and aother and another ....
		default:
			//generate
		}
	}
}

const (
	WordGeneratoEnum      = "WordGenerator"
	PasswordGeneratorEnum = "PasswordGenerator"
	IpGeneratorEnum       = "IpGenerator"
	MacGeneratorEnum      = "MacGenerator"
)

func (generator Generator) generateAttribute(valueName string) string {
	switch valueName {
	case WordGeneratoEnum:

		return "Keks"
	case PasswordGeneratorEnum:
		return "KEks password"
	case IpGeneratorEnum:
		return "Ip keks"
	case MacGeneratorEnum:
		return "Mac Keks"
	}
}
