package generators

import (
	"math/rand"

	"github.com/bomber-team/bomber-proto-contracts/golang/rest_contracts"
)

const (
	ru = "абвгдеёжзиклмнопрстуфхцшщцыъьэюя"
	en = "abcdefghiklmnoprstuqvwxyz"
)

func GenerateWord(config rest_contracts.GeneratorConfig_WordGeneratorConfig) string {

	stringSize := rand.Int31n(config.WordGeneratorConfig.MaxLetters-config.WordGeneratorConfig.MinLetters) + config.WordGeneratorConfig.MinLetters
	var index int32 = 0
	var resultString string = ""
	for ; index < stringSize; index++ {
		resultString += string(genByAlphabet(config.WordGeneratorConfig.Language))
	}
	return resultString
}

func genByAlphabet(typeA rest_contracts.Language) rune {
	sizeAlphabet := 0
	switch typeA {
	case rest_contracts.Language_RU:
		sizeAlphabet = len(ru)
	case rest_contracts.Language_EN:
		sizeAlphabet = len(en)
	default:
		sizeAlphabet = len(en)
	}
	return rune(en[rand.Intn(sizeAlphabet)])
}
