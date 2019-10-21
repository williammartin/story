package console_test

import (
	. "story/console"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("Console Writer", func() {

	var testWriter *Writer

	BeforeEach(func() {
		testWriter = &Writer{
			Out: NewBuffer(),
		}
	})

	Describe("DisplayText", func() {
		It("prints text with templated values to the out buffer", func() {
			testWriter.DisplayText("This is a test for the {{.Struct}} struct", map[string]interface{}{"Struct": "Console"})
			Expect(testWriter.Out).To(Say(`This is a test for the Console struct\n`))
		})
	})

	Describe("DisplayTable", func() {
		It("prints headers", func() {
			testWriter.DisplayTable([]string{"foo", "bar", "baz"}, nil)
			Expect(testWriter.Out).To(Say(`foo\tbar\tbaz`))
		})

		It("prints data", func() {
			data := [][]string{
				[]string{"qux", "yad", "zil"},
			}
			testWriter.DisplayTable([]string{"foo", "bar", "baz"}, data)
			Expect(testWriter.Out).To(Say(`qux\tyad\tzil`))
		})
	})
})
