package response

import (
	"github.com/gin-gonic/gin"
)

// ApiResponse is the standard response wrapper for all endpoints
type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // lebih flexible saat return data apapun disimpan ke Data tipe interface
}

type PaginationResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
	HasNext    bool        `json:"has_next"`
	HasPrev    bool        `json:"has_prev"`
}

// SendResponse sends a standard success response
func SendResponse(c *gin.Context, status int, msg string, data interface{}) {
	resp := ApiResponse{
		Success: true,
		Message: msg,
		Data:    data,
	}
	c.JSON(status, resp)
}

// SendError sends a standard error response
func SendError(c *gin.Context, status int, msg string) {
	resp := ApiResponse{
		Success: false,
		Message: msg,
		Data:    nil,
	}
	c.JSON(status, resp)
}
