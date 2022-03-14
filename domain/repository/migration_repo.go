package repository

/*
migrationリポジトリ
*/
type MigrationRepository interface {
	TodoMigrate() error // TODOモデル
}
