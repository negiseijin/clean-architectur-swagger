package repository

/*
DBリポジトリ
*/
type DBRepository interface {
	Migrate(item interface{}) error                                     // 移行
	Find(model interface{}, params interface{}, item interface{}) error // 検索
	Read(model interface{}) ([]map[string]interface{}, error)           // 全検索
	Create(model interface{}, item interface{}) error                   // 新規作成
	Update(model interface{}, item interface{}) error                   // 更新
	Delete(model interface{}, item interface{}) error                   // 削除
}
