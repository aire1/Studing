package shared

//типы формата *Data - это типы для Kafka
//типы формата *Status - это типы для передачи статуса через Redis

//куча дублей сделана лишь для возможного расширения в будущем
//(разумеется, его не будет, но я пытаюсь создать идеальную архитектуру)

type RegistrationData struct {
	Login    string `json:"login"`
	Passhash string `json:"passhash"`
	Taskid   string `json:"taskid"`
}

type RegistrationStatus struct {
	Result string `json:"result"`
	Info   string `json:"info"`
}

type AuthorizationGetStatus struct {
	Result string `json:"result"`
	Info   string `json:"info"`
}

type AuthorizationGetData struct {
	Login    string `json:"login"`
	Passhash string `json:"passhash"`
	Taskid   string `json:"taskid"`
}

type AuthorizationCheckData struct {
	Username string `json:"username"`
	JwtToken string `json:"jwtToken"`
	Taskid   string `json:"taskid"`
}

type AuthorizationCheckStatus struct {
	Result string `json:"result"`
	Info   string `json:"info"`
}
