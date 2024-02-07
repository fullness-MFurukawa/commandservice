package categories

import (
	"commandservice/errs"

	"github.com/google/uuid"
)

// 商品カテゴリを表すEntity
type Category struct {
	id   *CategoryId   // カテゴリ番号
	name *CategoryName // カテゴリ名
}

// ゲッター
func (ins *Category) Id() *CategoryId {
	return ins.id
}
func (ins *Category) Name() *CategoryName {
	return ins.name
}

// 値の変更
func (ins *Category) ChangeCategoryName(name *CategoryName) {
	ins.name = name
}

// 同一性検証
func (ins *Category) Equals(obj *Category) (bool, *errs.DomainError) {
	if obj == nil {
		return false, errs.NewDomainError("引数でnilが指定されました。")
	}
	result := ins.id.Equals(obj.Id())
	return result, nil
}

// コンストラクタ
func NewCategory(name *CategoryName) (*Category, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil { // UUIDを生成する
		return nil, errs.NewDomainError(err.Error())
	} else {
		if id, err := NewCategoryId(uid.String()); err != nil {
			return nil, errs.NewDomainError(err.Error())
		} else {
			return &Category{
				id:   id,
				name: name,
			}, nil
		}
	}
}

// Categoryエンティティを再構築する
func BuildCategory(id *CategoryId, name *CategoryName) *Category {
	return &Category{
		id:   id,
		name: name,
	}
}
