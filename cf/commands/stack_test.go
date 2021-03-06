package commands_test

import (
	"errors"

	testapi "github.com/cloudfoundry/cli/cf/api/stacks/fakes"
	"github.com/cloudfoundry/cli/cf/models"
	testcmd "github.com/cloudfoundry/cli/testhelpers/commands"
	testconfig "github.com/cloudfoundry/cli/testhelpers/configuration"
	testreq "github.com/cloudfoundry/cli/testhelpers/requirements"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/cli/cf/commands"
	. "github.com/cloudfoundry/cli/testhelpers/matchers"
)

var _ = Describe("stack command", func() {
	var (
		ui                  *testterm.FakeUI
		cmd                 ListStack
		repo                *testapi.FakeStackRepository
		requirementsFactory *testreq.FakeReqFactory
	)

	BeforeEach(func() {
		ui = &testterm.FakeUI{}
		config := testconfig.NewRepositoryWithDefaults()
		requirementsFactory = &testreq.FakeReqFactory{LoginSuccess: true}
		repo = &testapi.FakeStackRepository{}
		cmd = NewListStack(ui, config, repo)
	})

	Describe("login requirements", func() {
		It("fails if the user is not logged in", func() {
			requirementsFactory.LoginSuccess = false

			Expect(testcmd.RunCommand(cmd, []string{}, requirementsFactory)).To(BeFalse())
		})
	})

	It("returns the stack guid when '--guid' flag is provided", func() {
		stack1 := models.Stack{
			Name:        "Stack-1",
			Description: "Stack 1 Description",
			Guid:        "Stack-1-GUID",
		}

		repo.FindByNameReturns(stack1, nil)

		testcmd.RunCommand(cmd, []string{"Stack-1", "--guid"}, requirementsFactory)

		Expect(len(ui.Outputs)).To(Equal(1))
		Expect(ui.Outputs[0]).To(Equal("Stack-1-GUID"))
	})

	It("returns the empty string as guid when '--guid' flag is provided and stack doesn't exist", func() {
		stack1 := models.Stack{
			Name:        "Stack-1",
			Description: "Stack 1 Description",
			Guid:        "Stack-1-GUID",
		}

		repo.FindByNameReturns(stack1, nil)

		testcmd.RunCommand(cmd, []string{"Stack-1", "--guid"}, requirementsFactory)

		Expect(len(ui.Outputs)).To(Equal(1))
		Expect(ui.Outputs[0]).To(Equal("Stack-1-GUID"))
	})

	It("lists the stack requested", func() {
		repo.FindByNameReturns(models.Stack{}, errors.New("Stack Stack-1 not found"))

		testcmd.RunCommand(cmd, []string{"Stack-1", "--guid"}, requirementsFactory)

		Expect(len(ui.Outputs)).To(Equal(1))
		Expect(ui.Outputs[0]).To(Equal(""))
	})

	It("informs user if stack is not found", func() {
		repo.FindByNameReturns(models.Stack{}, errors.New("Stack Stack-1 not found"))

		testcmd.RunCommand(cmd, []string{"Stack-1"}, requirementsFactory)

		Expect(ui.Outputs).To(BeInDisplayOrder(
			[]string{"FAILED"},
			[]string{"Stack Stack-1 not found"},
		))
	})
})
