package utils

import (
	"bytes"
	"encoding/base64"
	"strings"
	"context"


	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DecodeBase64Image(b64 string) ([]byte, error) {
	if strings.Contains(b64, ",") {
		parts := strings.SplitN(b64, ",", 2)
		b64 = parts[1]
	}

	return base64.StdEncoding.DecodeString(b64)
}

func UploadBase64Image(
	ctx context.Context,
	client *s3.Client,
	bucket string,
	key string,
	b64 string,
	contentType string,
) error {

	data, err := DecodeBase64Image(b64)
	if err != nil {
		return err
	}

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &key,
		Body:        bytes.NewReader(data),
		ContentType: &contentType,
	})

	return err
}
