package categories_test

import (
	"commandservice/domain/models/categories"
	"commandservice/errs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Categoryエンティティ", Ordered, Label("Category構造体の生成"), func() {
	Context("インスタンス生成", Label("Create Category"), func() {
		It("新しいカテゴリを生成する", Label("NewCategory"), func() {
			category_name, _ := categories.NewCategoryName("食料品")
			category, _ := categories.NewCategory(category_name)
			Expect(category.Id()).ToNot(BeNil())
			Expect(category.Name().Value()).To(Equal("食料品"))
		})
		It("カテゴリを再構築する", Label("BuildCategory"), func() {
			category_id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
			category_name, _ := categories.NewCategoryName("食料品")
			category := categories.BuildCategory(category_id, category_name)
			Expect(category.Id().Value()).To(Equal("b1524011-b6af-417e-8bf2-f449dd58b5c0"))
			Expect(category.Name().Value()).To(Equal("食料品"))
		})
	})
})

var _ = Describe("Categoryエンティティ", Ordered, Label("Categoryの同一性検証"), func() {
	It("比較対象がnil", Label("nil検証"), func() {
		By("インスタンスの生成")
		category := categories.BuildCategory(nil, nil)
		By("Equals()にnilを指定")
		_, err := category.Equals(nil)
		By("errs.DomainErrorを評価する")
		Expect(err).To(Equal(errs.NewDomainError("引数でnilが指定されました。")))
	})
	It("異なるカテゴリID", Label("false検証"), func() {
		By("2つインスタンスの生成")
		category_name, _ := categories.NewCategoryName("食料品")
		category_a, _ := categories.NewCategory(category_name)
		category_name, _ = categories.NewCategoryName("食料品")
		category_b, _ := categories.NewCategory(category_name)
		By("Equals()にcategory_bを指定")
		result, _ := category_a.Equals(category_b)
		By("falseを評価する")
		Expect(result).To(Equal(false))
	})
	It("同一のカテゴリID", Label("trueの検証"), func() {
		By("2つインスタンスの生成")
		category_name, _ := categories.NewCategoryName("食料品")
		category_a, _ := categories.NewCategory(category_name)

		category_b := categories.BuildCategory(category_a.Id(), category_a.Name())
		By("Equals()にcategory_bを指定")
		result, _ := category_a.Equals(category_b)
		By("trueを評価する")
		Expect(result).To(Equal(true))
	})
})
