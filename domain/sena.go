package domain

import "time"

type Sena struct {
	Numbers []int     `bson:"numbers" json:"numbers"`
	Code    int       `bson:"code" json:"code"`
	Date    time.Time `bson:"date" json:"date"`
}
