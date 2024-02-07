package prepare

import (
	"crypto/tls"
	"embed"
	"io/fs"

	"github.com/fullness-MFurukawa/samplepb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// 埋め込むファイル指定
//
//go:embed commandservice.pem commandservice-key.pem
var content embed.FS

// gRPCサーバの生成とServer機能の登録
type CommandServer struct {
	Server *grpc.Server // gRPCServer
}

// 証明書、秘密鍵をロードする
func loadPem() credentials.TransportCredentials {
	cert, _ := fs.ReadFile(content, "commandservice.pem")
	key, _ := fs.ReadFile(content, "commandservice-key.pem")
	if certificate, err := tls.X509KeyPair(cert, key); err != nil {
		panic(err)
	} else {
		creds := credentials.NewServerTLSFromCert(&certificate)
		return creds
	}
}

// コンストラクタ TLSを利用
func NewCommandServer(category pb.CategoryCommandServer, product pb.ProductCommandServer) *CommandServer {
	// gRPCサーバを生成する
	server := grpc.NewServer(grpc.Creds(loadPem()))
	// CategoryCommandServerを登録する
	pb.RegisterCategoryCommandServer(server, category)
	// ProductCommandServerを登録する
	pb.RegisterProductCommandServer(server, product)
	return &CommandServer{Server: server}
}
