package entity

import (
	"fmt"
	"github.com/hingbong/student-management-system-golang/models"
	"strconv"
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

type MarkVO struct {
	MarkId     uint      `gorm:"primary_key;column:markid;AUTO_INCREMENT"json:"markId"`
	StuId      uint      `gorm:"not null;column:stuId"json:"stuId"`
	CourseName string    `gorm:"not null"json:"courseName"`
	BaseScore  float64   `gorm:"not null"json:"baseScore"`
	TestScore  float64   `gorm:"not null"json:"testScore"`
	FinalScore float64   `gorm:"not null"json:"finalScore"`
	AddDate    time.Time `gorm:"not null"json:"addDate"`
	Note       string    `json:"note"`
	StuName    string    `gorm:"not null;column:stuname"json:"stuName"`
}

func GetAllMarkWithStudentsName(name, addDate string) ([]*MarkVO, error) {
	db := models.DB
	if name != "" {
		db = db.Where("student_info.stuname LIKE ?", "%"+name+"%")
	}
	if addDate != "" {
		db = db.Where("DATE(stu_mark_info.add_date) = ?", addDate)
	}
	rows, err := db.Table("stu_mark_info").Select(`
        student_info.stuname,
		stu_mark_info.markid,
		stu_mark_info.stuId,
		stu_mark_info.course_name,
		stu_mark_info.base_score,
		stu_mark_info.test_score,
		stu_mark_info.final_score,
		stu_mark_info.add_date,
		stu_mark_info.note
	`).Joins(`JOIN student_info ON stu_mark_info.stuId = student_info.stuid`).Rows()
	if err != nil {
		return nil, err
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			return
		}
	}()
	vos := make([]*MarkVO, 0)
	for rows.Next() {
		vo := MarkVO{}
		err := rows.Scan(&vo.StuName, &vo.MarkId, &vo.StuId, &vo.CourseName, &vo.BaseScore, &vo.TestScore, &vo.FinalScore, &vo.AddDate, &vo.Note)
		if err != nil {
			return vos, err
		}
		vos = append(vos, &vo)
	}
	return vos, nil
}

func GetMarksByStuId(stuId uint) (marks []*Mark) {
	models.DB.Where("stuId = ?", stuId).Find(&marks)
	return
}

func DeleteMark(markId string) error {
	i, e := strconv.Atoi(markId)
	if e != nil {
		fmt.Println(e)
		return e
	}
	models.DB.Where("markid = ?", i).Delete(&Mark{})
	return nil
}
