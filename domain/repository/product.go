package repository

import "github.com/jdpadillaac/go-waitgroups-example/domain/entity"

type Product interface {
	Save(model entity.Product) error
	GetAll(model entity.Product) []entity.Product
}
