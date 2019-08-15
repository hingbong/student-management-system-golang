package entity

import (
	"fmt"
	"github.com/hingbong/student-management-system-golang/models"
	"strconv"
	"time"
)

type Student struct {
	StuId      uint      `gorm:"primary_key;column:stuid;AUTO_INCREMENT"json:"stuId"`
	StuName    string    `gorm:"not null;column:stuname"json:"stuName"`
	Profession int       `gorm:"not null"json:"profession"`
	Date       time.Time `gorm:"not null"json:"date"`
	Sex        int       `gorm:"not null"json:"sex"`
	IdNum      string    `gorm:"not null"json:"idNum"`
	TotalScore int       `gorm:"not null"json:"totalScore"`
	Note       string    `json:"note"`
}

func (Student) TableName() string {
	return "student_info"
}

func (s *Student) Insert() error {
	return models.DB.Create(s).Error
}

func (s *Student) Update() error {
	return models.DB.Model(s).Update(map[string]interface{}{
		"profession": s.Profession,
	}).Error
}

func UpdateTotalScoreById(stuId uint, totalScore float64) error {
	return models.DB.Model(new(Student)).Where("stuid = ?", stuId).Update(map[string]interface{}{
		"totalScore": totalScore,
	}).Error
}

func (s *Student) DeleteByStuId(stuId string) error {
	i, e := strconv.Atoi(stuId)
	if e != nil {
		return e
	}
	return models.DB.Delete(s, "stuid = ?", i).Error
}

func GetStudentByStuId(stuId string) (*Student, error) {
	var s Student
	i, e := strconv.Atoi(stuId)
	if e != nil {
		return nil, e
	}
	e = models.DB.First(&s, i).Error
	return &s, e
}

func GetAllStudents(name, profession string) (students []*Student, err error) {
	db := models.DB
	if name != "" {
		db = db.Where("stuname LIKE ?", "%"+name+"%")
	}
	if profession != "" {
		i, err := strconv.Atoi(profession)
		if err == nil {
			db = db.Where("profession = ?", i)
		}
	}
	err = db.Find(&students).Error
	return
}

func DeleteStudent(stuId string) error {
	i, e := strconv.Atoi(stuId)
	if e != nil {
		fmt.Println(e)
		return e
	}
	tx := models.DB.Begin()
	tx.Where("stuid = ?", i).Delete(&Student{})
	marks := GetMarksByStuId(uint(i))
	for _, v := range marks {
		e := DeleteMark(strconv.Itoa(int(v.StuId)))
		if e != nil {
			tx.Rollback()
			return e
		}
	}
	tx.Commit()
	return nil
}
