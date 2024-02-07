package impl_test

import (
	"commandservice/application"
	"commandservice/application/service"
	"commandservice/domain/models/categories"
	"commandservice/errs"
	"context"
	"fmt"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
)

var _ = Describe("categoryServiceImpl構造体", Ordered, Label("メソッドのテスト"), func() {
	var category *categories.Category
	var service service.CategoryService
	var ctx context.Context
	var container *fx.App
	// 前処理
	BeforeAll(func() {
		// テストデータの作成
		name, _ := categories.NewCategoryName("飲料水")
		category, _ = categories.NewCategory(name)
		// Contextの生成
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

	// Add()メソッドのテスト
	Context("Add()メソッドのテスト", Label("Add"), func() {
		// テスト
		It("カテゴリ登録が成功し、nilが返る", func() {
			result := service.Add(ctx, category)
			Expect(result).To(BeNil())
		})
		It("存在するカテゴリ名の場合、errs.CRUDErrorが返る", func() {
			result := service.Add(ctx, category)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("%sは既に登録されています。", category.Name().Value()))))
		})
	})
	// Update()メソッドのテスト
	Context("Update()メソッドのテスト", Label("Update"), func() {
		// テスト
		It("存在するobj_idを指定すると、nilが返る", func() {
			result := service.Update(ctx, category)
			log.Println("存在するobj_idを指定すると、nilが返る", result)
			Expect(result).To(BeNil())
		})
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			name, _ := categories.NewCategoryName("飲料水")
			up_category, _ := categories.NewCategory(name)
			result := service.Update(ctx, up_category)
			log.Println("存在しないobj_idを指定すると、errs.CRUDErrorが返る", result)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("カテゴリ番号:%sは存在しないため、更新できませんでした。", up_category.Id().Value()))))
		})
	})
	// Delete()メソッドのテスト
	Context("Delete()メソッドのテスト", Label("Delete"), func() {
		// テスト
		It("存在するobj_idを指定すると、nilが返る", func() {
			result := service.Delete(ctx, category)
			Expect(result).To(BeNil())
		})
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			name, _ := categories.NewCategoryName("飲料水")
			del_category, _ := categories.NewCategory(name)
			result := service.Delete(ctx, del_category)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("カテゴリ番号:%sは存在しないため、削除できませんでした。", del_category.Id().Value()))))
		})
	})
})
