package time

import (
	"strings"
	"time"
)

func GetTime() string {
	return strings.Split(time.Now().String(), ".")[0]
}
