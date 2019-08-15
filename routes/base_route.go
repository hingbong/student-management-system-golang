package routes

import (
	"github.com/hingbong/student-management-system-golang/routes/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func SetupRouter(s *http.Server) (e *echo.Echo) {
	e = echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Static("./static"))

	e.POST("/student_op", service.AddStudentPost)
	e.GET("/student_op", service.AllStudentsGet)
	e.GET("/student_op/:id", service.OneStudentsGet)
	e.PUT("/student_op", service.ModifyProfessionPut)
	e.DELETE("/student_op/:id", service.DeleteStudent)

	e.POST("/mark_op", service.AddMarkPost)
	e.GET("/mark_op", service.GetAllMarksGet)
	e.DELETE("/mark_op/:id", service.DeleteMarkDelete)
	e.Logger.Fatal(e.StartServer(s))
	return
}
