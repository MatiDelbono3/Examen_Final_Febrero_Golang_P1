package dtos

type SucursalResponse struct {
	Id     string  `json:"id" binding:"required"`
	Nombre float64 `json:"nombre"`
}
