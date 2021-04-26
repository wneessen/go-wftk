package kratos

type FormField struct {
	IsDisabled  bool          `json:"disabled,omitempty"`
	Messages    *[]ApiMessage `json:"messages,omitempty"`
	Name        string        `json:"name"`
	I18nName    string
	Placeholder string
	Pattern     string `json:"pattern,omitempty"`
	IsRequired  bool   `json:"required,omitempty"`
	Type        string `json:"type"`
	Value       string `json:"value"`
}
