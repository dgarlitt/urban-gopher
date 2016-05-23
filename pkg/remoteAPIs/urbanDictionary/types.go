package urbanDictionary

type errorResponse struct {
	Message string
}

// DictionaryResults struct is used for unmarshaling urban dictionary responses
type DictionaryResults struct {
	UserDefinitions []*UserDefinition `json:"list"`
	ResultType      string            `json:"result_type"`
	Sounds          []string          `json:"sounds"`
	Tags            []string          `json:"tags"`
	Error           error
}

// UserDefinition is an individual definition in an urban dictionary response
type UserDefinition struct {
	Author      string `json:"author"`
	CurrentVote string `json:"current_vote"`
	Defid       int    `json:"defid"`
	Text        string `json:"definition"`
	Example     string `json:"example"`
	Permalink   string `json:"permalink"`
	ThumbsDown  int    `json:"thumbs_down"`
	ThumbsUp    int    `json:"thumbs_up"`
	Word        string `json:"word"`
}
