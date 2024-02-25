package api

import (
	db "Task2DbWithGin/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DeleteValueHandler(c *gin.Context) {

	bookName := c.Param("bookName")

	fmt.Fprint(c.Writer, db.DeleteValue(bookName))

}
