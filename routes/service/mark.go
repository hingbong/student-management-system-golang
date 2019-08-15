package service

import (
	"fmt"
	"github.com/hingbong/student-management-system-golang/models"
	"github.com/hingbong/student-management-system-golang/models/entity"
	"github.com/hingbong/student-management-system-golang/utils"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func AddMarkPost(c echo.Context) error {
	mark := new(entity.Mark)
	e := c.Bind(mark)
	if e != nil {
		fmt.Println(e)
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage(e.Error()))
	}
	mark.AddDate = time.Now()
	mark.FinalScore = mark.BaseScore*0.4 + mark.TestScore*0.6

	tx := models.DB.Begin()

	marks := entity.GetMarksByStuId(mark.StuId)
	var totalScore float64
	for _, v := range marks {
		totalScore += v.FinalScore
	}
	totalScore += mark.FinalScore
	totalScore *= 0.1

	e = entity.UpdateTotalScoreById(mark.StuId, totalScore)
	if e != nil {
		tx.Rollback()
		fmt.Println(e)
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage(e.Error()))
	}
	e = mark.Insert()
	if e != nil {
		tx.Rollback()
		fmt.Println(e)
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage(e.Error()))
	}
	tx.Commit()
	return c.JSON(http.StatusOK, utils.SuccessJson())
}

func GetAllMarksGet(c echo.Context) error {
	name := c.FormValue("stuName")
	date := c.FormValue("addDate")
	vos, e := entity.GetAllMarkWithStudentsName(name, date)
	if e != nil {
		return c.JSON(http.StatusOK, utils.ErrorJsonWithMessage(e.Error()))
	}
	return c.JSON(http.StatusOK, utils.SuccessJsonWithData(vos))
}

func DeleteMarkDelete(c echo.Context) error {
	id := c.Param("id")
	e := entity.DeleteMark(id)
	if e != nil {
		return c.JSON(http.StatusOK, e.Error())
	}
	return c.JSON(http.StatusOK, utils.SuccessJson())
}
