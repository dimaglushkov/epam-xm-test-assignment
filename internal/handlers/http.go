package handlers

import (
	"errors"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type HTTPHandler struct {
	port           string
	companyService ports.CompanyService
	router         *gin.Engine
}

func NewHTTPHandler(port string, mode string, companyService ports.CompanyService) *HTTPHandler {
	if mode != "" {
		gin.SetMode(mode)
	}

	handler := new(HTTPHandler)
	handler.companyService = companyService
	handler.port = port

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/companies/:id", handler.getCompany)
	router.POST("/companies", AuthCheckMiddleware(), handler.createCompany)
	router.PATCH("/companies/:id", AuthCheckMiddleware(), handler.updateCompany)
	router.DELETE("/companies/:id", AuthCheckMiddleware(), handler.deleteCompany)

	handler.router = router
	return handler
}

func AuthCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("todo: auth middleware")
		c.Next()
	}
}

func (h *HTTPHandler) Run() error {
	return h.router.Run(":" + h.port)
}

func (h *HTTPHandler) getCompany(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		errorResponse(c, err)
		return
	}

	company, err := h.companyService.Get(c, id)
	if err != nil {
		errorResponse(c, err)
		return
	}

	c.JSON(200, company)
}

func (h *HTTPHandler) createCompany(c *gin.Context) {
	company := new(domain.Company)
	if err := c.ShouldBind(company); err != nil {
		errorResponse(c, err)
		return
	}

	if err := h.companyService.Create(c, company); err != nil {
		errorResponse(c, err)
		return
	}

	c.JSON(200, company)
}

func (h *HTTPHandler) updateCompany(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		errorResponse(c, err)
		return
	}

	fieldsToUpdate := make(map[string]any)
	if err := c.ShouldBind(&fieldsToUpdate); err != nil {
		errorResponse(c, err)
		return
	}

	if err := h.companyService.Update(c, id, fieldsToUpdate); err != nil {
		errorResponse(c, err)
		return
	}

	company, err := h.companyService.Get(c, id)
	if err != nil {
		errorResponse(c, err)
		return
	}
	c.JSON(200, company)
}

func (h *HTTPHandler) deleteCompany(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		errorResponse(c, err)
		return
	}
	if err := h.companyService.Delete(c, id); err != nil {
		errorResponse(c, err)
		return
	}
}

func errorResponse(c *gin.Context, err error) {
	code := http.StatusBadRequest
	if errors.Is(err, domain.ErrInternalServer) {
		code = http.StatusInternalServerError
	}
	var companyNotFoundErr *domain.CompanyNotFoundError
	if errors.As(err, &companyNotFoundErr) {
		code = http.StatusNotFound
	}

	c.AbortWithStatusJSON(code, gin.H{
		"error": err.Error(),
	})
}
