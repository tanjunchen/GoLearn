package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


// 顶级的 Describe 容器
// Describe块用于组织Specs，其中可以包含任意数量的：
//    BeforeEach：在Spec（It块）运行之前执行，嵌套Describe时最外层BeforeEach先执行
//    AfterEach：在Spec运行之后执行，嵌套Describe时最内层AfterEach先执行
//    JustBeforeEach：在It块，所有BeforeEach之后执行
//    Measurement

// 可以在 Describe 块内嵌套 Describe、Context、When 块
var _ = Describe("Book", func() {
	var (
		// 通过闭包在BeforeEach和It之间共享数据
		longBook  Book
		shortBook Book
	)
	// 此函数用于初始化Spec的状态，在It块之前运行。如果存在嵌套Describe，则最
	// 外面的BeforeEach最先运行
	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			// 通过It来创建一个Spec
			It("should be a novel", func() {
				// Gomega的Expect用于断言
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})

})

type Book struct {
	Title string
	Author string
	Pages int
}

func (b Book) CategoryByLength() string {
	if b.Pages < 300{
		return "SHORT STORY"
	}
	return "NOVEL"
}