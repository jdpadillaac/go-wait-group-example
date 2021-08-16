package repository

import "github.com/jdpadillaac/go-waitgroups-example/domain/entity"

type ProductCategory interface {
	Save(model entity.ProductCategory) (entity.ProductCategory, error)
	FindByID(ID int) (entity.ProductCategory, error)
}
