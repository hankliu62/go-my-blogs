package utils

import (
	"math"

	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
	"hankliu.com.cn/go-my-blog/constants"
	"hankliu.com.cn/go-my-blog/types"
)

// GetResultsMeta get query list result meta property
func GetResultsMeta(page, pageSize, total int) types.StringMap {
	pageCount := float64(total) / float64(pageSize)
	meta := types.StringMap{
		"currentPage": page,
		"perPage":     pageSize,
		"totalCount":  total,
		"pageCount":   math.Ceil(pageCount),
	}

	return meta
}

// ParsePagingCondition parse pagination from gin Context
func ParsePagingCondition(c *gin.Context) (uint32, uint32) {
	page := constants.DEFAULT_PAGE_INDEX
	pageSize := constants.DEFAULT_PAGE_SIZE

	if queryPage, pageOk := c.GetQuery("page"); pageOk {
		if ppage, err := strconv.ParseInt(queryPage, 10, 32); err == nil {
			page = uint32(ppage)
			if queryPageSize, pageSizeOk := c.GetQuery("pageSize"); pageSizeOk {
				if ppageSize, errSize := strconv.ParseInt(queryPageSize, 10, 32); errSize == nil {
					pageSize = uint32(ppageSize)
				}
			}
		}
	}

	return page, pageSize
}
