package persistence

import (
	"github.com/negiseijin/clean-architectur-swagger/domain/repository"
	"gorm.io/gorm"
)

type DBRepository struct {
	DB *gorm.DB
}

// Create implements repository.DBRepository
func (repo *DBRepository) Create(model interface{}, item interface{}) error {
	db := repo.DB.Debug().Model(model)
	db.Create(item)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// Delete implements repository.DBRepository
func (repo *DBRepository) Delete(model interface{}, item interface{}) error {
	db := repo.DB.Debug().Model(model)
	db.Delete(item)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// Find implements repository.DBRepository
func (repo *DBRepository) Find(model interface{}, params interface{}, item interface{}) error {
	db := repo.DB.Debug().Model(model)
	db.Where(params).Find(item)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// Migrate implements repository.DBRepository
func (repo *DBRepository) Migrate(item interface{}) error {
	err := repo.DB.Debug().AutoMigrate(item)
	if err != nil {
		return err
	}
	return nil
}

// Read implements repository.DBRepository
func (repo *DBRepository) Read(model interface{}) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	db := repo.DB.Debug().Model(model)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}
	return results, nil
}

// Update implements repository.DBRepository
func (repo *DBRepository) Update(model interface{}, item interface{}) error {
	db := repo.DB.Debug().Model(model)
	db.Updates(item)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func NewDBRepository(db *gorm.DB) repository.DBRepository {
	return &DBRepository{DB: db}
}
