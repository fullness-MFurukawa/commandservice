package products

import (
	"commandservice/errs"
	"fmt"
	"regexp"
	"unicode/utf8"
)

// 商品番号を保持する値オブジェクト(UUIDを保持する)
type ProductId struct {
	value string
}

// valueフィールドのゲッター
func (ins *ProductId) Value() string {
	return ins.value
}

// 同一性検証
func (ins *ProductId) Equlas(value *ProductId) bool {
	if ins == value { // アドレスが同じ?
		return true
	}
	// 値の比較
	return ins.value == value.Value()
}

// コンストラクタ
func NewProductId(value string) (*ProductId, *errs.DomainError) {
	// フィールドの長さ
	const LENGTH int = 36
	// UUIDの正規表現
	const REGEXP string = "([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})"
	// 引数の文字数チェック
	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("商品IDの長さは%d文字でなければなりません。", LENGTH))
	}
	// 引数の正規表現(UUID)チェック
	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, errs.NewDomainError("商品IDはUUIDの形式でなければなりません。")
	}
	return &ProductId{value: value}, nil
}
