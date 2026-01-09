package dtos

type EstadisticasRequest struct {
	Tipo    string    `json:"tipo"`
	Valores []float64 `json:"valores"`
}
