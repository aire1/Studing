package shared

//типы формата *Data - это типы для Kafka
//типы формата *Status - это типы для передачи статуса через Redis

//куча дублей сделана лишь для возможного расширения в будущем
//(разумеется, его не будет, но я пытаюсь в идеальную архитектуру)

// Базовая структура для Kafka, поля которой, так или иначе, содержат все сообщения
type BaseTaskData struct {
	Login  string `json:"login"`
	TaskId string `json:"taskid"`
}

// Базовая структура для Redis, поля которой, так или иначе, содержат все сообщения
type BaseTaskStatus struct {
	Result string `json:"result"`
	Info   string `json:"info"`
}

type RegistrationData struct {
	BaseTaskData
	Passhash string `json:"passhash"`
}

type RegistrationStatus struct {
	BaseTaskStatus
}

type AuthorizationGetStatus struct {
	BaseTaskStatus
}

type AuthorizationGetData struct {
	BaseTaskData
	Passhash string `json:"passhash"`
}

type AuthorizationCheckData struct {
	BaseTaskData
	JwtToken string `json:"jwtToken"`
}

type AuthorizationCheckStatus struct {
	BaseTaskStatus
}

type CreateTaskData struct {
	BaseTaskData
	Info string `json:"info"`
}

type CreateTaskStatus struct {
	BaseTaskStatus
}
