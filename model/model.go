package model

type Job struct {
	JobName string
	JobData string
	JobFunc string
}

func (Job) TableName() string {
	return "job"
}

type JobsFunctions struct {
	Jobs []Job
}

type Email struct {
	To      string
	Subject string
	Body    string
}

type Contact struct {
	Id          int    `gorm:"primary_key"`
	FirstName   string `gorm:"<-"`
	LastName    string `gorm:"<-"`
	Email       string `gorm:"<-"`
	PhoneNumber string `gorm:"<-"`
}

func (Contact) TableName() string {
	return "contact"
}

type WhatsAppMessage struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             struct {
		Body string `json:"body"`
	} `json:"text"`
}
