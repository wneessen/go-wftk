package password

import (
	"fmt"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
)

func CheckPasswordStrength(p string) {
	foo := zxcvbn.PasswordStrength(p, nil)
	fmt.Println(foo)

}
