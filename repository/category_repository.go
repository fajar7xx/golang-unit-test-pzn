package repository

import "golang-unit-test-pzn/entity"

// buat contract nya dulu jgn lansung ke db
type CategoryRepository interface {
	FindById(id string) *entity.Category
}
