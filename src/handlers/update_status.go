package handler

import (
	"FTDS3Status/src/models"
	"FTDS3Status/src/providers"
	"context"
)

// The batch status updater is handled only by worker rq
func choiceCommandsBatchsStatus(req models.Request) (models.Response, error) {
	// Maneja status batch
	switch req.Status {
	case "pending":
		err := providers.AddBatchPendingStatus(req)
		if err != nil {
			return models.Response{
				Message: err.Error(),
			}, err
		}
		return models.Response{
			Message: "Ok, batch actualizado como pending",
		}, nil

	case "running":
		err := providers.AddBatchRunningStatus(req)
		if err != nil {
			return models.Response{
				Message: err.Error(),
			}, err
		}
		return models.Response{
			Message: "Ok, batch actualizado como running",
		}, nil

	case "failed":
		err := providers.AddBatchFailedStatus(req)
		if err != nil {
			return models.Response{
				Message: err.Error(),
			}, err
		}
		return models.Response{
			Message: "ups, batch actualizado como failed",
		}, nil

	case "completed":
		err := providers.AddBatchCompletedStatus(req)
		if err != nil {
			return models.Response{
				Message: err.Error(),
			}, err
		}
		return models.Response{
			Message: "yeah, batch marcado como completado!",
		}, nil
	case "success":
		//todo another thing
	default:
		return models.Response{
			Message: "Unknown status add-batch-status: " + req.Status,
		}, nil
	}
	return models.Response{
		Message: "Unknown status add-batch-status: " + req.Status,
	}, nil
}

// The task status updater is handled only by stepfunctions
func choicheCommandTaskStatus(req models.Request) (models.Response, error) {
	switch req.Status {
	case "pending":
		err := providers.AddTaskPendingStatus(req)
		if err != nil {
			return models.Response{Message: "error updating task pending status"}, err
		}
		return models.Response{Message: "task pending status updated"}, nil

	case "running":
		err := providers.AddTaskRunningStatus(req)
		if err != nil {
			return models.Response{Message: "error updating task running status"}, err
		}
		return models.Response{Message: "task running status updated"}, nil

	case "failed":
		err := providers.AddTaskFailedStatus(req)
		if err != nil {
			return models.Response{Message: "error updating task failed status"}, err
		}
		return models.Response{Message: "task failed status updated"}, nil

	case "completed":
		err := providers.AddTaskCompletedStatus(req)
		if err != nil {
			return models.Response{Message: "error updating task completed status"}, err
		}
		return models.Response{Message: "task completed status updated"}, nil

	default:
		return models.Response{Message: "invalid task status: " + req.Status}, nil
	}
}

func HandleRequest(ctx context.Context, req models.Request) (models.Response, error) {
	switch req.Command {
	case "add-batch-status":
		ok, err := choiceCommandsBatchsStatus(req)
		if err != nil {
			return models.Response{
				Message: "Unknown command: " + err.Error(),
			}, nil
		}
		return ok, nil
	case "add-task-status":
		choicheCommandTaskStatus(req)
	default:
		return models.Response{
			Message: "Unknown command: " + req.Command,
		}, nil
	}
	return models.Response{
		Message: "Unknown command: " + req.Command,
	}, nil
}
