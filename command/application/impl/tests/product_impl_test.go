package impl_test

import (
	"commandservice/application"
	"commandservice/application/service"
	"commandservice/domain/models/categories"
	"commandservice/domain/models/products"
	"commandservice/errs"
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
)

var _ = Describe("productServiceImpl構造体", Ordered, Label("メソッドのテスト"), func() {
	var product *products.Product
	var category *categories.Category
	var service service.ProductService
	var ctx context.Context
	var container *fx.App
	// 前処理
	BeforeAll(func() {
		category_id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
		category_name, _ := categories.NewCategoryName("文房具")
		category = categories.BuildCategory(category_id, category_name)
		product_name, _ := products.NewProductName("ボールペン")
		product_price, _ := products.NewProductPrice(uint32(150))
		product, _ = products.NewProduct(product_name, product_price, category)
		ctx = context.Background()
		// サービスのインスタンス生成
		container = fx.New(
			application.SrvDepend,
			fx.Populate(&service),
		)
		// fxを起動し、起動時にエラーがないことを確認する
		err := container.Start(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
	// 後処理
	AfterAll(func() {
		// fxを停止する
		err := container.Stop(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
	// テスト
	Context("Add()メソッドのテスト", Label("Add"), func() {
		It("商品登録が成功し、nilが返る", func() {
			result := service.Add(ctx, product)
			Expect(result).To(BeNil())
		})
		It("存在する商品名の場合、errs.CRUDErrorが返る", func() {
			result := service.Add(ctx, product)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("%sは既に登録されています。", product.Name().Value()))))
		})
	})
	Context("Update()メソッドのテスト", Label("Update"), func() {
		It("存在するobj_idを指定すると、nilが返る", func() {
			product_name, _ := products.NewProductName("ボールペン(黒)")
			product_price, _ := products.NewProductPrice(uint32(200))
			up_product := products.BuildProduct(product.Id(), product_name, product_price, category)
			result := service.Update(ctx, up_product)
			Expect(result).To(BeNil())
		})
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			product_name, _ := products.NewProductName("ボールペン(黒)")
			product_price, _ := products.NewProductPrice(uint32(200))
			up_product, _ := products.NewProduct(product_name, product_price, category)
			result := service.Update(ctx, up_product)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("商品番号:%sは存在しないため、更新できませんでした。", up_product.Id().Value()))))
		})
	})
	Context("Delete()メソッドのテスト", Label("Delete"), func() {
		It("存在するobj_idを指定すると、nilが返る", func() {
			result := service.Delete(ctx, product)
			Expect(result).To(BeNil())
		})
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			product_name, _ := products.NewProductName("ボールペン(黒)")
			product_price, _ := products.NewProductPrice(uint32(200))
			del_product, _ := products.NewProduct(product_name, product_price, category)
			result := service.Delete(ctx, del_product)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("商品番号:%sは存在しないため、削除できませんでした。", del_product.Id().Value()))))
		})
	})
})
