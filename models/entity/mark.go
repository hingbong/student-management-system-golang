package entity

import (
	"system-management-system/main/models"
	"time"
)

type Mark struct {
	MarkId     uint      `gorm:"primary_key;column:markid;AUTO_INCREMENT"json:"markId"`
	StuId      uint      `gorm:"not null;column:stuId"json:"stuId"`
	CourseName string    `gorm:"not null"json:"courseName"`
	BaseScore  float64   `gorm:"not null"json:"baseScore"`
	TestScore  float64   `gorm:"not null"json:"testScore"`
	FinalScore float64   `gorm:"not null"json:"finalScore"`
	AddDate    time.Time `gorm:"not null"json:"addDate"`
	Note       string    `json:"note"`
}

func (Mark) TableName() string {
	return "stu_mark_info"
}

func (mark *Mark) Insert() error {
	return models.DB.Create(mark).Error
}
