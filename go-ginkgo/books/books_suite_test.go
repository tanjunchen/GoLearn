package books_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	// 将 Ginkgo 的 Fail 函数传递给 Gomega，Fail 函数用于标记测试失败
	// 这是 Ginkgo 和 Gomega 唯一的交互点
	// 如果 Gomega 断言失败，就会调用 Fail 进行处理
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}
