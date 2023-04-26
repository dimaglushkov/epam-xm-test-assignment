package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/ports"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type HTTPHandler struct {
	port           string
	companyService ports.CompanyService
	router         *gin.Engine
	signKey        string
}

func NewHTTPHandler(port, mode, signKey string, companyService ports.CompanyService) *HTTPHandler {
	if mode != "" {
		gin.SetMode(mode)
	}

	handler := new(HTTPHandler)
	handler.companyService = companyService
	handler.port = port
	handler.signKey = signKey

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/companies/:id", handler.getCompany)
	router.POST("/companies", handler.AuthCheckMiddleware(), handler.createCompany)
	router.PATCH("/companies/:id", handler.AuthCheckMiddleware(), handler.updateCompany)
	router.DELETE("/companies/:id", handler.AuthCheckMiddleware(), handler.deleteCompany)

	handler.router = router

	return handler
}

func (h *HTTPHandler) AuthCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(h.signKey), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

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

	c.JSON(http.StatusOK, company)
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

	c.JSON(http.StatusOK, company)
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

	c.JSON(http.StatusOK, company)
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
