package password

import (
	zxcvbn "github.com/nbutton23/zxcvbn-go"
)

type PwStrengthIndicator struct {
	Score         int
	Entropy       float64
	CrackTime     float64
	CrackTimeText string
}

// Test strenght of password and return the score
func CheckPasswordStrength(p string) *PwStrengthIndicator {
	pwScoreObj := zxcvbn.PasswordStrength(p, nil)
	strengthObj := &PwStrengthIndicator{
		Score:         pwScoreObj.Score,
		Entropy:       pwScoreObj.Entropy,
		CrackTime:     pwScoreObj.CrackTime,
		CrackTimeText: pwScoreObj.CrackTimeDisplay,
	}

	return strengthObj
}
