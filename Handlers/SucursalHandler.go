package handlers

import (
	"net/http"
	"strconv"

	dtos "examen_final_febrero_golang_P1/Dtos"
	service "examen_final_febrero_golang_P1/Services"

	"github.com/gin-gonic/gin"
)

type SucursalHandler struct {
	service service.SucursalService
}

func NewSucursalHandler(service service.SucursalService) *SucursalHandler {
	return &SucursalHandler{
		service: service,
	}
}

// ==================== CREAR SUCURSAL ====================
func (h *SucursalHandler) Crear(c *gin.Context) {
	var req dtos.SucursalRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.Crear(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ==================== LISTAR SUCURSALES ====================
func (h *SucursalHandler) Listar(c *gin.Context) {
	resp, err := h.service.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ==================== CÁLCULOS ESTADÍSTICOS ====================
func (h *SucursalHandler) Calculos(c *gin.Context) {
	var req dtos.EstadisticasRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado, err := h.service.Calculos(req.Valores, req.Tipo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"resultado": resultado})
}

// ==================== PROYECCIÓN ====================
func (h *SucursalHandler) ObtenerTablaProyeccion(c *gin.Context) {
	montoInicialStr := c.Query("montoInicial")
	tasaCrecimientoStr := c.Query("tasaCrecimiento")
	aniosStr := c.Query("anios")

	montoInicial, err := strconv.ParseFloat(montoInicialStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "montoInicial inválido"})
		return
	}

	tasaCrecimiento, err := strconv.Atoi(tasaCrecimientoStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tasaCrecimiento inválida"})
		return
	}

	anios, err := strconv.Atoi(aniosStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "anios inválidos"})
		return
	}

	req := dtos.SucursalItemRequest{
		MontoInicial:    montoInicial,
		TasaCrecimiento: tasaCrecimiento,
		Anios:           anios,
	}

	resultado, err := h.service.ObtenerTablaProyeccion(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultado)
}
