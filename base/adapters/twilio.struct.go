package adapters

//||------------------------------------------------------------------------------------------------||
//|| SendGrid Personalization Struct
//||------------------------------------------------------------------------------------------------||

type SendGridPersonalization struct {
	To []SendGridEmail `json:"to"`
}

//||------------------------------------------------------------------------------------------------||
//|| SendGrid Email Struct
//||------------------------------------------------------------------------------------------------||

type SendGridEmail struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| SendGrid Content Struct
//||------------------------------------------------------------------------------------------------||

type SendGridContent struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

//||------------------------------------------------------------------------------------------------||
//|| SendGrid Request Struct
//||------------------------------------------------------------------------------------------------||

type SendGridRequest struct {
	Personalizations []SendGridPersonalization `json:"personalizations"`
	From             SendGridEmail             `json:"from"`
	Subject          string                    `json:"subject"`
	Content          []SendGridContent         `json:"content"`
}
