package v1

type ActionState string

const (
	ActionCompletedState ActionState = "completed"
	ActionRunningState   ActionState = "running"
)

type CommandResult string

const (
	CommandSucceedResult CommandResult = "succeed"
	CommandFailedResult  CommandResult = "failed"
)
