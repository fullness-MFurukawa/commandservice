package products_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// 商品エンティティ、値オブジェクトテスト用エントリーポイント
func TestEntityPackage(t *testing.T) {
	RegisterFailHandler(Fail)                      // テストが失敗した場合のハンドラを登録する
	RunSpecs(t, "domain/models/productsパッケージのテスト") // 登録されたすべてのテストを実行する
}
