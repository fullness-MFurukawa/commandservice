package products

import (
	"context"
	"database/sql"
)

// 商品をアクセスするリポジトリインターフェイス
type ProductRepository interface {
	// 同名の商品があるか
	Exists(ctx context.Context, tran *sql.Tx, product *Product) error
	// 新しい商品を永続化する
	Create(ctx context.Context, tran *sql.Tx, product *Product) error
	// 商品を変更する
	UpdateById(ctx context.Context, tran *sql.Tx, product *Product) error
	// 商品を削除する
	DeleteById(ctx context.Context, tran *sql.Tx, product *Product) error
}
