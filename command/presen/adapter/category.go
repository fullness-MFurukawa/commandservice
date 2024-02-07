package adapter

import (
	"commandservice/domain/models/categories"

	"github.com/fullness-MFurukawa/samplepb/pb"
)

// パラメータと実行結果の変換インターフェス
type CategoryAdapter interface {
	// CategoryUpParamからCategoryに変換する
	ToEntity(param *pb.CategoryUpParam) (*categories.Category, error)
	// 実行結果からCategoryUpResultに変換する
	ToResult(result any) *pb.CategoryUpResult
}
