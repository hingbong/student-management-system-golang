package routes

import (
	"fmt"
	"github.com/afanti-com/utils-go/idCardNo"
	"github.com/labstack/echo"
	"net/http"
	"system-management-system/main/models"
	"time"
)

const EMPTY_STRING = ""

func addStudentPost(c echo.Context) error {
	student := new(models.Student)
	param := c.FormValue("date")
	err := c.Bind(student)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, errorJsonWithMessage(err.Error()))
	}

	parse, err := time.Parse("2006-01-02", param)
	if err != nil {
		return c.JSON(http.StatusOK, errorJsonWithMessage("日期格式错误"))
	}
	student.Date = parse

	if student.StuName == "" {
		return c.JSON(http.StatusOK, errorJsonWithMessage("学生姓名错误"))
	}

	if student.Profession < 1 || student.Profession > 3 {
		return c.JSON(http.StatusOK, errorJsonWithMessage("学生专业错误"))
	}

	if student.Sex < 0 || student.Sex > 1 {
		return c.JSON(http.StatusOK, errorJsonWithMessage("学生性别错误"))
	}

	// check id number
	students, err := models.GetAllStudents(EMPTY_STRING, EMPTY_STRING)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, errorJsonWithMessage(err.Error()))
	}
	for _, v := range students {
		if v.IdNum == student.IdNum {
			return c.JSON(http.StatusOK, errorJsonWithMessage("身份证错误"))
		}
		if ok := idCardNo.Verification(student.IdNum); !ok {
			return c.JSON(http.StatusOK, errorJsonWithMessage("身份证错误"))
		}
	}

	// insert
	err = student.Insert()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, errorJsonWithMessage(err.Error()))
	}
	return c.JSON(http.StatusOK, successJson())
}

func allStudentsGet(c echo.Context) error {
	name := c.FormValue("stuName")
	profession := c.FormValue("profession")
	students, err := models.GetAllStudents(name, profession)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, errorJsonWithMessage(err.Error()))
	}
	return c.JSON(http.StatusOK, successJsonWithData(students))
}

func oneStudentsGet(c echo.Context) error {
	id := c.Param("id")
	student, err := models.GetStudentByStuId(id)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, errorJsonWithMessage(err.Error()))
	}
	return c.JSON(http.StatusOK, successJsonWithData(student))
}

func modifyProfessionPut(c echo.Context) error {
	student := new(models.Student)
	err := c.Bind(student)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, errorJsonWithMessage(err.Error()))
	}
	err = student.Update()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, errorJsonWithMessage(err.Error()))
	}
	return c.JSON(http.StatusOK, successJson())
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	err := models.DeleteStudent(id)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, errorJsonWithMessage(err.Error()))
	}
	return c.JSON(http.StatusOK, successJson())
}
