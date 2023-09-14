package helper

import (
	"fmt"
	"time"
)

func GenerateCode() string {
	return fmt.Sprint(time.Now().Nanosecond())[:6]
}