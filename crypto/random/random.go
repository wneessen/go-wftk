package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const pwChar string = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`
const pwCharHuman string = `abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ2345678`
const pwSpecial string = `!"#$%&'()*+,-.\/:;<=>?@[]^_{}|}`
const pwSpHuman string = `"#%*+-/:;=\\_|~`

// Provide random bytes
func GenerateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Generate a random (human-)readable string
func GenerateRandomString(sLength int, isHumanReadable bool, isNoSpecial bool) (string, error) {
	charList := pwChar
	if isHumanReadable {
		charList = pwCharHuman
	}
	if !isNoSpecial {
		if isHumanReadable {
			charList = charList + pwSpHuman
		} else {
			charList = charList + pwSpecial
		}
	}
	listLen := len(charList)
	charArray := []byte(charList)
	returnString := make([]byte, sLength)
	for i := 0; i < sLength; i++ {
		randNum, err := GenerateRandomNum(listLen)
		if err != nil {
			return "", err
		}
		returnString[i] = charArray[randNum]
	}

	return string(returnString), nil
}

// Generate a random number with given maximum value
func GenerateRandomNum(maxNum int) (int, error) {
	if maxNum <= 0 {
		err := fmt.Errorf("provided maxNum is <= 0: %v", maxNum)
		return 0, err
	}
	maxNumBigInt := big.NewInt(int64(maxNum))
	if !maxNumBigInt.IsUint64() {
		err := fmt.Errorf("big.NewInt() generation returned negative value: %v", maxNumBigInt)
		return 0, err
	}
	randNum64, err := rand.Int(rand.Reader, maxNumBigInt)
	if err != nil {
		return 0, err
	}
	randNum := int(randNum64.Int64())
	if randNum < 0 {
		err := fmt.Errorf("generated random number does not fit as int64: %v", randNum64)
		return 0, err
	}
	return randNum, nil
}
