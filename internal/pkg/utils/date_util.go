package utils

import "time"

var GetTimeNowJakarta = func () (time.Time, error) {
	jakarta, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(jakarta), nil
}