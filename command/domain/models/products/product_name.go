package products

import (
	"commandservice/errs"
	"fmt"
	"unicode/utf8"
)

// 商品名を保持する値オブジェクト
type ProductName struct {
	value string
}

// ゲッター
func (ins *ProductName) Value() string {
	return ins.value
}

// コンストラクタ
func NewProductName(value string) (*ProductName, *errs.DomainError) {
	const MIN_LENGTH int = 5  // フィールドの最小文字数
	const MAX_LENGTH int = 30 // フィールドの最大文字数
	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("商品名の長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}
	return &ProductName{value: value}, nil
}
