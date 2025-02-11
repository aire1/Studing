package shared

import "encoding/json"

//типы формата *Data - это типы для Kafka
//типы формата *Status - это типы для передачи статуса через Redis

//куча дублей сделана лишь для возможного расширения в будущем
//(разумеется, его не будет, но я пытаюсь в идеальную архитектуру)

// Универсальный сериализуемый тип
type Serializable[T any] struct{}

func (s *Serializable[T]) Marshal() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Serializable[T]) Unmarshal(data []byte) error {
	return json.Unmarshal(data, s)
}

// Базовая структура для Kafka, поля которой, так или иначе, содержат все сообщения
type BaseTaskData struct {
	Login  string `json:"login"`
	TaskId string `json:"taskid"`
}

type RegistrationData struct {
	Serializable[RegistrationData]
	BaseTaskData
	Passhash string `json:"passhash"`
}

type RegistrationStatus struct {
	Serializable[RegistrationStatus]
	Result string `json:"result"`
	Info   string `json:"info"`
}

type AuthorizationGetStatus struct {
	Serializable[AuthorizationGetStatus]
	Result string `json:"result"`
	Info   string `json:"info"`
}

type AuthorizationGetData struct {
	Serializable[AuthorizationGetData]
	BaseTaskData
	Passhash string `json:"passhash"`
}

type AuthorizationCheckData struct {
	Serializable[AuthorizationCheckData]
	BaseTaskData
	JwtToken string `json:"jwtToken"`
}

type AuthorizationCheckStatus struct {
	Serializable[AuthorizationCheckStatus]
	Result string `json:"result"`
	Info   string `json:"info"`
}

type CreateTaskData struct {
	Serializable[CreateTaskData]
	BaseTaskData
	Info string `json:"info"`
}
