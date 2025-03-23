package providers

import (
	"FTDS3Status/src/models"
	settings "FTDS3Status/src/settings"
	"fmt"
	"time"
)

// 🔄 Estados posibles: pending, running, failed, success, completed

func AddTaskStatusPending(req models.Request) {
	putTaskStatus(req, settings.PENDING_TASK_TAG, false)
}

func AddTaskStatusRunning(req models.Request) {
	putTaskStatus(req, settings.RUNNING_TASK_TAG, false)
}

func AddTaskStatusFailed(req models.Request) {
	putTaskStatus(req, settings.FAILED_TASK_TAG, true)
}

func AddTaskStatusCompleted(req models.Request) {
	putTaskStatus(req, settings.COMPLETED_TASK_TAG, true)
}

// Función genérica para guardar estados de tareas
func putTaskStatus(req models.Request, statusTag string, withPayload bool) {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	filename := fmt.Sprintf("%s.json", timestamp)

	s3Key := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s",
		settings.SERVICE_NAME,
		settings.ALL_BATCH_TAG,
		req.BatchID,
		settings.ALL_TASKS,
		req.TaskID,
		statusTag,
		filename,
	)

	if withPayload {
		dataFile, err := convertPayloadToBytes(req.Payload)
		if err != nil {
			fmt.Printf("failed to marshal payload for task status %s: %v\n", statusTag, err)
			return
		}
		putObjectOnS3(s3Key, dataFile)
	} else {
		putEmptyObjectOnS3(s3Key)
	}
}
