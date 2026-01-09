package dtos

type SucursalRequest struct {
	Nombre       string  `json:"nombre" binding:"required"`
	Ciudad       string  `json:"ciudad"`
	SuperficieM2 float64 `json:"superficieM2"`
}
