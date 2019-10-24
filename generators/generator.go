package generators

import "github.com/bomber-team/rest-bomber/generators/gparams"

type (
	/*Generator - main structure for Generators merged*/
	Generator struct {
		IP   *IPGenerator
		MAC  *MacGenerator
		WORD *WordGenerator
		PASS *PasswordGenerator
	}

	/*IGenerator - interface with methods needs for Generator*/
	IGenerator interface {
		New(params string) *Generator // initialize new instance of generator
		GenerateForJSONBody(parsed map[string]interface{}) map[string]interface{}
	}
)

/*New - initialize new instance of generator*/
func (gen *Generator) New(params string) *Generator {
	gen.IP = &IPGenerator{}
	gen.MAC = &MacGenerator{}
	word := &WordGenerator{}
	pass := &PasswordGenerator{}

	gen.PASS = pass.New()
	gen.WORD = word.New("")
	return gen
}

/*GenerateForJSONBody - generating value for requested parameterized types strings*/
func (gen *Generator) GenerateForJSONBody(parsed map[string]interface{}) (resultOut map[string]interface{}) {
	for key, value := range parsed {
		switch value.(type) {
		case string:
			resultOut[key] = gen.generateAttribute(value.(int))
		case int:
			resultOut[key] = gen.generateAttribute(value.(int))
		case interface{}:
			resultOut = gen.GenerateForJSONBody(value.(map[string]interface{}))
		//handle another json object and another and another ....
		default:
			resultOut[key] = gen.generateAttribute(value.(int))
		}
	}
	return
}

//Switch between generator and assign value from user
func (gen *Generator) generateAttribute(valueName int) string {
	switch valueName {
	case gparams.WordGeneratorEnum:
		return gen.WORD.Generate(0)
	case gparams.PasswordGeneratorEnum:
		return gen.PASS.Generate()
	case gparams.IpGeneratorV4Enum:
		return gen.IP.Generate(gparams.IPV4)
	case gparams.IpGeneratorV6Enum:
		return gen.IP.Generate(gparams.IPV6)
	case gparams.MacGeneratorEnum:
		return gen.MAC.Generate()
	default:
		return ""
	}
}
