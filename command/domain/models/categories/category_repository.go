package categories

import (
	"context"
	"database/sql"
)

// カテゴリをアクセスするリポジトリインターフェイス
type CategoryRepository interface {
	// 同名の商品カテゴリが存在確認結果を返す
	Exists(ctx context.Context, tran *sql.Tx, category *Category) error
	// 新しい商品カテゴリを永続化する
	Create(ctx context.Context, tran *sql.Tx, category *Category) error
	// 商品カテゴリを変更する
	UpdateById(ctx context.Context, tran *sql.Tx, category *Category) error
	// 商品カテゴリを削除する
	DeleteById(ctx context.Context, tran *sql.Tx, category *Category) error
}
