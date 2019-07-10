package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type Student struct {
	StuId      uint      `gorm:"primary_key;column:stuid"json:"stuId"`
	StuName    string    `gorm:"not null;column:stuname"json:"stuName"`
	Profession int       `gorm:"not null"json:"profession"`
	Date       time.Time `gorm:"not null"json:"date"`
	Sex        int       `gorm:"not null"json:"sex"`
	IdNum      string    `gorm:"not null"json:"idNum"`
	TotalScore int       `gorm:"not null"json:"totalScore"`
	Note       string    `json:"note"`
}

type Mark struct {
	MarkId     uint      `gorm:"primary_key;column:markid"`
	StuId      string    `gorm:"not null";column:"stuId"`
	CourseName string    `gorm:"not null"`
	BaseCourse float64   `gorm:"not null"`
	TestScore  float64   `gorm:"not null"`
	FinalScore float64   `gorm:"not null"`
	AddDate    time.Time `gorm:"not null"`
	Note       string
}

func (Student) TableName() string {
	return "student_info"
}

func (Mark) TableName() string {
	return "stu_mark_info"
}

var DB *gorm.DB

func InitDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:root@/sms?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		DB = db
		return
	}
	panic(err.Error())
}

func (s *Student) Insert() error {
	return DB.Create(s).Error
}

func (s *Student) Update() error {
	return DB.Model(s).Update(map[string]interface{}{
		"profession": s.Profession,
	}).Error
}

func (s *Student) DeleteByStuId(stuId string) error {
	i, e := strconv.Atoi(stuId)
	if e != nil {
		return e
	}
	return DB.Delete(s, "stuid = ?", i).Error
}

func GetStudentByStuId(stuId string) (*Student, error) {
	var s Student
	i, e := strconv.Atoi(stuId)
	if e != nil {
		return nil, e
	}
	e = DB.Where("stuid = ?", i).First(&s).Error
	return &s, e
}

func GetAllStudents() (students []*Student, err error) {
	err = DB.Find(&students).Error
	return
}

func DeleteStudent(stuId string) error {
	student, e := GetStudentByStuId(stuId)
	if e != nil {
		fmt.Println(e)
		return e
	}
	DB.Delete(&student)
	return nil
}
