package generators

import (
	"math/rand"
	"restbomber/gparams"
	"strings"
	"time"
)

/*PasswordGenerator - generator for password by defined pattern*/
type (
	PasswordGenerator struct {
		PatternType string                   `json:"pattern"`
		HeadParams  []gparams.PasswordParams `json:"params_generator"` // Head params need contain elements in strict order (0 - UpperCase, 1 - lowerCase, 2 - Special Characters, 3 - Numbers)
		generator   *rand.Rand
	}

	// IPasswordGenerator - interface which contain all methods for structure passwordgenerator
	IPasswordGenerator interface {
		New() *PasswordGenerator
		Generate() string // generating new password with defining patterns
		maker(chunks [][]rune) string
		sizeChunks(chunks [][]rune) int // amount characters do merge
		GetRandomRune(sizeChunks, sizeChunk int) (int, int)
		RemoveUsedValueChunk(chunk, value int, chunks [][]rune) [][]rune
		gUpperLetters() []rune      // generating upper letters
		gLowerLetters() []rune      // generating lower letters
		gNumbers() []rune           // generating numbers
		gSpecialCharacters() []rune // generating special characters
		cycleGen() []rune           // generating iteration
		getAlphaBet(tp int) string  // get alphabet for generation cycle
	}
)

/*New - initialize new instance of generator with time salt*/
func (gen *PasswordGenerator) New() *PasswordGenerator {
	source := rand.NewSource(time.Now().UnixNano())
	gen.generator = rand.New(source)
	return gen
}

/*Generate - generating new password with defining pattern and params*/
func (gen *PasswordGenerator) Generate() string {
	chunks := make([][]rune, len(gen.HeadParams))
	for index, value := range gen.HeadParams {
		switch value.Alphabet {
		case gparams.ENLower:
		case gparams.RULower:
			chunks[index] = append(chunks[index], gen.gLowerLetters(&value)...)
			break
		case gparams.ENUpper:
		case gparams.RUUpper:
			chunks[index] = append(chunks[index], gen.gUpperLetters(&value)...)
			break
		case gparams.SPCH:
			chunks[index] = append(chunks[index], gen.gSpecialCharacters(&value)...)
			break
		case gparams.NM:
			chunks[index] = append(chunks[index], gen.gNumbers(&value)...)
			break
		default:
			chunks[index] = append(chunks[index], ' ')
			break
		}
	}
	return gen.maker(chunks)
}

/*maker - mixer of chunks values*/
func (gen *PasswordGenerator) maker(chunks [][]rune) (result string) {
	currentSize := gen.sizeChunks(chunks)
	currentIteration := 0

	for currentIteration < currentSize {
		randChunk := gen.generator.Intn(len(chunks))
		rchunk, rvalue := gen.GetRandomRune(randChunk, len(chunks[randChunk]))
		result += string(chunks[rchunk][rvalue])
		chunks = gen.RemoveUsedValueChunk(rchunk, rvalue, chunks)
		currentIteration++
	}
	return result
}

/*RemovedUsedValueChunk - removing  usefull chunks value*/
func (gen *PasswordGenerator) RemoveUsedValueChunk(chunk, value int, chunks [][]rune) [][]rune {
	chunkNew := append(chunks[chunk][:value], chunks[chunk][value+1:]...)
	chunksRes := append(chunks[:chunk], chunkNew)
	chunksRes = append(chunksRes, chunks[chunk+1:]...)
	return chunksRes
}

func (gen *PasswordGenerator) GetRandomRune(currentSizeChunks, currentSizeChunk int) (int, int) {
	return gen.generator.Intn(currentSizeChunks), gen.generator.Intn(currentSizeChunk)
}

func (gen *PasswordGenerator) sizeChunks(chunks [][]rune) (size int) {
	for _, value := range chunks {
		for i, _ := range value {
			size++
			if i != 0 {
			}
		}
	}
	return size
}

func (gen *PasswordGenerator) gUpperLetters(param *gparams.PasswordParams) []rune {
	return gen.cycleGen(param)
}

func (gen *PasswordGenerator) gLowerLetters(param *gparams.PasswordParams) []rune {
	return gen.cycleGen(param)
}

func (gen *PasswordGenerator) gNumbers(param *gparams.PasswordParams) []rune {
	return gen.cycleGen(param)
}

func (gen *PasswordGenerator) gSpecialCharacters(param *gparams.PasswordParams) []rune {
	return gen.cycleGen(param)
}

func (gen *PasswordGenerator) cycleGen(params *gparams.PasswordParams) []rune {
	amountIterations := gen.generator.Intn(params.Max-params.Min) + params.Min
	currentIteration := 0

	buffer := make([]rune, amountIterations)

	alphabet := gen.getAlphaBet(params.Alphabet)

	for currentIteration < amountIterations {
		buffer[currentIteration] = rune(alphabet[gen.generator.Intn(len(alphabet))])
	}
	return buffer
}

func (gen *PasswordGenerator) getAlphaBet(tp int) string {
	switch tp {
	case gparams.ENLower:
		return strings.ToLower(gparams.ENAlphaBet)
	case gparams.ENUpper:
		return gparams.ENAlphaBet
	case gparams.RULower:
		return strings.ToLower(gparams.RUAlphaBet)
	case gparams.RUUpper:
		return gparams.RUAlphaBet
	case gparams.SPCH:
		return gparams.SPAlphaBet
	case gparams.NM:
		return gparams.NMAlphaBet
	default:
		return ""
	}
}
