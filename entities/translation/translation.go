package translation

type TranslationInput struct {
	Inputs []Input `json:"inputs"`
}

type Input struct {
	StorageType string   `json:"storageType"`
	Source      Source   `json:"source"`
	Targets     []Target `json:"targets"`
}

type Source struct {
	SourceUrl     string `json:"sourceUrl"`
	StorageSource string `json:"storageSource"`
	Language      string `json:"language"`
}

type Target struct {
	TargetUrl     string `json:"targetUrl"`
	Language      string `json:"language"`
	StorageSource string `json:"storageSource"`
}

type Glossary struct {
	TargetUrl string `json:"glossaryUrl"`
	Language  string `json:"format"`
}

type TranslationRequest struct {
	Text string `json:"Text"`
}

type TranslationResponse struct {
	Translations []struct {
		Text string `json:"text"`
		To   string `json:"to"`
	} `json:"translations"`
	Error *ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
}
