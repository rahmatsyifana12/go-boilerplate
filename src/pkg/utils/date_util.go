package utils

import "time"

// DateUtil is the interface that defines the methods for dates utility
type DateUtil interface {
	GetTimeNowJakarta() time.Time
}

// DateUtilImpl is the implementation of interface DateUtil
type DateUtilImpl struct {}

func (du DateUtilImpl) GetTimeNowJakarta() time.Time {
	jakarta, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	return time.Now().In(jakarta)
}