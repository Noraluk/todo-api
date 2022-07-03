package base

import (
	"todo-api/pkg/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository[T any] struct {
	db           *gorm.DB
	associations []Association
}

type Association struct {
	Type string
	Name string
}

func (repo *Repository[T]) Init(associations []Association) {
	var dbs = database.DB
	for _, association := range associations {
		if association.Type == "preload" {
			dbs = dbs.Preload(association.Name)
		} else if association.Type == "joins" {
			dbs = dbs.Joins(association.Name)
		}
	}
	repo.db = dbs
	repo.associations = associations
}

func (repo *Repository[T]) Migration(model *T) {
	repo.db.AutoMigrate(model)
}

func (repo *Repository[T]) reset() {
	repo.Init(repo.associations)
}

func (repo *Repository[T]) Where(query any, value ...any) *gorm.DB {
	repo.db = repo.db.Where(query, value)
	return repo.db
}

func (repo *Repository[T]) Or(query any, value ...any) *gorm.DB {
	repo.db = repo.db.Or(query, value)
	return repo.db
}

func (repo *Repository[T]) PreloadCondition(associationName string, conditions ...any) *gorm.DB {
	var dbs = database.DB
	for _, association := range repo.associations {
		if association.Type == "preload" {
			if associationName == association.Name {
				dbs = dbs.Preload(association.Name, conditions...)
			} else {
				dbs = dbs.Preload(association.Name)
			}
		} else if association.Type == "joins" {
			if associationName == association.Name {
				dbs = dbs.Joins(association.Name, conditions...)
			} else {
				dbs = dbs.Joins(association.Name)
			}
		}
	}
	repo.db = dbs
	return repo.db
}

func (repo *Repository[T]) Order(field string, orderType string) *gorm.DB {
	repo.db = repo.db.Order(field + " " + orderType)
	return repo.db
}

func (repo *Repository[T]) Between(query any, value1 any, value2 any) *gorm.DB {
	repo.db = repo.db.Where(query, value1, value2)
	return repo.db
}

func (repo *Repository[T]) FindOneById(id int) (T, *gorm.DB) {
	var models T
	db := repo.db.First(&models, id)
	repo.reset()
	return models, db
}

func (repo *Repository[T]) FindOne(cond any) (T, *gorm.DB) {
	var model T
	db := repo.db.Where(cond).First(&model)
	repo.reset()
	return model, db
}

func (repo *Repository[T]) FindOrCreate(model *T, cond any) (*T, *gorm.DB) {
	db := repo.db.Where(cond).FirstOrCreate(model)
	repo.reset()
	return model, db
}

func (repo *Repository[T]) Find(cond any) ([]T, *gorm.DB) {
	var models []T
	db := repo.db.Where(cond).Find(&models)
	repo.reset()
	return models, db
}

func (repo *Repository[T]) FindPaginate(cond any, page int, limit int) ([]T, *gorm.DB) {
	var models []T
	offset := (page - 1) * limit
	db := repo.db.Offset(offset).Limit(limit).Where(cond).Find(&models)
	repo.reset()
	return models, db
}

func (repo *Repository[T]) Count() (int64, *gorm.DB) {
	var model T
	count := int64(0)
	db := repo.db.Model(&model).Count(&count)
	return count, db
}

func (repo *Repository[T]) Insert(model *T) (*T, *gorm.DB) {
	db := repo.db.Create(model)
	repo.reset()
	return model, db
}

func (repo *Repository[T]) Inserts(models *[]T) (*[]T, *gorm.DB) {
	db := repo.db.Create(models)
	repo.reset()
	return models, db
}

func (repo *Repository[T]) Update(model *T, cond any) (*T, *gorm.DB) {
	db := repo.db.Model(model).Where(cond).Updates(model)
	repo.reset()
	return model, db
}

func (repo *Repository[T]) Upsert(model *T, col string, updateCols []string) (*T, *gorm.DB) {
	db := repo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: col}},
		DoUpdates: clause.AssignmentColumns(updateCols),
	}).Create(model)
	repo.reset()
	return model, db
}

func (repo *Repository[T]) Delete(model *T) (*T, *gorm.DB) {
	db := repo.db.Delete(model)
	repo.reset()
	return model, db
}
