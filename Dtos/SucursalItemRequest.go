package dtos

type SucursalItemRequest struct {
	MontoInicial    float64 `json:"montoInicial"`
	TasaCrecimiento int     `json:"tasaCrecimiento"`
	Anios           int     `json:"anios"`
}
