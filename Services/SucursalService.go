package Service

import (
	"context"
	"errors"
	"math"
	"sort"
	"time"

	dtos "examen_final_febrero_golang_P1/Dtos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type sucursalService interface {
	Crear(req dtos.SucursalRequest) (dtos.SucursalResponse, error)
	Listar() ([]dtos.SucursalResponse, error)
	Calculos(valores []float64, tipo string) (float64, error)
	ObtenerTablaProyeccion(dtos.SucursalItemRequest) ([]dtos.SucursalItemResponse, error)
}

type SucursalService struct {
	collection *mongo.Collection
}

func NewEntidadService(collection *mongo.Collection) *SucursalService {
	return &SucursalService{
		collection: collection,
	}

}

const (
	CalculoPromedio = "PROMEDIO"
	CalculoVarianza = "VARIANZA"
	CalculoDesvio   = "DESVIO"
	CalculoMediana  = "MEDIANA"
	CalculoMaximo   = "MAXIMO"
	CalculoMinimo   = "MINIMO"
)

/*
Calculos
Función GENERAL para cualquier tipo de cálculo numérico
*/
func (service *SucursalService) Calculos(valores []float64, tipo string) (float64, error) {

	// Validaciones generales
	if len(valores) == 0 {
		return 0, errors.New("la lista de valores no puede estar vacía")
	}

	for _, v := range valores {
		if v < 0 {
			return 0, errors.New("los valores no pueden ser negativos")
		}
	}

	switch tipo {

	case CalculoPromedio:
		suma := 0.0
		for _, v := range valores {
			suma += v
		}
		return suma / float64(len(valores)), nil

	case CalculoVarianza:
		media, _ := service.Calculos(valores, CalculoPromedio)

		var suma float64
		for _, v := range valores {
			d := v - media
			suma += d * d
		}
		return suma / float64(len(valores)), nil

	case CalculoDesvio:
		varianza, _ := service.Calculos(valores, CalculoVarianza)
		return math.Sqrt(varianza), nil

	case CalculoMediana:
		ordenados := make([]float64, len(valores))
		copy(ordenados, valores)
		sort.Float64s(ordenados)

		mid := len(ordenados) / 2
		if len(ordenados)%2 == 0 {
			return (ordenados[mid-1] + ordenados[mid]) / 2, nil
		}
		return ordenados[mid], nil

	case CalculoMaximo:
		max := valores[0]
		for _, v := range valores {
			if v > max {
				max = v
			}
		}
		return max, nil

	case CalculoMinimo:
		min := valores[0]
		for _, v := range valores {
			if v < min {
				min = v
			}
		}
		return min, nil

	default:
		return 0, errors.New("tipo de cálculo no soportado")
	}
}

func (service *SucursalService) Crear(req dtos.SucursalRequest) (dtos.SucursalResponse, error) {

	// Validaciones
	if req.Nombre == "" {
		return dtos.SucursalResponse{}, errors.New("el nombre es obligatorio")
	}

	if req.Ciudad == "" || req.Ciudad == "" {
		return dtos.SucursalResponse{}, errors.New("latitud inválida")
	}

	// Documento a persistir
	doc := bson.M{
		"Nombre":       req.Nombre,
		"Ciudad":       req.Ciudad,
		"SuperficieM2": req.SuperficieM2,
	}

	result, err := service.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return dtos.SucursalResponse{}, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return dtos.SucursalResponse{
		Id:       id.Hex(),
		Nombre:   req.Nombre,
		CreadoEn: time.Now(),
	}, nil
}
func (s *SucursalService) Listar() ([]dtos.SucursalResponse, error) {

	cursor, err := s.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var resultado []dtos.SucursalResponse

	for cursor.Next(context.Background()) {

		var doc struct {
			ID       primitive.ObjectID `bson:"_id"`
			Nombre   string             `bson:"nombre"`
			CreadoEn time.Time          `bson:"creado_en"`
		}

		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}

		resultado = append(resultado, dtos.SucursalResponse{
			Id:       doc.ID.Hex(),
			Nombre:   doc.Nombre,
			CreadoEn: doc.CreadoEn,
		})
	}

	return resultado, nil
}
func (f *SucursalService) ObtenerTablaProyeccion(proyeccion dtos.SucursalItemRequest) ([]dtos.SucursalItemResponse, error) {
	if proyeccion.MontoInicial <= 0 || proyeccion.TasaCrecimiento <= 0 || proyeccion.Anios <= 0 {
		return nil, errors.New("Monto inicial, tasa de crecimiento y años deben ser mayores a 0")
	}
	var tabla []dtos.SucursalItemResponse
	monto := proyeccion.MontoInicial
	for i := 1; i <= proyeccion.Anios; i++ {
		crecimiento := monto * float64(proyeccion.TasaCrecimiento/100)
		montoFinal := monto + crecimiento
		tabla = append(tabla, dtos.SucursalItemResponse{
			Anios: i,
			Monto: montoFinal,
		})
		monto = float64(montoFinal)
	}

	return tabla, nil
}
