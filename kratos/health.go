package kratos

import (
	"encoding/json"
)

type kratosHealthStatus struct {
	HealthStatus string `json:"status"`
}

// Check if Kratos is alive
func (k *KratosClientObj) HealthIsAlive() (bool, error) {
	reqUrl := k.GetPublicUrl() + "health/alive"
	resBytes, err := HttpReqGet(reqUrl, k.config.httpClient, "")
	if err != nil {
		return false, err
	}

	var healtStatus kratosHealthStatus
	err = json.Unmarshal(resBytes, &healtStatus)
	if err != nil {
		return false, err
	}

	if healtStatus.HealthStatus != "ok" {
		return false, nil
	}

	return true, nil
}
