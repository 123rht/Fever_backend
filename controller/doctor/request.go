package doctor

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func getPageInfo(c *gin.Context) (int, int) {
	//获取分页参数
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		page int
		size int
		err  error
	)

	page, err = strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}
	return page, size
}
