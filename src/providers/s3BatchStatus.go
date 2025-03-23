package providers

import (
	"FTDS3Status/src/models"
	settings "FTDS3Status/src/settings"
	"fmt"
	"time"
)

// ✅ Tipos de status que requieren payload
var batchStatusesWithPayload = map[string]bool{
	settings.PENDING_BATCH_TAG:   true,
	settings.FAILED_BATCH_TAG:    true,
	settings.SUCCESS_BATCH_TAG:   true,
	settings.COMPLETED_BATCH_TAG: true,
}

// ✅ Función pública por cada estado

func AddBatchPendingStatus(req models.Request) error {
	return saveBatchStatus(req, settings.PENDING_BATCH_TAG)
}

func AddBatchRunningStatus(req models.Request) error {
	return saveBatchStatus(req, settings.RUNNING_BATCH_TAG)
}

func AddBatchFailedStatus(req models.Request) error {
	return saveBatchStatus(req, settings.FAILED_BATCH_TAG)
}

func AddBatchSuccessStatus(req models.Request) error {
	return saveBatchStatus(req, settings.SUCCESS_BATCH_TAG)
}

func AddBatchCompletedStatus(req models.Request) error {
	return saveBatchStatus(req, settings.COMPLETED_BATCH_TAG)
}

// 🧠 Lógica central unificada

func saveBatchStatus(req models.Request, statusTag string) error {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	filename := fmt.Sprintf("%s.json", timestamp)
	s3Key := buildBatchS3Key(req.BatchID, statusTag, filename)

	// Si el estado requiere payload
	if batchStatusesWithPayload[statusTag] {
		dataFile, err := convertPayloadToBytes(req.Payload)
		if err != nil {
			return fmt.Errorf("failed to convert payload: %w", err)
		}
		if err := putObjectOnS3(s3Key, dataFile); err != nil {
			return fmt.Errorf("failed to upload file with payload: %w", err)
		}
		return nil
	}

	// Estado sin payload → archivo vacío
	if err := putEmptyObjectOnS3(s3Key); err != nil {
		return fmt.Errorf("failed to upload empty file: %w", err)
	}
	return nil
}

// 🧱 Construcción del path S3 para batches
func buildBatchS3Key(batchID, statusTag, filename string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s",
		settings.SERVICE_NAME,
		settings.ALL_BATCH_TAG,
		batchID,
		settings.BATCH_STATUS_TAG,
		statusTag,
		filename,
	)
}
