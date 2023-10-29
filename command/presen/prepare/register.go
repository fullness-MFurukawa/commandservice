package prepare

import (
	"github.com/fullness-MFurukawa/samplepb/pb"
	"google.golang.org/grpc"
)

// gRPCサーバの生成とServer機能の登録
type CommandServer struct {
	Server *grpc.Server // gRPCServer
}

// コンストラクタ 平文を利用する
func NewCommandServer(category pb.CategoryCommandServer, product pb.ProductCommandServer) *CommandServer {
	// gRPCサーバを生成する
	server := grpc.NewServer()
	// CategoryCommandServerを登録する
	pb.RegisterCategoryCommandServer(server, category)
	// ProductCommandServerを登録する
	pb.RegisterProductCommandServer(server, product)
	return &CommandServer{Server: server}
}
