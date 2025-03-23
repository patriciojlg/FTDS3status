package mock

import (
	"FTDS3Status/src/models"
	"FTDS3Status/src/settings"
	"context"
)

func MockAddPendingTask(idBatch string, taskID string) (context.Context, models.Request) {
	ctx := context.Background()
	req := models.Request{
		Command:     settings.ADD_TASK_STATUS,
		BatchID:     idBatch,
		TaskID:      taskID,
		Status:      "pending",
		ServiceTask: "sec",
		Payload: map[string]interface{}{
			"rut": taskID,
		},
	}
	return ctx, req
}

func MockAddRunningTask(idBatch string, taskID string) (context.Context, models.Request) {
	ctx := context.Background()
	req := models.Request{
		Command:     settings.ADD_TASK_STATUS,
		BatchID:     idBatch,
		TaskID:      taskID,
		Status:      "running",
		ServiceTask: "sec",
		Payload: map[string]interface{}{
			"rut": taskID,
		},
	}
	return ctx, req
}

func MockAddFailedTask(idBatch string, taskID string) (context.Context, models.Request) {
	ctx := context.Background()
	req := models.Request{
		Command:     settings.ADD_TASK_STATUS,
		BatchID:     idBatch,
		TaskID:      taskID,
		Status:      "failed",
		ServiceTask: "sec",
		Payload: map[string]interface{}{
			"error": "rut no encontrado",
			"rut":   taskID,
		},
	}
	return ctx, req
}

func MockAddCompletedTask(idBatch string, taskID string) (context.Context, models.Request) {
	ctx := context.Background()
	req := models.Request{
		Command:     settings.ADD_TASK_STATUS,
		BatchID:     idBatch,
		TaskID:      taskID,
		Status:      "completed",
		ServiceTask: "sec",
		Payload: map[string]interface{}{
			"rut":          taskID,
			"certificados": []string{"Ingeniería", "SEC Clase D"},
		},
	}
	return ctx, req
}
