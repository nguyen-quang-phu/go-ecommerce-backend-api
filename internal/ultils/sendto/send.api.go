package sendto

type MailRequest struct {
	ToEmail     string `json:"to_email,omitempty"`
	MessageBody string `json:"message_body,omitempty"`
	Subject     string `json:"subject,omitempty"`
	Attachment  string `json:"attachment,omitempty"`
}

func sendEmailToJavaByAPI(otp string, email string, purpost string) error {
	return nil
}
