package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/paytm-temp/cms-backend/pkg/api/v1/services"
    "github.com/paytm-temp/cms-backend/pkg/models/case"
)

type CaseHandler struct {
    caseService *services.CaseService
}

func NewCaseHandler() *CaseHandler {
    return &CaseHandler{
        caseService: services.NewCaseService(),
    }
}

func (h *CaseHandler) GetAllCases(c *gin.Context) {
    cases, err := h.caseService.GetAllCases()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"cases": cases})
}

func (h *CaseHandler) GetCaseByID(c *gin.Context) {
    id := c.Param("id")
    caseItem, err := h.caseService.GetCaseByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Case not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"case": caseItem})
}

func (h *CaseHandler) CreateCase(c *gin.Context) {
    var newCase cases.Case
    if err := c.ShouldBindJSON(&newCase); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdCase, err := h.caseService.CreateCase(newCase)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"case": createdCase})
}

func (h *CaseHandler) UpdateCase(c *gin.Context) {
    id := c.Param("id")
    var updatedCase cases.Case
    if err := c.ShouldBindJSON(&updatedCase); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result, err := h.caseService.UpdateCase(id, updatedCase)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Case not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"case": result})
}

func (h *CaseHandler) DeleteCase(c *gin.Context) {
    id := c.Param("id")
    if err := h.caseService.DeleteCase(id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Case not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Case deleted successfully"})
}

func (h *CaseHandler) GetStats(c *gin.Context) {
    stats, err := h.caseService.GetStats()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"stats": stats})
}

func RegisterRoutes(r *gin.Engine) {
    handler := NewCaseHandler()
    v1 := r.Group("/api/v1")
    {
        cases := v1.Group("/cases")
        {
            cases.GET("", handler.GetAllCases)
            cases.GET("/:id", handler.GetCaseByID)
            cases.POST("", handler.CreateCase)
            cases.PUT("/:id", handler.UpdateCase)
            cases.DELETE("/:id", handler.DeleteCase)
            cases.GET("/stats", handler.GetStats)
        }
    }
}
