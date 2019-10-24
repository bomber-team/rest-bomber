package gparams

/*WordParams - parameters for word generator*/
type WordParams struct {
	Min     int    `json:"min"`
	Max     int    `json:"max"`
	Charset string `json:"charset"`
}

const (
	// DefaultCharsetEnLower - default pattern for charset inject
	DefaultCharsetEnLower = "abcdefghijklmnopqrstuvwxyz"
	// MinLength - minimal length of generating word by default
	MinLength = 5
	// MaxLength - maximal length of generating word by default
	MaxLength = 150
)
