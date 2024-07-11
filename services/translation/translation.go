package translation

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"eira/document/entities/translation"
)

type TranslationService interface {
	TranslateDocument(input *translation.TranslationInput) (string, error)
	TranslateDocumentFromFile(filePath string, subscriptionKey string) (bool, error)
}

type ACSTranslationService struct {
	Endpoint string
	Key      string
	Region   string
}

func DefACSTranslationService() *ACSTranslationService {
	return &ACSTranslationService{}
}

func (acs *ACSTranslationService) TranslateDocument(input *translation.TranslationInput) (string, error) {

	// Build the request URL
	// uri := acs.Endpoint + "/translator/document/batches"
	// u, _ := url.Parse(uri + "?api-version={date}")
	// method := "POST"
	// q := u.Query()
	// q.Add("from", "en")
	// q.Add("to", "id")
	// u.RawQuery = q.Encode()

	// fmt.Println(u.String())
	// Create the request body
	// body := translation.TranslationInput{Inputs: []translation.Input{{
	// 	StorageType: "file",
	// 	Source:      translation.Source{SourceUrl: cfg.Storage.InputUrl, Language: "en", StorageSource: cfg.Storage.InputContainer},
	// 	Targets:     []translation.Target{{TargetUrl: cfg.Storage.OutputUrl, Language: "id", StorageSource: cfg.Storage.OutputContainer}}}}}
	// b, err := json.Marshal(body)
	// fmt.Println(string(b))
	// if err != nil {
	// 	return "", err
	// }

	// Build the HTTP POST request
	// req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(b))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", err
	// }
	// // Add required headers
	// req.Header.Add("Ocp-Apim-Subscription-Key", acs.Key)
	// req.Header.Add("Ocp-Apim-Subscription-Region", acs.Region)
	// req.Header.Add("Content-Type", "application/json")

	// // Call the Translator API
	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", err
	// }
	// defer res.Body.Close()

	// fmt.Println("response status:", res.Status)
	// fmt.Println("response headers", res.Header)
	// // respBody, _ := ioutil.ReadAll(res.Body)
	// // fmt.Println("Response body:", string(respBody))

	// // Decode the JSON response
	// var translationResponses []translation.TranslationResponse
	// if err := json.NewDecoder(res.Body).Decode(&translationResponses); err != nil {
	// 	fmt.Println(err)
	// 	return "", err
	// }

	// // Format and return the response
	// if len(translationResponses) > 0 && len(translationResponses[0].Translations) > 0 {
	// 	fmt.Println(err)
	// 	return translationResponses[0].Translations[0].Text, nil
	// }
	response := "test"
	return response, fmt.Errorf("no translations found")
}

func (acs *ACSTranslationService) TranslateDocumentFromFile(filePath string, subscriptionKey string) (bool, error) {
	url := "https://rndaitranslator.cognitiveservices.azure.com/translator/document:translate?targetLanguage=id&api-version=2024-05-01"
	headers := map[string]string{
		"Ocp-Apim-Subscription-Key": subscriptionKey,
	}

	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("document", filepath.Base(filePath))
	if err != nil {
		return false, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return false, err
	}
	err = writer.Close()
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		result, err := io.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}
		err = os.WriteFile("result.pptx", result, 0644)
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		fmt.Println("Translation failed. Status code:", resp.StatusCode)
		return false, fmt.Errorf("Translation failed. Status code: %d", resp.StatusCode)
	}
}
