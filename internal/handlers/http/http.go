package http

import (
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
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

	router.GET("/company/:id", handler.getCompany).Use(Validate())

	handler.router = router
	return handler
}

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		log.Print(latency)
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
