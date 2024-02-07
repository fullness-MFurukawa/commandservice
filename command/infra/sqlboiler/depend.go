package sqlboiler

import (
	"commandservice/infra/sqlboiler/repository"

	"go.uber.org/fx"
)

// SQLBoilerを利用したRepositoryの依存定義
var RepDepend = fx.Options(
	fx.Provide(
		// Repositoryインターフェイス実装のコンストラクタを指定
		repository.NewcategoryRepositorySQLBoiler, // カテゴリ用Reposititory
		repository.NewproductRepositorySQLBoiler,  // 商品用Repository
	),
)
