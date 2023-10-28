package products

import (
	"commandservice/domain/models/categories"
	"commandservice/errs"

	"github.com/google/uuid"
)

// 商品を表すEntity
type Product struct {
	id       *ProductId           // 商品番号
	name     *ProductName         // 商品名
	price    *ProductPrice        // 単価
	category *categories.Category // カテゴリ
}

// ゲッター
func (ins *Product) Id() *ProductId {
	return ins.id
}
func (ins *Product) Name() *ProductName {
	return ins.name
}
func (ins *Product) Price() *ProductPrice {
	return ins.price
}
func (ins *Product) Category() *categories.Category {
	return ins.category
}

// 値の変更
func (ins *Product) ChangeProductName(name *ProductName) {
	ins.name = name
}
func (ins *Product) ChangeProductPrice(price *ProductPrice) {
	ins.price = price
}
func (ins *Product) ChangeCategory(category *categories.Category) {
	ins.category = category
}

// 同一性検証メソッド
func (ins *Product) Equals(obj *Product) (bool, *errs.DomainError) {
	if obj == nil {
		return false, errs.NewDomainError("引数でnilが指定されました。")
	}
	result := ins.id.Equlas(obj.Id())
	return result, nil
}

// コンストラクタ
func NewProduct(name *ProductName, price *ProductPrice, category *categories.Category) (*Product, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil { // UUIDを生成する
		return nil, errs.NewDomainError(err.Error())
	} else {
		// 商品ID用値オブジェクトを生成する
		if id, err := NewProductId(uid.String()); err != nil {
			return nil, err
		} else {
			// 商品エンティティのインスタンスを生成して返す
			return &Product{
				id:       id,
				name:     name,
				price:    price,
				category: category,
			}, nil
		}
	}
}

// 商品エンティティの再構築
func BuildProduct(id *ProductId, name *ProductName, price *ProductPrice, category *categories.Category) *Product {
	product := Product{ // 商品エンティティのインスタンスを生成して返す
		id:       id,
		name:     name,
		price:    price,
		category: category,
	}
	return &product
}
