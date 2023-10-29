package presen

import (
	"commandservice/application"
	"commandservice/presen/adapter"
	"commandservice/presen/prepare"
	"commandservice/presen/server"

	"go.uber.org/fx"
)

var CommandDepend = fx.Options(
	application.SrvDepend, // アプリケーション層の依存定義
	fx.Provide( // プレゼンテーション層の依存定義
		adapter.NewcategoryAdapaterImpl, // カテゴリ変換
		adapter.NewproductAdapaterImpl,  // 商品変換
		server.NewcategoryServer,        // カテゴリサーバ
		server.NewprductServer,          // 商品サーバ
		prepare.NewCommandServer,        // gRPCサーバ
	),
	// メソッドの起動
	fx.Invoke(prepare.CommandServiceLifecycle),
)
