package translation

import "eira/document/entities/translation"

type TranslationUseCase interface {
	TranslateDocument(input *translation.TranslationInput) (string, error)
	TranslateDocumentFromFile(filePath string, subscriptionKey string) (bool, error)
}
