package providers

import (
	"FTDS3Status/src/models"
	"FTDS3Status/src/settings"
	"fmt"
	"time"
)

func AddTaskPendingStatus(req models.Request) error {
	return writeTaskStatus(req, settings.PENDING_TASK_TAG, true)
}

func AddTaskRunningStatus(req models.Request) error {
	return writeTaskStatus(req, settings.RUNNING_TASK_TAG, false)
}

func AddTaskFailedStatus(req models.Request) error {
	return writeTaskStatus(req, settings.FAILED_TASK_TAG, true)
}

func AddTaskCompletedStatus(req models.Request) error {
	return writeTaskStatus(req, settings.COMPLETED_TASK_TAG, true)
}

// Reutilizable
func writeTaskStatus(req models.Request, status string, withPayload bool) error {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	filename := fmt.Sprintf("%s.json", timestamp)
	s3Key := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s",
		settings.SERVICE_NAME,
		settings.ALL_BATCH_TAG,
		req.BatchID,
		settings.ALL_TASKS,
		req.TaskID,
		status,
		filename,
	)

	if withPayload {
		dataFile, err := convertPayloadToBytes(req.Payload)
		if err != nil {
			fmt.Println("error converting payload:", err)
			return err
		}
		return putObjectOnS3(s3Key, dataFile)
	} else {
		return putEmptyObjectOnS3(s3Key)
	}
}
