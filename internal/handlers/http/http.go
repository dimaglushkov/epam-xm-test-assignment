package http

import (
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	port   string
	svc    ports.CompanyServicePort
	router *gin.Engine
}

func New(port string, mode string, svc ports.CompanyServicePort) *Handler {
	if mode != "" {
		gin.SetMode(mode)
	}

	handler := new(Handler)
	handler.svc = svc
	handler.port = port
	router := gin.Default()

	router.GET("/company/:id", handler.getCompany)
	router.POST("/company", handler.createCompany).Use(AuthCheckMiddleware())

	handler.router = router
	return handler
}

func AuthCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("todo: auth middleware")
		c.Next()
	}
}

func (h *Handler) Run() error {
	return h.router.Run(":" + h.port)
}

func (h *Handler) getCompany(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
	}

}

func (h *Handler) createCompany(c *gin.Context) {
	reqBody := new(domain.Company)
	if err := c.Bind(reqBody); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if validationErrors := reqBody.Validate(); len(validationErrors) > 0 {
		messages := make([]string, len(validationErrors))
		for i, err := range validationErrors {
			messages[i] = err.Error()
		}
		errorResponse(c, http.StatusBadRequest, strings.Join(messages, ";\n"))
		return
	}

	c.JSON(200, reqBody)
}
