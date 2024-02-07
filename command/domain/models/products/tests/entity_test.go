package products_test

import (
	"commandservice/domain/models/categories"
	"commandservice/domain/models/products"
	"commandservice/errs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Productエンティティ", Ordered, Label("Product構造体の生成"), func() {
	// 前処理
	BeforeAll(func() {
	})
	var _ = Describe("Productエンティティ", Ordered, Label("Product構造体の生成"), func() {
		Context("インスタンス生成", Label("Create Product"), func() {
			It("新しい商品のインスタンス生成", Label("NewProduct"), func() {
				name, _ := products.NewProductName("チョコレート")
				price, _ := products.NewProductPrice(150)
				product, _ := products.NewProduct(name, price, nil)
				Expect(product.Id().Value()).ToNot(BeNil())
				Expect(product.Name().Value()).To(Equal("チョコレート"))
				Expect(product.Price().Value()).To(Equal(uint32(150)))
				Expect(product.Category()).To(BeNil())
			})
			It("商品のインスタンスを構築する", Label("BuildProduct"), func() {
				id, _ := products.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
				name, _ := products.NewProductName("チョコレート")
				price, _ := products.NewProductPrice(150)
				product := products.BuildProduct(id, name, price, nil)
				Expect(product.Id().Value()).To(Equal("ac413f22-0cf1-490a-9635-7e9ca810e544"))
				Expect(product.Name().Value()).To(Equal("チョコレート"))
				Expect(product.Price().Value()).To(Equal(uint32(150)))
				Expect(product.Category()).To(BeNil())
			})
		})
	})
})

var _ = Describe("Productエンティティ", Ordered, Label("Productの同一性検証"), func() {
	var category *categories.Category
	var product *products.Product

	// 前処理
	BeforeAll(func() {
		category_name, _ := categories.NewCategoryName("食料品")
		category, _ = categories.NewCategory(category_name)
		product_name, _ := products.NewProductName("ポテトチップス")
		product_price, _ := products.NewProductPrice(uint32(200))
		product, _ = products.NewProduct(product_name, product_price, category)
	})

	// エラーの検証
	Context("エラーの検証", func() {
		It("比較対象がnil", Label("nil検証"), func() {
			By("nilを指定し,DomainErrorを返すことを検証する")
			_, err := product.Equals(nil)
			Expect(err).To(Equal(errs.NewDomainError("引数でnilが指定されました。")))
		})
	})

	// 比較結果の検証
	Context("比較結果の検証", func() {
		It("異なる商品ID", Label("false検証"), func() {
			product_name, _ := products.NewProductName("ポテトチップス")
			product_price, _ := products.NewProductPrice(uint32(200))
			p, _ := products.NewProduct(product_name, product_price, category)
			By("異なるProductを指定し,falseを返すことを検証する")
			result, _ := product.Equals(p)
			Expect(result).To(Equal(false))
		})
		It("同一の商品ID", Label("trueの検証"), func() {
			p := products.BuildProduct(
				product.Id(),
				product.Name(),
				product.Price(), category)
			By("同一のProductを指定し,trueを返すことを検証する")
			result, _ := product.Equals(p)
			Expect(result).To(Equal(true))
		})
	})
})
