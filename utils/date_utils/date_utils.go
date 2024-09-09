package date_utils

import "time"

const (
	apiDateLayout = time.RFC3339
)

func GetNowString() string {
	return time.Now().UTC().Format(apiDateLayout)
}
