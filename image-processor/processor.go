package processor

import (
	"bytes"
	"image"
	"image/jpeg"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ProcessImage(imageURL, bucketName string) error {
	// Download the image
	resp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode and compress the image
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return err
	}

	var compressed bytes.Buffer
	err = jpeg.Encode(&compressed, img, &jpeg.Options{Quality: 75})
	if err != nil {
		return err
	}

	// Upload the compressed image to S3
	return uploadToS3(compressed.Bytes(), bucketName, "compressed/"+getFileName(imageURL))
}

func uploadToS3(data []byte, bucket, key string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		return err
	}

	svc := s3.New(sess)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	return err
}

func getFileName(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
