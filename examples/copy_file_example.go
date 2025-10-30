package main

import (
	"fmt"
	"log"

	storage_go "github.com/supabase-community/storage-go"
)

func main() {
	// Initialize the storage client
	storageClient := storage_go.NewClient("https://your-project.supabase.co/storage/v1", "your-api-key", nil)

	// Example 1: Copy file within the same bucket
	fmt.Println("Example 1: Copy file within the same bucket")
	response1, err := storageClient.CopyFile(
		"avatars",                    // source bucket
		"public/avatar1.png",         // source file path
		"private/avatar1-copy.png",   // destination file path
	)
	if err != nil {
		log.Printf("Error copying file: %v\n", err)
	} else {
		fmt.Printf("File copied successfully: %s\n", response1.Key)
	}

	// Example 2: Copy file across buckets
	fmt.Println("\nExample 2: Copy file across buckets")
	destinationBucket := "avatars-backup"
	response2, err := storageClient.CopyFile(
		"avatars",                    // source bucket
		"public/avatar1.png",         // source file path
		"backup/avatar1.png",         // destination file path
		storage_go.CopyFileOptions{
			DestinationBucket: &destinationBucket,
		},
	)
	if err != nil {
		log.Printf("Error copying file across buckets: %v\n", err)
	} else {
		fmt.Printf("File copied across buckets successfully: %s\n", response2.Key)
	}

	// Example 3: Copy file without metadata
	fmt.Println("\nExample 3: Copy file without copying metadata")
	copyMetadata := false
	response3, err := storageClient.CopyFile(
		"avatars",                    // source bucket
		"public/avatar1.png",         // source file path
		"public/avatar1-no-meta.png", // destination file path
		storage_go.CopyFileOptions{
			CopyMetadata: &copyMetadata,
		},
	)
	if err != nil {
		log.Printf("Error copying file without metadata: %v\n", err)
	} else {
		fmt.Printf("File copied without metadata: %s\n", response3.Key)
	}

	// Example 4: Copy file across buckets without metadata
	fmt.Println("\nExample 4: Copy file across buckets without metadata")
	response4, err := storageClient.CopyFile(
		"avatars",                    // source bucket
		"public/avatar1.png",         // source file path
		"backup/avatar1-no-meta.png", // destination file path
		storage_go.CopyFileOptions{
			DestinationBucket: &destinationBucket,
			CopyMetadata:      &copyMetadata,
		},
	)
	if err != nil {
		log.Printf("Error copying file across buckets without metadata: %v\n", err)
	} else {
		fmt.Printf("File copied across buckets without metadata: %s\n", response4.Key)
	}
}
