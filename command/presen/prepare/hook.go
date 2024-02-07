package prepare

import (
	"commandservice/infra/sqlboiler/handler"
	"context"
	"fmt"
	"log"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc/reflection"
)

func CommandServiceLifecycle(lifecycle fx.Lifecycle, server *CommandServer) {
	lifecycle.Append(fx.Hook{
		// fx起動時の処理
		OnStart: func(ctx context.Context) error {
			// データベース接続とコネクションプールを生成する
			if err := handler.DBConnect(); err != nil {
				panic(err)
			}
			port := 8082 // 8082ポートを利用するListenerを生成する
			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
			if err != nil {
				return err
			}
			reflection.Register(server.Server) // サーバリフレクションの設定
			// 作成したgRPCサーバを起動する
			go func() {
				log.Printf("Command Server 開始 ポート番号: %v", port)
				server.Server.Serve(listener)
			}()
			return nil
		},
		// fx終了時の処理
		OnStop: func(ctx context.Context) error {
			server.Server.GracefulStop() // gRPCサーバを停止する
			log.Println("Command Server 停止")
			return nil
		},
	})
}
