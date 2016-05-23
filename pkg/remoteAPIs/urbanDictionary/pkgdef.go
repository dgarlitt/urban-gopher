/*
Package urbanDictionary provides an interface to make a
request to the Urban Dictionary API provided by mashape.

The package can be implemented by passing a Params
struct to the LookupDefinition function.

Example:

	params := &urbanDictionary.Params{
		Term:   "batman",
		APIKey: "yourMashApeAPIKey",
	}

	dictionary, err := urbanDictionary.LookupDefinition(params)
*/
package urbanDictionary
