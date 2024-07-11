package acs

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func GetServiceClientSAS(accountURL string, sasToken string) *azblob.Client {
	// Create a new service client with an existing SAS token

	// Append the SAS to the account URL with a "?" delimiter
	accountURLWithSAS := fmt.Sprintf("%s?%s", accountURL, sasToken)
	fmt.Println(accountURLWithSAS)
	client, err := azblob.NewClientWithNoCredential(accountURLWithSAS, nil)
	fmt.Println(err)

	return client
}

func GetServiceClientSharedKey(accountName string, accountKey string) *azblob.Client {
	// Create a new service client with shared key credential
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	handleError(err)

	accountURL := fmt.Sprintf("https://%s.blob.core.windows.net", accountName)

	client, err := azblob.NewClientWithSharedKeyCredential(accountURL, credential, nil)
	handleError(err)

	return client
}

func GenerateSharedKeySASToken(accountName string, accountKey string, containerName string) {

	credential, _ := azblob.NewSharedKeyCredential(accountName, accountKey)

	fmt.Println(credential)

	// Append the SAS token to the BlobURL
	// sasURL := fmt.Sprintf("%s?%s", blobURL, sasToken)
	// fmt.Println("SAS URL:", sasURL)
}
