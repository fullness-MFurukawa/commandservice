package repository

import (
	"commandservice/domain/models/products"
	"commandservice/errs"
	"commandservice/infra/sqlboiler/handler"
	"commandservice/infra/sqlboiler/models"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// ProductRepositoryインターフェイスの実装
type productRepositorySQLBoiler struct{}

// コンストラクタ
func NewproductRepositorySQLBoiler() products.ProductRepository {
	// フック関数の登録
	models.AddProductHook(boil.AfterInsertHook, ProductAfterInsertHook)
	models.AddProductHook(boil.AfterUpdateHook, ProductAfterUpdateHook)
	models.AddProductHook(boil.AfterDeleteHook, ProductAfterDeleteHook)
	return &productRepositorySQLBoiler{}
}

// 同名の商品の存在確認結果を返す
func (rep *productRepositorySQLBoiler) Exists(ctx context.Context, tran *sql.Tx, product *products.Product) error {
	// レコードの存在確認条件を作成する
	condition := models.ProductWhere.Name.EQ(product.Name().Value())
	// レコードの存在を確認するクエリを実行する
	exists, err := models.Products(condition).Exists(ctx, tran)
	if err != nil {
		return handler.DBErrHandler(err)
	}
	if !exists { // 同じ名称の商品は存在していない?
		return nil
	} else {
		return errs.NewCRUDError(fmt.Sprintf("%sは既に登録されています。", product.Name().Value()))
	}
}

// 新しい商品を永続化する
func (rep *productRepositorySQLBoiler) Create(ctx context.Context, tran *sql.Tx, product *products.Product) error {
	// SqlBoilerのモデルを生成する
	new_product := models.Product{
		ID:         0,
		ObjID:      product.Id().Value(),
		Name:       product.Name().Value(),
		Price:      int(product.Price().Value()),
		CategoryID: product.Category().Id().Value(),
	}
	// 商品を永続化する
	if err := new_product.Insert(ctx, tran, boil.Whitelist("obj_id", "name", "price", "category_id")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 商品を変更する
func (rep *productRepositorySQLBoiler) UpdateById(ctx context.Context, tran *sql.Tx, product *products.Product) error {
	// 更新対象を取得する
	up_model, err := models.Products(qm.Where("obj_id = ?", product.Id().Value())).One(ctx, tran)
	if up_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("商品番号:%sは存在しないため、更新できませんでした。", product.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 取得したモデルの値を変更する
	up_model.Name = product.Name().Value()
	up_model.Price = int(product.Price().Value())
	// 更新を実行する
	if _, err = up_model.Update(ctx, tran, boil.Whitelist("obj_id", "name", "price")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 商品を削除する
func (rep *productRepositorySQLBoiler) DeleteById(ctx context.Context, tran *sql.Tx, product *products.Product) error {
	// 削除対象を取得する
	del_model, err := models.Products(qm.Where("obj_id = ?", product.Id().Value())).One(ctx, tran)
	if del_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("商品番号:%sは存在しないため、削除できませんでした。", product.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 削除を実行する
	if _, err = del_model.Delete(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 登録処理後に実行されるフック
func ProductAfterInsertHook(ctx context.Context, exec boil.ContextExecutor, product *models.Product) error {
	log.Printf("商品ID:%s 商品名:%s 単価:%d カテゴリ番号: %s を登録しました。\n",
		product.ObjID, product.Name, product.Price, product.CategoryID)
	return nil
}

// 変更処理後に実行されるフック
func ProductAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor, product *models.Product) error {
	log.Printf("商品ID:%s 商品名:%s 単価:%d カテゴリ番号: %s を変更しました。\n",
		product.ObjID, product.Name, product.Price, product.CategoryID)
	return nil
}

// 削除処理後に実行されるフック
func ProductAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, product *models.Product) error {
	log.Printf("商品ID:%s 商品名:%s 単価:%d カテゴリ番号: %s を削除しました。\n",
		product.ObjID, product.Name, product.Price, product.CategoryID)
	return nil
}
