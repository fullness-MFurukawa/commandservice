package adapet_test

import (
	"commandservice/domain/models/categories"
	"commandservice/domain/models/products"
	"commandservice/errs"
	"commandservice/presen/adapter"

	"github.com/fullness-MFurukawa/samplepb/pb"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("productAdapaterImpl構造体", Ordered, Label("メソッドのテスト"), func() {
	var category *categories.Category
	var adapt adapter.ProductAdapter

	// 前処理
	BeforeAll(func() {
		adapt = adapter.NewproductAdapaterImpl()

		id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
		category = categories.BuildCategory(id, nil)
	})

	// テスト
	Context("ToEntity()メソッド", func() {
		It("Id、Name、Price、CategoryIdフィールドに値を設定する", func() {
			param := pb.ProductUpParam{
				Crud:       pb.CRUD_UPDATE,
				Id:         "ac413f22-0cf1-490a-9635-7e9ca810e544",
				Name:       "水性ボールペン(黒)",
				Price:      120,
				CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
			}
			result, _ := adapt.ToEntity(&param)
			product_id, _ := products.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
			product_name, _ := products.NewProductName("水性ボールペン(黒)")
			product_price, _ := products.NewProductPrice(uint32(120))
			p := products.BuildProduct(product_id, product_name, product_price, category)
			Expect(result).To(Equal(p))
		})
		It("Name、Price、CategoryIdフィールドに値を設定する", func() {
			param := pb.ProductUpParam{
				Crud:       pb.CRUD_INSERT,
				Id:         "",
				Name:       "水性ボールペン(黒)",
				Price:      120,
				CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
			}
			result, _ := adapt.ToEntity(&param)
			Expect((result.Id())).NotTo(BeNil())
			Expect(result.Name().Value()).To(Equal("水性ボールペン(黒)"))
			Expect(result.Price().Value()).To(Equal(uint32(120)))
			Expect(result.Category().Id().Value()).To(Equal("b1524011-b6af-417e-8bf2-f449dd58b5c0"))
		})
	})

	Context("ToResult()メソッド", func() {

		It("entity.Productを渡すと、pb.Productを保持したProductUpResultを返す", func() {
			product_id, _ := products.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
			product_name, _ := products.NewProductName("水性ボールペン(黒)")
			product_price, _ := products.NewProductPrice(uint32(120))
			product := products.BuildProduct(product_id, product_name, product_price, category)
			result := adapt.ToResult(product)

			c := pb.Category{Id: "b1524011-b6af-417e-8bf2-f449dd58b5c0", Name: ""}
			p := pb.Product{Id: "ac413f22-0cf1-490a-9635-7e9ca810e544",
				Name: "水性ボールペン(黒)", Price: 120, Category: &c}
			Expect(result.Product).To(Equal(&p))
			Expect(result.Error).To(BeNil())
		})

		It("errs.DomainErrorを渡すと、pb.Errorを保持したProductUpResultを返す", func() {
			err := errs.NewDomainError("水性ボールペン(黒)は既に登録されています。")
			result := adapt.ToResult(err)
			e := pb.Error{Type: "Domain Error", Message: "水性ボールペン(黒)は既に登録されています。"}
			Expect(result.Error).To(Equal(&e))
		})

		It("errs.CRUDErrorを渡すと、pb.Errorを保持したproductUpResultを返す", func() {
			err := errs.NewCRUDError("水性ボールペン(黒)は既に登録されています。")
			result := adapt.ToResult(err)
			e := pb.Error{Type: "CRUD Error", Message: "水性ボールペン(黒)は既に登録されています。"}
			Expect(result.Error).To(Equal(&e))
		})

		It("errs.InternalErrorを渡すと、pb.Errorを保持したProductUpResultを返す", func() {
			err := errs.NewInternalError("水性ボールペン(黒)は既に登録されています。")
			result := adapt.ToResult(err)
			e := pb.Error{Type: "Internal Error", Message: "只今、サービスを提供できません。"}
			Expect(result.Error).To(Equal(&e))
		})
	})
})
