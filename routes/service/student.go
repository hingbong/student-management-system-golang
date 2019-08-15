package service

import (
	"fmt"
	"github.com/afanti-com/utils-go/idCardNo"
	"github.com/hingbong/student-management-system-golang/models/entity"
	"github.com/hingbong/student-management-system-golang/utils"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func AddStudentPost(c echo.Context) error {
	student := new(entity.Student)
	param := c.FormValue("date")
	err := c.Bind(student)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage(err.Error()))
	}

	parse, err := time.Parse("2006-01-02", param)
	if err != nil {
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage("日期格式错误"))
	}
	student.Date = parse

	if student.StuName == "" {
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage("学生姓名错误"))
	}

	if student.Profession < 1 || student.Profession > 3 {
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage("学生专业错误"))
	}

	if student.Sex < 0 || student.Sex > 1 {
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage("学生性别错误"))
	}

	// check id number
	students, err := entity.GetAllStudents(utils.EmptyString, utils.EmptyString)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage(err.Error()))
	}
	for _, v := range students {
		if v.IdNum == student.IdNum {
			return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage("身份证错误"))
		}
		if ok := idCardNo.Verification(student.IdNum); !ok {
			return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage("身份证错误"))
		}
	}

	// insert
	err = student.Insert()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.SuccessJson())
}

func AllStudentsGet(c echo.Context) error {
	name := c.FormValue("stuName")
	profession := c.FormValue("profession")
	students, err := entity.GetAllStudents(name, profession)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.SuccessJsonWithData(students))
}

func OneStudentsGet(c echo.Context) error {
	id := c.Param("id")
	student, err := entity.GetStudentByStuId(id)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.SuccessJsonWithData(student))
}

func ModifyProfessionPut(c echo.Context) error {
	student := new(entity.Student)
	err := c.Bind(student)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage(err.Error()))
	}
	err = student.Update()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.SuccessJson())
}

func DeleteStudent(c echo.Context) error {
	id := c.Param("id")
	err := entity.DeleteStudent(id)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage(err.Error()))
	}
	return c.JSON(http.StatusOK, utils.SuccessJson())
}
