package constant

const (
	SESSION_KEY              = "OCTOPUS_SESSIONS"
	SYSTEM_WORKSPACE_DEFAULT = "default-workspace"
	SYSTEM_TYPE_AI           = "OCTOPUS_AI"
	SYSTEM_TYPE_ADMIN        = "OCTOPUS_ADMIN"

	SYSTEM_ROOT_NAME = "octopus"

	PREPARING = "preparing"
	PENDING   = "pending"
	RUNNING   = "running"
	FAILED    = "failed"
	SUCCEEDED = "succeeded"
	STOPPED   = "stopped"
	SUSPENDED = "suspended"
	UNKNOWN   = "unknown"

	JOB_TYPE    = "jobType"
	NotebookJob = "notebookjob"
	TrainJob    = "trainjob"

	REDIS_MINIO_REMOVING_OBJECT_SET = "minio-removing-object-set"
)
