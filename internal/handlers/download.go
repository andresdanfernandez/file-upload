package handlers

import (
	"net/http"
	"os"
	"time"

	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

func DownloadHandler(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing file key in URL"})
		return
	}

	// Debug: Log environment variables
	log.Printf("AWS_REGION: %s", os.Getenv("AWS_REGION"))
	log.Printf("AWS_ACCESS_KEY_ID: %s", os.Getenv("AWS_ACCESS_KEY_ID"))
	log.Printf("S3_BUCKET_NAME: %s", os.Getenv("S3_BUCKET_NAME"))
	
	// Check if AWS credentials are set
	if os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		log.Println("AWS credentials not found in environment")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AWS credentials not configured"})
		return
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})

	if err != nil {
		log.Println("Failed to create AWS session:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not initialize AWS session"})
		return
	}

	s3Client := s3.New(sess)

	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(10 * time.Minute)
	if err != nil {
		log.Println("Failed to presign URL:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate presigned URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": urlStr})
}
