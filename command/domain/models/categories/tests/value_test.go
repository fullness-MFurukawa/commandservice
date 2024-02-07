package categories_test

import (
	"commandservice/domain/models/categories"
	"commandservice/errs"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Categoryエンティティを構成する値オブジェクト", Ordered, Label("CategoryId構造体の生成"), func() {
	var empty_str *errs.DomainError        // 空文字列　長さ36に違反する
	var length_over *errs.DomainError      // 36文字より大きい文字列　長さ36に違反する
	var not_uuid *errs.DomainError         // UUID以外の文字列を指定する
	var category_id *categories.CategoryId // UUID文字列を指定する
	var uid string
	// 前処理
	BeforeAll(func() {
		_, empty_str = categories.NewCategoryId("")
		_, length_over = categories.NewCategoryId("aaaaaaaaaabbbbbbbbbbccccccccccdddddddddd")
		_, not_uuid = categories.NewCategoryId("aaaaaaaaaabbbbbbbbbbccccccccccdddddd")
		id, _ := uuid.NewRandom()
		uid = id.String()
		category_id, _ = categories.NewCategoryId(id.String())
	})
	// 文字数の検証
	Context("文字数の検証", func() {
		It("空文字列の場合、errs.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("カテゴリIDの長さは36文字でなければなりません。")))
		})
		It("36文字より大きい文字列の場合、errs.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("カテゴリIDの長さは36文字でなければなりません。")))
		})
	})
	// UUID形式の検証
	Context("UUID形式の検証", func() {
		It("uuid以外の文字列の場合、errs.DomainErrorが返る", func() {
			Expect(not_uuid).To(Equal(errs.NewDomainError("カテゴリIDはUUIDの形式でなければなりません。")))
		})
		It("uuid文字列の場合、CategoryIdが返る", func() {
			Expect(category_id.Value()).To(Equal(uid))
		})
	})
	// 同一性の検証
	Context("同一性の検証", func() {
		It("アドレスが同じ場合はtrueが返る", func() {
			result := category_id.Equals(category_id)
			Expect(result).To(Equal(true))
		})
		It("値が同じ場合はtrueが返る", func() {
			c_id, _ := categories.NewCategoryId(uid)
			result := category_id.Equals(c_id)
			Expect(result).To(Equal(true))
		})
		It("値が異なる場合はfalseが返る", func() {
			uid, _ := uuid.NewRandom()
			c_id, _ := categories.NewCategoryId(uid.String())
			result := category_id.Equals(c_id)
			Expect(result).To(Equal(false))
		})
	})
})
var _ = Describe("Categoryエンティティを構成する値オブジェクト", Ordered, Label("CategoryName構造体の生成"), func() {
	var empty_str *errs.DomainError            // 空文字列　長さ20に違反する
	var length_over *errs.DomainError          // 20文字より大きい文字列　長さ36に違反する
	var category_name *categories.CategoryName // 20文字以内の文字列を指定する
	// 前処理
	BeforeAll(func() {
		_, empty_str = categories.NewCategoryName("")
		_, length_over = categories.NewCategoryName("aaaaaaaaaabbbbbbbbbbccccccccccd")
		category_name, _ = categories.NewCategoryName("食料品")
	})
	// 文字数の検証
	Context("文字数の検証", func() {
		It("空文字列の場合、errs.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("カテゴリ名の長さは2文字以上、20文字以内です。")))
		})
		It("20文字より大きいの場合,errs.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("カテゴリ名の長さは2文字以上、20文字以内です。")))
		})
	})
	// 文字数の検証
	Context("有効な文字数の検証", func() {
		It("2文字以上20文字以内場合,CategoryNameが返る", func() {
			Expect(category_name.Value()).To(Equal("食料品"))
		})
	})
})
