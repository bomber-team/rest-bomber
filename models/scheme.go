package models

type (
	/*Scheme - contain main info for datas in request*/
	Scheme struct {
		Headers *HeadersMap `json:"headers"`
		Params  *ParamsMap  `json:"params"`
	}

	/*HeadersMap - contain headers which insert in request packet*/
	HeadersMap struct {
		Method    string `json:"method"`
		Authority string `json:"authority"`
	}

	/*ParamsMap - contain params which insert in request query or body*/
	ParamsMap struct {
		Query []*RequestParams `json:"request"`
		Body  []*BodyParams    `json:"body"`
	}

	/*RequestParams - query params for request*/
	RequestParams struct {
		Name          string `json:"name"`
		GeneratorNeed bool   `json:"generatorNeed"`
		Value         string `json:"value"`
	}

	/*BodyParams - params inserting in body of request packet*/
	BodyParams struct {
		Name            string          `json:"name"`
		GeneratorNeed   bool            `json:"generatorNeed"`
		Generator       string          `json:"generator"`
		GeneratorConfig GeneratorConfig `json:"generatorConfig"`
		ValueObject     *BodyParams     `json:"value_obj"`
		ValueString     string          `json:"value_str"`
	}

	/*GeneratorConfig - configuration for generators*/
	GeneratorConfig struct {
		MinLetters int    `json:"min_letters"`
		MaxLetters int    `json:"max_letters"`
		Language   string `json:"language"`
		Alphabet   string `json:"alphabet"`
	}
)
