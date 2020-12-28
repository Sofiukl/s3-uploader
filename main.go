package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/sofiukl/s3-uploader/utils"

	"github.com/aws/aws-sdk-go/aws"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	start := time.Now()
	if len(os.Args) != 3 {
		panic("please pass argument source and foldername")
	}
	source := os.Args[1]
	folderName := os.Args[2]

	fmt.Println("Source path:", source)
	fmt.Println("Folder name where file will be uploaded:", folderName)

	// Create session
	config, err := utils.LoadConfig(".")
	sess, err := awsSession.NewSessionWithOptions(awsSession.Options{
		Profile: config.AWSProfile,
		Config: aws.Config{
			Region: aws.String(config.AWSRegion),
		},
	})
	if err != nil {
		utils.PrintErrorf("Fail to create aws session %s", err)
		return
	}

	// read files
	readDir(source, folderName, sess, config)
	fmt.Printf("upload took %v\n", time.Since(start))
}

func readDir(source string, folderName string, sess *awsSession.Session, config utils.Config) {
	ch := make(chan string)
	files, err := ioutil.ReadDir(source)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		if f.IsDir() {
			s := filepath.Join(source, f.Name())
			fn := filepath.Join(folderName, f.Name())
			readDir(s, fn, sess, config)
		} else {
			s := filepath.Join(source, f.Name())
			go func(s string, filename string, folderName string, ch chan string) {
				uploadFileS3(s, filename, folderName, ch, sess, config)
			}(s, f.Name(), folderName, ch)
		}

	}

	for _, f := range files {
		if !f.IsDir() {
			fmt.Println(<-ch)
		}
	}
}

func uploadFileS3(path string, filename string, foldername string, ch chan string, sess *awsSession.Session, config utils.Config) {
	uploadedFile, _ := os.Open(path)
	defer uploadedFile.Close()

	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.AWSBucketName),
		Key:    aws.String(filepath.Join(foldername, filename)),
		Body:   uploadedFile,
	})
	if err != nil {
		utils.PrintErrorf("Unable to upload %q to %q, %v", filename, config.AWSBucketName, err)
	}
	ch <- fmt.Sprintf("Successfully uploaded %q to bucket %q\n", filepath.Join(foldername, filename), config.AWSBucketName)
}
