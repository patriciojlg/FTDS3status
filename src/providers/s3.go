package providers

import (
	settings "FTDS3Status/src/settings"
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client   *s3.S3
	bucketName string
)

func init() {
	// Crear sesión compartida
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(settings.AWS_REGION), // Asegúrate de definir esta variable
	}))
	s3Client = s3.New(sess)

	// Cargar configuraciones
	bucketName = settings.BUCKET_NAME

}
func putObjectOnS3(s3Key string, data []byte) error {
	// Subida al bucket
	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(s3Key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("application/json"),
	})
	return err

}
func putEmptyObjectOnS3(s3Key string) error {
	// Subida al bucket
	var data []byte
	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(s3Key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("application/json"),
	})
	return err
}

func convertPayloadToBytes(payload map[string]interface{}) ([]byte, error) {
	var empty []byte
	data, err := json.Marshal(payload)
	if err != nil {
		return empty, err
	}
	return data, nil
}
