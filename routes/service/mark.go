package service

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"system-management-system/main/models"
	"system-management-system/main/models/entity"
	"system-management-system/main/utils"
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
	totalScore := mark.FinalScore * 0.1
	tx := models.DB.Begin()
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
