package model

import "time"

type Base struct {
	ID         int64     `json:"id" form:"id" comment:"主键id"`
	StartTime  time.Time `json:"start_time" form:"start_time" comment:"开始时间"`
	EndTime    time.Time `json:"end_time" form:"end_time" comment:"截止时间"`
	DeleteTime time.Time `json:"delete_time" form:"delete_time" comment:"删除时间"`
}
