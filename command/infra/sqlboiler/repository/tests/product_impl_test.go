package repository_test

import (
	"commandservice/domain/models/categories"
	"commandservice/domain/models/products"
	"commandservice/errs"
	"commandservice/infra/sqlboiler/repository"
	"context"
	"database/sql"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ = Describe("productRepositorySQLBoiler構造体", Ordered, Label("ProductRepositoryインターフェイスメソッドのテスト"), func() {
	var category *categories.Category
	var rep products.ProductRepository
	var ctx context.Context
	var tran *sql.Tx
	// 前処理
	BeforeAll(func() {
		// リポジトリの生成
		rep = repository.NewproductRepositorySQLBoiler() // Repositoryの生成
		// カテゴリエンティティの生成
		category_id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
		category_name, _ := categories.NewCategoryName("文房具")
		category = categories.BuildCategory(category_id, category_name)
	})
	// テスト毎の前処理
	BeforeEach(func() {
		ctx = context.Background()       // Contextの生成
		tran, _ = boil.BeginTx(ctx, nil) // トランザクションの開始
	})
	// テスト毎の後処理
	AfterEach(func() {
		tran.Rollback() // トランザクションのロールバック
	})
	// Exists()メソッドのテスト
	Context("同名の商品の存在確認結果を返す", Label("Exists"), func() {
		It("存在しない商品の場合nilが返る", func() {
			productName, _ := products.NewProductName("水性ボールペン")
			product := products.BuildProduct(nil, productName, nil, nil)
			result := rep.Exists(ctx, tran, product)
			Expect(result).To(BeNil())
		})
		It("存在する所品の場合、errs.CRUDErrorが返る", func() {
			productName, _ := products.NewProductName("水性ボールペン(青)")
			product := products.BuildProduct(nil, productName, nil, nil)
			result := rep.Exists(ctx, tran, product)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("%sは既に登録されています。", productName.Value()))))
		})
	})
	// Create()メソッドのテスト
	Context("新しい商品を永続化する", Label("Create"), func() {
		It("商品が登録成功し、nilが返る", func() {
			product_name, _ := products.NewProductName("ボールペン")
			product_price, _ := products.NewProductPrice(uint32(150))
			product, _ := products.NewProduct(product_name, product_price, category)
			result := rep.Create(ctx, tran, product)
			Expect(result).To(BeNil())
		})
		//	productテーブルのobj_idをユニーク設定後テスト!!
		It("obj_idが同じカテゴリを追加すると、errs.CRUDErrorが返る", func() {
			product_id, _ := products.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
			product_name, _ := products.NewProductName("ボールペン")
			product_price, _ := products.NewProductPrice(uint32(200))
			product := products.BuildProduct(product_id, product_name, product_price, category)
			result := rep.Create(ctx, tran, product)
			err := errs.NewCRUDError("一意制約違反です。")
			Expect(result).To(Equal(err))
		})
	})
	// UpdateById()メソッドのテスト
	Context("商品を変更する", Label("UpdateById"), func() {
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			product_name, _ := products.NewProductName("ボールペン")
			product_price, _ := products.NewProductPrice(uint32(200))
			product, _ := products.NewProduct(product_name, product_price, category)
			result := rep.UpdateById(ctx, tran, product)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("商品番号:%sは存在しないため、更新できませんでした。", product.Id().Value()))))
		})
		It("存在するobj_idを指定すると、nilが返る", func() {
			product_id, _ := products.NewProductId("8f81a72a-58ef-422b-b472-d982e8665292")
			product_name, _ := products.NewProductName("ボールペン")
			product_price, _ := products.NewProductPrice(uint32(200))
			product := products.BuildProduct(product_id, product_name, product_price, category)
			result := rep.UpdateById(ctx, tran, product)
			Expect(result).To(BeNil())
		})
	})
	// DeleteById()メソッドのテスト
	Context("商品を削除する", Label("DeleteById"), func() {
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			product_name, _ := products.NewProductName("ボールペン")
			product_price, _ := products.NewProductPrice(uint32(200))
			product, _ := products.NewProduct(product_name, product_price, nil)
			result := rep.DeleteById(ctx, tran, product)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("商品番号:%sは存在しないため、削除できませんでした。", product.Id().Value()))))
		})
		It("存在するobj_idを指定すると、nilが返る", func() {
			// 削除対象のカテゴリを登録する
			product_id, _ := products.NewProductId("8f81a72a-58ef-422b-b472-d982e8665292")
			product_name, _ := products.NewProductName("ボールペン")
			product_price, _ := products.NewProductPrice(uint32(200))
			product := products.BuildProduct(product_id, product_name, product_price, category)
			rep.Create(ctx, tran, product)
			// 登録したカテゴリを削除する
			result := rep.DeleteById(ctx, tran, product)
			Expect(result).To(BeNil())
		})
	})
})
