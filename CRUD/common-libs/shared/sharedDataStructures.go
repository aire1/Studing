package shared

//типы формата *Data - это типы для Kafka

type Note struct {
	Id        int    `json:"id"`
	UserId    string `json:"UserId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Базовая структура для Kafka, поля которой, так или иначе, содержат все сообщения
type BaseTaskData struct {
	Login  string `json:"login"`
	TaskId string `json:"taskid"`
}

type RegistrationData struct {
	BaseTaskData
	Passhash string `json:"passhash"`
}

type AuthorizationGetData struct {
	BaseTaskData
	Passhash string `json:"passhash"`
}

type AuthorizationCheckData struct {
	BaseTaskData
	JwtToken string `json:"jwtToken"`
}

type CreateNoteData struct {
	BaseTaskData
	Note
}

type GetNoteData struct {
	BaseTaskData
	Offset int
	Count  int
}

type DeleteNoteData struct {
	BaseTaskData
	NoteId int
}

type UpdateNoteData struct {
	BaseTaskData
	Note
}
