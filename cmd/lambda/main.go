package main

import (
	handlers "FTDS3Status/src/handlers"
	mocks "FTDS3Status/src/mocks"
	"FTDS3Status/src/models"
	"FTDS3Status/src/settings"
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
)

func mockNewPending() (context.Context, models.Request) {
	ctx := context.Background()
	id := uuid.New() // genera un UUID v4
	idStr := id.String()
	req := models.Request{
		Command: settings.ADD_BATCH_STATUS,
		BatchID: idStr,
		TaskID:  "",
		Status:  "pending",
		Payload: map[string]interface{}{
			"result": "todo bien",
		},
	}
	return ctx, req
}
func mockAddRunningBatch(idBatch string) (context.Context, models.Request) {
	ctx := context.Background()

	req := models.Request{
		Command: settings.ADD_BATCH_STATUS,
		BatchID: idBatch,
		TaskID:  "",
		Status:  "running",
		Payload: map[string]interface{}{
			"result": "todo bien",
		},
	}
	return ctx, req
}

func mockAddFailedBatch(idBatch string) (context.Context, models.Request) {
	ctx := context.Background()

	req := models.Request{
		Command: settings.ADD_BATCH_STATUS,
		BatchID: idBatch,
		TaskID:  "",
		Status:  "failed",
		Payload: map[string]interface{}{
			"result": "todo bien",
		},
	}
	return ctx, req
}

func mockAddCompletedBatch(idBatch string) (context.Context, models.Request) {
	ctx := context.Background()

	req := models.Request{
		Command: settings.ADD_BATCH_STATUS,
		BatchID: idBatch,
		TaskID:  "",
		Status:  "completed",
		Payload: map[string]interface{}{
			"result": "todo bien",
		},
	}
	return ctx, req
}
func main() {
	// Si se ejecuta en AWS Lambda, usar Lambda Handler
	if _, exists := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); exists {
		lambda.Start(handlers.HandleRequest)
	} else {
		ctxRunningStatus, reqMocPendingTask := mocks.MockAddPendingTask("600b5e67-89e4-4d34-86bb-dd3867257db3", "15838946-0")
		handlers.HandleRequest(ctxRunningStatus, reqMocPendingTask)
		ctxRunningStatus, mockAddRunningTask := mocks.MockAddRunningTask("600b5e67-89e4-4d34-86bb-dd3867257db3", "15838946-0")
		handlers.HandleRequest(ctxRunningStatus, mockAddRunningTask)
		ctxRunningStatus, mockComplete := mocks.MockAddCompletedTask("600b5e67-89e4-4d34-86bb-dd3867257db3", "15838946-0")
		handlers.HandleRequest(ctxRunningStatus, mockComplete)
		ctxRunningStatus, mockFailed := mocks.MockAddFailedTask("600b5e67-89e4-4d34-86bb-dd3867257db3", "15838946-0")
		handlers.HandleRequest(ctxRunningStatus, mockFailed)
	}
}
