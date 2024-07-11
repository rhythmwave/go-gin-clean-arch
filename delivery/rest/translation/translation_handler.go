package translation

import (
	"net/http"

	useCase "eira/document/usecase/acs/translation"

	"github.com/gin-gonic/gin"
)

type TranslationHandler struct {
	TranslationUsecase useCase.TranslationUseCase
}

func (t *TranslationHandler) TranslateFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "No file uploaded")
		return
	}

	// Save the file to a temporary folder
	err = c.SaveUploadedFile(file, "/tmp/"+file.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to save file")
		return
	}

	c.String(http.StatusOK, "File saved to temporary folder")

	// Translate the document
	status, err := t.TranslationUsecase.TranslateDocumentFromFile("/tmp/"+file.Filename, "subscriptionKey")
	if err != nil {
		c.String(http.StatusInternalServerError, "Translation failed")
		return
	}
	if status {
		msg := "Translation succeeded"
		c.String(http.StatusOK, msg)
	}
	c.String(http.StatusInternalServerError, "Translation failed")
}
