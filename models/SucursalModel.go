package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sucursal struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Nombre   string             `bson:"nombre"`
	CreadoEn time.Time          `bson:"creadoEn"`
}
