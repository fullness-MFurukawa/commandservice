package server

import (
	"commandservice/application/service"
	"commandservice/presen/adapter"
	"context"

	"github.com/fullness-MFurukawa/samplepb/pb"
)

// カテゴリ更新サーバの実装
type categoryServer struct {
	adapter adapter.CategoryAdapter // カテゴリ変換
	service service.CategoryService // カテゴリ更新サービス
	// 生成されたUnimplementedCategoryCommandServerをエンベデッド
	pb.UnimplementedCategoryCommandServer
}

// コンストラクタ
func NewcategoryServer(adapter adapter.CategoryAdapter, service service.CategoryService) pb.CategoryCommandServer {
	return &categoryServer{adapter: adapter, service: service}
}

// カテゴリの追加 pb.CategoryCommandServerインターフェイスのメソッド実装
func (ins *categoryServer) Create(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	// pb.CategoryUpParamをentity.Categoryに変換する
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil // CategoryUpResultにエラーを設定
	} else {
		// サービスのAdd()メソッドを実行する
		if err := ins.service.Add(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil // CategoryUpResultにエラーを設定
		}
		return ins.adapter.ToResult(category), nil // CategoryUpResultにCategoryを設定
	}
}

// カテゴリの変更 pb.CategoryCommandServerインターフェイスのメソッド実装
func (ins *categoryServer) Update(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	// pb.CategoryUpParamをentity.Categoryに変換する
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil // CategoryUpResultにエラーを設定
	} else {
		// サービスのUpdate()メソッドを実行する
		if err := ins.service.Update(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil // CategoryUpResultにエラーを設定
		}
		return ins.adapter.ToResult(category), nil // CategoryUpResultにCategoryを設定
	}
}

// カテゴリの削除 pb.CategoryCommandServerインターフェイスのメソッド実装
func (ins *categoryServer) Delete(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	// pb.CategoryUpParamをentity.Categoryに変換する
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil // CategoryUpResultにエラーを設定
	} else {
		// サービスのDelete()メソッドを実行する
		if err := ins.service.Delete(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil // CategoryUpResultにエラーを設定
		}
		// CategoryUpResultにCategoryを設定して返す
		return ins.adapter.ToResult(category), nil
	}
}
