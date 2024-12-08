package utils

import "time"

func FormateDate(t time.Time, layout string) string {
	return t.Format(layout)
}
