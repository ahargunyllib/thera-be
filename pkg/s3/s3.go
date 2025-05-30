package s3

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/ahargunyllib/thera-be/internal/infra/env"
	"github.com/ahargunyllib/thera-be/pkg/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type CustomS3Interface interface {
	Upload(file *multipart.FileHeader, bucketName, fileName string) (string, error)
	UploadFromBytes(bytes []byte, contentType, bucketName, fileName string) (string, error)
}

type CustomS3Struct struct {
	// session  *session.Session
	// uploader *s3manager.Uploader
	client *s3.Client
}

var S3 = getS3()

func getS3() CustomS3Interface {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(env.AppEnv.AWSAccessKeyID,
			env.AppEnv.AWSSecretAccessKey,
			"")),
		config.WithRegion("apac"),
	)
	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[S3][getS3] failed to load AWS config")
		return nil
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", env.AppEnv.R2AccountID))
	})

	// session := session.Must(session.NewSession(&aws.Config{
	// 	Region: aws.String("ap-southeast-3"), // Asia Pacific (Jakarta)
	// 	Credentials: credentials.NewStaticCredentials(
	// env.AppEnv.AWSAccessKeyID,
	// env.AppEnv.AWSSecretAccessKey,
	// "",
	// 	),
	// }))

	// uploader := s3manager.NewUploader(session)

	return &CustomS3Struct{
		// session:  session,
		// uploader: uploader,
		client: client,
	}
}

func (s *CustomS3Struct) Upload(file *multipart.FileHeader, bucketName, fileName string) (string, error) {
	// Validate input
	if file == nil {
		err := errors.New("file is nil")
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[S3][Upload] invalid input")
		return "", err
	}

	// Open file
	fileContent, err := file.Open()
	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[S3][Upload] failed to open file")
		return "", err
	}
	defer func() {
		if cerr := fileContent.Close(); cerr != nil {
			log.Warn(log.CustomLogInfo{
				"error": cerr.Error(),
			}, "[S3][Upload] failed to close file")
		}
	}()

	// Determine content type with fallback
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream" // Fallback for unknown file types
	}

	// Upload to S3
	// result, err := s.uploader.Upload(&s3manager.UploadInput{
	// 	Bucket:      aws.String(bucketName),
	// 	Key:         aws.String(fileName),
	// 	Body:        fileContent,
	// 	ACL:         aws.String("public-read"),
	// 	ContentType: aws.String(contentType),
	// })
	// if err != nil {
	// 	log.Error(log.CustomLogInfo{
	// 		"error": err.Error(),
	// 	}, "[S3][Upload] failed to upload file")
	// 	return "", err
	// }

	_, err = s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(fileName),
		Body:        fileContent,
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[S3][Upload] failed to upload file")
		return "", err
	}

	// Return public URL
	return fmt.Sprintf("https://%s.hmdtif-filkom.my.id/%s", bucketName, fileName), nil
}

func (s *CustomS3Struct) UploadFromBytes(fileBytes []byte, contentType, bucketName, fileName string) (string, error) {
	// Validate input
	if fileBytes == nil {
		err := errors.New("fileBytes is nil")
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[S3][UploadFromBytes] invalid input")
		return "", err
	}

	// Upload to S3
	_, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
	})

	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[S3][UploadFromBytes] failed to upload file")
		return "", err
	}

	return fmt.Sprintf("https://%s.hmdtif-filkom.my.id/%s", bucketName, fileName), nil
}
