package acs

import (
	"context"
	"eira/document/config"
	"fmt"
	"io"
	"os"
)

func uploadBlobFile(containerName string, blobName string, cfg config.Config) {

	// Load client
	client := GetServiceClientSAS(cfg.Storage.AccountUrl, cfg.Storage.InputToken)

	// Open the file for reading
	file, err := os.OpenFile("path/to/sample/file", os.O_RDONLY, 0)
	fmt.Println(err)

	defer file.Close()

	// Upload the file to the specified container with the specified blob name
	_, err = client.UploadFile(context.TODO(), containerName, blobName, file, nil)
	fmt.Println(err)
}

func UploadBlobStream(file io.Reader, containerName string, blobName string, cfg config.Config) error {
	ctx := context.Background()

	// Load client
	// client := GetServiceClientSAS(cfg.Storage.AccountUrl, cfg.Storage.InputToken)
	client := GetServiceClientSharedKey(cfg.Storage.Account, cfg.Storage.Key)

	fmt.Println(client)
	// ListContainers(client)

	// Upload the file to the specified container with the specified blob name
	_, err := client.UploadStream(ctx, containerName, blobName, file, nil)
	fmt.Println(err)

	if err != nil {
		return err
	}
	return nil
}
