package translation

import (
	entity "eira/document/entities/translation"
	service "eira/document/services/translation"
)

type TranslationUseCaseImpl struct {
	AiProviders service.TranslationService
}

func (t *TranslationUseCaseImpl) TranslateDocument(input *entity.TranslationInput) (string, error) {

	return t.AiProviders.TranslateDocument(input)
}

func (t *TranslationUseCaseImpl) TranslateDocumentFromFile(filePath string, subscriptionKey string) (bool, error) {

	result, err := t.AiProviders.TranslateDocumentFromFile(filePath, subscriptionKey)

	return result, err
}
