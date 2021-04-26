package kratos

type GenericError struct {
	ErrorCode int64         `json:"error_code,omitempty"`
	Debug     string        `json:"debug,omitempty"`
	Details   *ErrorDetails `json:"details,omitempty"`
	Message   string        `json:"message,omitempty"`
	Reason    string        `json:"reason,omitempty"`
	Request   string        `json:"request,omitempty"`
	Status    string        `json:"status,omitempty"`
}

type ErrorDetails struct {
	RedirectUrl string `json:"redirect_to,omitempty"`
}
