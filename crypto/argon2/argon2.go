package argon2

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/wneessen/go-wftk/crypto/random"
	"golang.org/x/crypto/argon2"
	"strings"
)

type Params struct {
	// The amount of memory used by the algorithm (in kibibytes).
	Memory uint32

	// The number of iterations over the memory.
	Iterations uint32

	// The number of threads (or lanes) used by the algorithm.
	// Recommended value is between 1 and runtime.NumCPU().
	Parallelism uint8

	// Length of the random salt. 16 bytes is recommended for password hashing.
	SaltLength uint32

	// Length of the generated key. 16 bytes or more is recommended.
	KeyLength uint32
}

// Argon2 package default params
var DefaultParams = &Params{
	Memory:      256 * 1024,
	Iterations:  20,
	Parallelism: 2,
	SaltLength:  20,
	KeyLength:   32,
}

// Create an Argon2id hash
func CreateHash(pwString string, p *Params) (string, error) {
	hashSalt, err := random.GenerateRandomBytes(p.SaltLength)
	if err != nil {
		return "", err
	}

	pwKey := argon2.IDKey([]byte(pwString), hashSalt, p.Iterations, p.Memory,
		p.Parallelism, p.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(hashSalt)
	b64Key := base64.RawStdEncoding.EncodeToString(pwKey)

	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory,
		p.Iterations, p.Parallelism, b64Salt, b64Key), nil
}

// Validate an Argon2id hash
func ValidateHash(pwString, argonHash string) (bool, error) {
	hashParams, hashSalt, pwKey, err := decodeHash(argonHash)
	if err != nil {
		return false, err
	}

	// Create a hash from the given PwString
	testKey := argon2.IDKey([]byte(pwString), hashSalt, hashParams.Iterations, hashParams.Memory,
		hashParams.Parallelism, hashParams.KeyLength)

	// Check key length are identical
	orgKeyLen := int32(len(pwKey))
	testKeyLen := int32(len(testKey))
	if subtle.ConstantTimeEq(orgKeyLen, testKeyLen) == 0 {
		return false, nil
	}

	// Check the testKey and original pwKey are identical
	if subtle.ConstantTimeCompare(pwKey, testKey) == 1 {
		return true, nil
	}

	return false, nil
}

// Decode an Argon2id hash
func decodeHash(h string) (*Params, []byte, []byte, error) {
	hashVals := strings.Split(h, "$")
	if len(hashVals) != 6 {
		return nil, nil, nil, fmt.Errorf("argon2id: hash is not in the correct format")
	}

	var version int
	_, err := fmt.Sscanf(hashVals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, fmt.Errorf("argon2id: incompatible version of argon2")
	}

	hashParams := &Params{}
	_, err = fmt.Sscanf(hashVals[3], "m=%d,t=%d,p=%d", &hashParams.Memory, &hashParams.Iterations,
		&hashParams.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	hashSalt, err := base64.RawStdEncoding.Strict().DecodeString(hashVals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	hashParams.SaltLength = uint32(len(hashSalt))

	pwKey, err := base64.RawStdEncoding.Strict().DecodeString(hashVals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	hashParams.KeyLength = uint32(len(pwKey))

	return hashParams, hashSalt, pwKey, nil
}
