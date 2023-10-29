package service

import (
	"commandservice/domain/models/categories"
	"context"
)

// Category更新サービスインターフェイス
type CategoryService interface {
	// カテゴリを登録する
	Add(ctx context.Context, category *categories.Category) error
	// カテゴリを変更する
	Update(ctx context.Context, category *categories.Category) error
	// カテゴリを削除する
	Delete(ctx context.Context, categoryId *categories.Category) error
}
