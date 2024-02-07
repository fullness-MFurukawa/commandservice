package categories_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// カテゴリエンティティ、値オブジェクトテスト用エントリーポイント
func TestEntityPackage(t *testing.T) {
	RegisterFailHandler(Fail)                 // テストが失敗した場合のハンドラを登録する
	RunSpecs(t, "domain/entityパッケージのテストスイート") // 登録されたすべてのテストを実行する
}
