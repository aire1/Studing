package shared

type TaskStatus interface {
	GetResult() string
	SetResult(result string)
}

// Базовая структура для Redis, поля которой, так или иначе, содержат все сообщения
type BaseTaskStatus struct {
	Result string `json:"result"`
	Info   string `json:"info"`
}

func (b *BaseTaskStatus) GetResult() string {
	return b.Result
}

func (b *BaseTaskStatus) SetResult(result string) {
	b.Result = result
}

type RegistrationStatus struct {
	BaseTaskStatus
}

type AuthorizationGetStatus struct {
	BaseTaskStatus
}

type AuthorizationCheckStatus struct {
	BaseTaskStatus
}

type CreateNoteStatus struct {
	BaseTaskStatus
}

type GetNoteStatus struct {
	BaseTaskStatus
	Notes []Note
}
