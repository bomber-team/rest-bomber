package gparams

/*PasswordParams - params for generator passwords*/
type PasswordParams struct {
	Min      int    `json:"min"`      // minimal amount repeat
	Max      int    `json:"max"`      // maximal amount of repeat
	Position string `json:"position"` // type of generate position ()
	Alphabet int    `json:"alpha"`    // require alphabet(default is en = 0|EN = 1|ru = 2|RU = 3|special = 4|numbers = 5)
}

const (
	// ENAlphaBet - en resource for generating
	ENAlphaBet = "ABCDEFGHIKLMNOPRSTWXYZ"
	// RUAlphaBet - ru resource for generating
	RUAlphaBet = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦШЩЬЪЫЭЮЯ"
	// SPAlphaBet - spectial characters for generating
	SPAlphaBet = "~!@#%^&*()_-=+"
	// NMAlphaBet - numbers for generating
	NMAlphaBet = "0123456789"

	//#####types
	ENLower = iota + 0
	ENUpper = iota + 1
	RULower = iota + 2
	RUUpper = iota + 3

	// Special Characters
	SPCH = iota + 4
	// Numbers
	NM = iota + 5
)
