package products_test

import (
	"commandservice/domain/models/products"
	"commandservice/errs"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Productエンティティを構成する値オブジェクト", Ordered, Label("ProductId構造体の生成"), func() {
	var empty_str *errs.DomainError    // 空文字列　長さ36に違反する
	var length_over *errs.DomainError  // 36文字より大きい文字列　長さ36に違反する
	var not_uuid *errs.DomainError     // UUID以外の文字列を指定する
	var product_id *products.ProductId // UUID文字列を指定する
	var uid string

	// 前処理
	BeforeAll(func() {
		_, empty_str = products.NewProductId("")
		_, length_over = products.NewProductId("aaaaaaaaaabbbbbbbbbbccccccccccdddddddddd")
		_, not_uuid = products.NewProductId("aaaaaaaaaabbbbbbbbbbccccccccccdddddd")
		id, _ := uuid.NewRandom()
		uid = id.String()
		product_id, _ = products.NewProductId(id.String())
	})

	// 文字数の検証
	Context("文字数の検証", Label("文字数"), func() {
		It("空文字列の場合、errs.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("商品IDの長さは36文字でなければなりません。")))
		})
		It("36文字より大きい文字列の場合、errs.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("商品IDの長さは36文字でなければなりません。")))
		})
	})

	// UUID形式の検証
	Context("UUID形式の検証", Label("UUID形式"), func() {
		It("uuid以外の文字列の場合、errs.DomainErrorが返る", func() {
			Expect(not_uuid).To(Equal(errs.NewDomainError("商品IDはUUIDの形式でなければなりません。")))
		})
		It("36文字のuuid文字列の場合、ProductIdが返る", func() {
			id, _ := products.NewProductId(uid)
			Expect(product_id).To(Equal(id))
		})
	})
})

var _ = Describe("Productエンティティを構成する値オブジェクト", Ordered, Label("ProductName構造体の生成"), func() {
	var empty_str *errs.DomainError        // 空文字列　長さ5文字以上、30文字以内に違反する
	var length_over *errs.DomainError      // 30文字より大きい文字列　長さ5文字以上、30文字以内に違反する
	var product_name *products.ProductName // 30文字以内の文字列を指定する
	// 前処理
	BeforeAll(func() {
		_, empty_str = products.NewProductName("")
		_, length_over = products.NewProductName("aaaaaaaaaabbbbbbbbbbccccccccccd")
		product_name, _ = products.NewProductName("水性ボールペン")
	})
	// 文字数の検証
	Context("文字数の検証", Label("無効な文字数"), func() {
		It("空文字列の場合、errs.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("商品名の長さは5文字以上、30文字以内です。")))
		})
		It("30文字より大きいの場合,errs.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("商品名の長さは5文字以上、30文字以内です。")))
		})
	})
	// 文字数の検証
	Context("有効な文字数の検証", Label("有効な文字数"), func() {
		It("5文字以上30文字以内場合,ProductNameが返る", func() {
			Expect(product_name.Value()).To(Equal("水性ボールペン"))
		})
	})
})

var _ = Describe("Productエンティティを構成する値オブジェクト", Ordered, Label("ProductPrice構造体の生成"), func() {
	var min_err *errs.DomainError            // 50未満の単価
	var max_err *errs.DomainError            // 10000より大きい単価
	var product_price *products.ProductPrice // 商品単価
	// 前処理
	BeforeAll(func() {
		_, min_err = products.NewProductPrice(49)
		_, max_err = products.NewProductPrice(10001)
		product_price, _ = products.NewProductPrice(1500)
	})
	// 範囲外の単価の検証
	Context("範囲外の単価の検証", Label("無効な範囲"), func() {
		It("50未満の場合、errs.DomainErrorが返る", func() {
			Expect(min_err).To(Equal(errs.NewDomainError("単価は50以上、10000以下です。")))
		})
		It("10000より大きいの単価の場合、errs.DomainErrorが返る", func() {
			Expect(max_err).To(Equal(errs.NewDomainError("単価は50以上、10000以下です。")))
		})
	})
	// 範囲内の単価の検証
	Context("範囲外の単価の検証", Label("有効な範囲"), func() {
		It("50以上10000以下の単価の場合、ProductPriceが返る", func() {
			Expect(product_price.Value()).To(Equal(uint32(1500)))
		})
	})
})
