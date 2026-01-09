package dtos

import "time"

type SucursalResponse struct {
	Id       string    `json:"id" binding:"required"`
	Nombre   string    `json:"nombre"`
	CreadoEn time.Time `json:"creadoEn"`
}
