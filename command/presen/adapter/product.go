package adapter

import (
	"commandservice/domain/models/products"

	"github.com/fullness-MFurukawa/samplepb/pb"
)

// パラメータと実行結果の変換インターフェス
type ProductAdapter interface {
	// ProductUpParamからProductに変換する
	ToEntity(param *pb.ProductUpParam) (*products.Product, error)
	// サービス実行結果からProductUpResultに変換する
	ToResult(result any) *pb.ProductUpResult
}
