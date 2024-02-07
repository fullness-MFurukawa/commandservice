package impl

import (
	"commandservice/application/service"
	"commandservice/domain/models/products"
	"commandservice/infra/sqlboiler/handler"
	"context"
)

// ProductServiceインターフェイスの実装
type productServiceImpl struct {
	rep         products.ProductRepository
	transaction // transaction構造体のエンベデッド
}

// コンストラクタ
func NewproductServiceImpl(rep products.ProductRepository) service.ProductService {
	return &productServiceImpl{rep: rep}
}

// カテゴリを登録する
func (ins *productServiceImpl) Add(ctx context.Context, product *products.Product) error {
	// トランザクションを開始する
	tran, err := ins.begin(ctx)
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 実行結果に応じてトランザクションのコミットロールバック制御する
	defer func() {
		err = ins.complete(tran, err)
	}()
	// 既に登録済であるか、確認する
	if err = ins.rep.Exists(ctx, tran, product); err != nil {
		return err
	}
	// 商品を登録する
	if err = ins.rep.Create(ctx, tran, product); err != nil {
		return err
	}
	return err
}

// カテゴリを変更する
func (ins *productServiceImpl) Update(ctx context.Context, product *products.Product) error {
	// トランザクションを開始する
	tran, err := ins.begin(ctx)
	if err != nil {
		handler.DBErrHandler(err)
	}
	// 実行結果に応じてトランザクションのコミットロールバック制御する
	defer func() {
		err = ins.complete(tran, err)
	}()
	// 商品を変更する
	if err = ins.rep.UpdateById(ctx, tran, product); err != nil {
		return err
	}
	return err
}

// 商品を削除する
func (ins *productServiceImpl) Delete(ctx context.Context, product *products.Product) error {
	// トランザクションを開始する
	tran, err := ins.begin(ctx)
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 実行結果に応じてトランザクションのコミットロールバック制御する
	defer func() {
		err = ins.complete(tran, err)
	}()
	// 商品を削除する
	if err = ins.rep.DeleteById(ctx, tran, product); err != nil {
		return err
	}
	return err
}
