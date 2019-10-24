package generators

import (
	"math/rand"
	"time"

	"github.com/bomber-team/rest-bomber/generators/gparams"
)

type (
	/*WordGenerator -*/
	WordGenerator struct {
		PatternType string `json:"pattern"`
		generator   *rand.Rand
		Charset     string `json:"charset"`
	}
	/*IWordGenerator - main interface for word generator*/
	IWordGenerator interface {
		New(charset string) *WordGenerator // initialize new instance with charset(if nill then set default charset)
		Generate() string                  // generate new word
		generateWithLength(length int)     // generate word with defined length
	}
)

/*New - initialize new instance of generator with defined charset*/
func (gen *WordGenerator) New(charset string) *WordGenerator {
	if charset != "" {
		gen.Charset = charset
	}

	gen.generator = rand.New(rand.NewSource(time.Now().UnixNano()))

	return gen
}

/*Generate - generating new word with length or random length if length == 0*/
func (gen *WordGenerator) Generate(length int) string {
	if length == 0 {
		length = gen.generator.Intn(gparams.MaxLength-gparams.MinLength) + gparams.MinLength
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = gen.Charset[gen.generator.Intn(len(gen.Charset))]
	}
	return string(b)
}
