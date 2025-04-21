package vo

type UserRegistratorRequest struct {
	Email   string `json:"email,omitempty"`
	Purpose string `json:"purpose,omitempty"`
}
