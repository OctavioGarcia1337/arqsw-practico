package div

import (
	"errors"
)

func Division(num, div float32) (float32, error) {
	if div == 0 {
		err := errors.New("Error no se puede dividir por 0")
		return -1, err
	}
	var res float32 = num / div
	return res, nil
}
