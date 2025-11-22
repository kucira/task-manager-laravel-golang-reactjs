package models

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CursorPagination struct {
	Limit      int    `json:"limit"`
	NextCursor string `json:"nextCursor"`
	HasMore    bool   `json:"hasMore"`
	Search     string `json:"search"`
	Sort       string `json:"sort"`
}

func CursorPaginate(c *gin.Context) (CursorPagination, func(db *gorm.DB) *gorm.DB) {
	cursorParam := c.Query("cursor")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := strings.ToLower(c.DefaultQuery("sort", "asc"))
	search := strings.TrimSpace(c.Query("search"))

	if limit <= 0 {
		limit = 10
	}

	p := CursorPagination{
		Limit:  limit,
		Search: search,
		Sort:   sort,
	}

	scope := func(db *gorm.DB) *gorm.DB {
		// Apply search if exists
		if p.Search != "" {
			db = db.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(p.Search)+"%")
		}

		// Cursor logic
		if cursorParam != "" {
			cursor, err := strconv.Atoi(cursorParam)
			if err == nil {
				if p.Sort == "desc" {
					db = db.Where("id < ?", cursor)
				} else {
					db = db.Where("id > ?", cursor)
				}
			}
		}

		// Sorting
		if p.Sort == "desc" {
			db = db.Order("id DESC")
		} else {
			db = db.Order("id ASC")
		}

		return db.Limit(p.Limit)
	}

	return p, scope
}
