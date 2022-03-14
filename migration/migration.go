package migration

import (
	"github.com/negiseijin/clean-architectur-swagger/domain/model"
	"github.com/negiseijin/clean-architectur-swagger/domain/repository"
	"github.com/negiseijin/clean-architectur-swagger/infrastructure/persistence"
	"gorm.io/gorm"
)

type Migration struct {
	Repo repository.DBRepository
}

// TodoMigrate implements repository.MigrationRepository
func (m *Migration) TodoMigrate() error {
	item := model.Todo{}
	err := m.Repo.Migrate(item)
	if err != nil {
		return err
	}
	return nil
}

func NewMigration(DB *gorm.DB) repository.MigrationRepository {
	migration := Migration{
		Repo: persistence.NewDBRepository(DB)}
	return &migration
}
