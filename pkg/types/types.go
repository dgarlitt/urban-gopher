package types

// Definition - struct for definitions response
type Definition struct {
	Word string `json:"word"`
	Text string `json:"definition"` // definition text
}

// ProviderParams - required parameters for using a Definition Provider
type ProviderParams struct {
	Term   string
	APIKey string
}

// DefinitionProvider - interface for provider of definitions
type DefinitionProvider interface {
	LookupDefinition(params *ProviderParams) (*Definition, error)
}
