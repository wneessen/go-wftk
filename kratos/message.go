package kratos

type ApiMessage struct {
	Context struct{} `json:"object,omitempty"`
	Id      int64    `json:"id,omitempty"`
	Text    string   `json:"text,omitempty"`
	Type    string   `json:"type,omitempty"`
}
