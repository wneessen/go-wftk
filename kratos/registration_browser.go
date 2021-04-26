package kratos

import (
	"encoding/json"
)

type RegistrationFlow struct {
	IsActive  bool                    `json:"active,omitempty"`
	ExpiresAt EpochTimeStamp          `json:"expires_at"`
	Id        Uuid4String             `json:"id"`
	IssuedAt  EpochTimeStamp          `json:"issued_at"`
	Messages  *[]ApiMessage           `json:"messages,omitempty"`
	Methods   RegistrationFlowMethods `json:"methods"`
	RequstUrl string                  `json:"requst_url"`
	Type      string                  `json:"type"`
	Error     GenericError            `json:"error"`
}

type RegistrationFlowMethods struct {
	Password RegistrationFlowMethod `json:"password"`
}

type RegistrationFlowMethod struct {
	Method string                       `json:"method"`
	Config RegistrationFlowMethodConfig `json:"config"`
}

type RegistrationFlowMethodConfig struct {
	Action   string        `json:"action"`
	Fields   []FormField   `json:"fields"`
	Messages *[]ApiMessage `json:"messages,omitempty"`
	Method   string        `json:"method"`
}

// Get the registration flow by providing the flow id
func (k *KratosClientObj) GetRegistrationFlow(i string) (RegistrationFlow, error) {
	reqUrl := k.GetPublicUrl() + "self-service/registration/flows?id=" + i
	resBytes, err := HttpReqGet(reqUrl, k.config.httpClient, "")
	if err != nil {
		return RegistrationFlow{}, err
	}

	var regFlow RegistrationFlow
	err = json.Unmarshal(resBytes, &regFlow)
	if err != nil {
		return RegistrationFlow{}, err
	}

	return regFlow, nil
}
