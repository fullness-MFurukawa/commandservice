package server_test

import (
	"commandservice/infra/sqlboiler/handler"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHelperPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "presen/serverパッケージのテスト")
}

// 全テストスイートが実行される前に1度だけ実行される関数
var _ = BeforeSuite(func() {
	absPath, _ := filepath.Abs("../../../infra/sqlboiler/config/database.toml")
	// database.tomlファイルのパスを環境変数に設定する
	os.Setenv("DATABSE_TOML_PATH", absPath)
	err := handler.DBConnect() // データベースに接続する
	Expect(err).NotTo(HaveOccurred(), "データベース接続が失敗したのでテストを中止します。")
})
