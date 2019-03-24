package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

type httpCode struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		db, err := gorm.Open("postgres", "host=pgd_db port=5432 user=db_user dbname=db_name password=mypassword sslmode=disable")
		if err != nil {
			resCode(c, http.StatusInternalServerError, err)
		} else {
			resCode(c, http.StatusOK, db)
		}
		defer db.Close()
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func resCode(ctx *gin.Context, code int, data ...interface{}) {
	er := httpCode{
		Status:  code,
		Message: http.StatusText(code),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, er)
}
