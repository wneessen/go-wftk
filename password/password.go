package password

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"github.com/nbutton23/zxcvbn-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PwStrengthIndicator struct {
	Score         int
	Entropy       float64
	CrackTime     float64
	CrackTimeText string
}

type HIBPResp struct {
	HasBeenPwned bool
	NumHits      int64
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

// Query HIBP database for provided password string
func CheckHIBP(p string) (HIBPResp, error) {
	shaSum := fmt.Sprintf("%x", sha1.Sum([]byte(p)))
	firstPart := shaSum[0:5]
	secondPart := shaSum[5:]
	respObj := HIBPResp{
		HasBeenPwned: false,
		NumHits:      0,
	}

	httpClient := &http.Client{Timeout: time.Second * 2}
	httpRes, err := httpClient.Get("https://api.pwnedpasswords.com/range/" + firstPart)
	if err != nil {
		return respObj, err
	}
	defer func() {
		if err := httpRes.Body.Close(); err != nil {
			fmt.Printf("An error occured while closing HTTP response body: %v\n", err)
		}
	}()

	scanObj := bufio.NewScanner(httpRes.Body)
	for scanObj.Scan() {
		scanLine := strings.SplitN(scanObj.Text(), ":", 2)
		if strings.ToLower(scanLine[0]) == secondPart {
			respObj.HasBeenPwned = true
			getHits, err := strconv.ParseInt(scanLine[1], 10, 64)
			if err != nil {
				fmt.Printf("HIBP: Unable to parse number: %v. Setting to 1", err)
				getHits = 1
			}
			respObj.NumHits = getHits
			break
		}
	}
	if err := scanObj.Err(); err != nil {
		return respObj, err
	}

	return respObj, nil
}
