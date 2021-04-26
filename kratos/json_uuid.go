package kratos

import (
	"fmt"
	"github.com/google/uuid"
)

type Uuid4String string

// UnmarshalJSON function for parsing/validating UUID4
func (u *Uuid4String) UnmarshalJSON(id []byte) error {
	uuidString := string(id)
	parseUuid, err := uuid.Parse(uuidString)
	if err != nil {
		return fmt.Errorf("uuid parsing failed: %v", err)
	}
	*(*string)(u) = parseUuid.String()

	return nil
}
