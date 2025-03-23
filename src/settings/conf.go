package settings

const (
	// commands
	ADD_BATCH_STATUS = "add-batch-status"
	ADD_TASK_STATUS  = "add-task-status"
	//s3 parents paths
	SERVICE_NAME     = "FTD_CODELCO"
	ALL_BATCH_TAG    = "ALL-BATCHS"
	BATCH_STATUS_TAG = "BATCH-STATUS"
	ALL_TASKS        = "ALL-TASK"
	///s3  task tags
	PENDING_TASK_TAG   = "pending-task"
	RUNNING_TASK_TAG   = "running-task"
	COMPLETED_TASK_TAG = "completed-task"
	FAILED_TASK_TAG    = "failed-task"
	///s3  batch tags
	PENDING_BATCH_TAG   = "pending-batch"
	RUNNING_BATCH_TAG   = "running-batch"
	COMPLETED_BATCH_TAG = "completed-batch"
	SUCCESS_BATCH_TAG   = "success-batch"
	FAILED_BATCH_TAG    = "failed-batch"

	// s3
	BUCKET_NAME = "ftd-prueba-status"
	AWS_REGION  = "us-east-1"
)
