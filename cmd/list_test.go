package cmd_test

import (
	"errors"
	"story"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "story/cmd"
	"story/cmd/cmdfakes"
)

var _ = Describe("ListCommand", func() {

	var (
		ui         *cmdfakes.FakeUI
		appFetcher *cmdfakes.FakeAppFetcher

		listCmd    *ListCommand
		executeErr error
	)

	BeforeEach(func() {
		ui = &cmdfakes.FakeUI{}
		appFetcher = &cmdfakes.FakeAppFetcher{}

		listCmd = &ListCommand{
			UI:         ui,
			AppFetcher: appFetcher,
		}
	})

	JustBeforeEach(func() {
		executeErr = listCmd.Execute(nil)
	})

	When("there are no apps", func() {
		BeforeEach(func() {
			appFetcher.FetchAppsReturns(nil, nil)
		})

		It("displays an informative message", func() {
			Expect(executeErr).NotTo(HaveOccurred())

			Expect(ui.DisplayTextCallCount()).NotTo(BeZero())
			Expect(ui.DisplayTextArgsForCall(0)).To(Equal("No apps found."))
			Expect(ui.DisplayTextArgsForCall(1)).To(Equal("Create your first app using `story apps create`"))
		})
	})

	When("there are apps", func() {
		BeforeEach(func() {
			appFetcher.FetchAppsReturns([]story.App{{Name: "foo"}, {Name: "bar"}}, nil)
		})

		It("displays a table with the app info", func() {
			Expect(executeErr).NotTo(HaveOccurred())

			Expect(ui.DisplayTableCallCount()).NotTo(BeZero())
			headers, data := ui.DisplayTableArgsForCall(0)
			Expect(headers).To(Equal([]string{"NAME"}))
			Expect(data).To(Equal([][]string{
				[]string{"foo"},
				[]string{"bar"},
			}))
		})
	})

	When("fetching apps errors", func() {
		BeforeEach(func() {
			appFetcher.FetchAppsReturns(nil, errors.New("explode"))
		})

		It("propagates the error", func() {
			Expect(executeErr).To(MatchError("explode"))
		})
	})

})
