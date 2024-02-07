package adapet_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConvImplPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "presen/adapterパッケージのテスト")
}
