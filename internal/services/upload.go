package services

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func sanitizeS3Key(name string) string {
	// Only allow alphanumeric, dash, underscore, dot
	result := make([]rune, 0, len(name))
	for _, r := range name {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' || r == '.' {
			result = append(result, r)
		} else {
			result = append(result, '_')
		}
	}
	return string(result)
}

func UploadFile(filename string, data io.Reader) (string, error) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})

	svc := s3.New(sess)
	safeFilename := sanitizeS3Key(filename)
	key := fmt.Sprintf("%d_%s", time.Now().UnixNano(), safeFilename)

	buffer := new(bytes.Buffer)
	_, err := io.Copy(buffer, data)
	if err != nil {
		return "", err
	}

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buffer.Bytes()),
	})

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", os.Getenv("S3_BUCKET_NAME"), key), nil
}

func DeleteFileFromS3(key string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})
	if err != nil {
		return err
	}

	svc := s3.New(sess)
	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	return svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(key),
	})
}
