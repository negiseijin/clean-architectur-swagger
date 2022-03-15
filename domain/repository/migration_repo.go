package repository

/*
migrationリポジトリ
*/
type MigrationRepository interface {
	TodoMigrate() error // TODOモデル
	Migrate() error     // migrate実行
}
