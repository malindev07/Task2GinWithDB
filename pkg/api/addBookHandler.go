package api

import (
	db "Task2DbWithGin/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddBookHadnler(c *gin.Context) {

	// биндим занчения из post  в json формат и передаем в структуру book
	//if err := c.BindJSON(&book); err != nil {
	//	fmt.Fprint(c.Writer, "Возникла ошибка")
	//	c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	//}

	c.JSON(http.StatusOK, db.AddBook(c))

}
