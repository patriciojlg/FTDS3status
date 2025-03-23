package models

type Response struct {
	Message string `json:"message"`
	S3Key   string `json:"s3_key,omitempty"`
}
