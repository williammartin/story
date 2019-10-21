package acceptance_test

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var storyPath string

var _ = BeforeSuite(func() {
	var err error
	storyPath, err = gexec.Build("../story/main.go")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func TestIntegration(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(time.Second * 5)
	RunSpecs(t, "Story Acceptance Suite")
}

func execStory(cmd *exec.Cmd) *gexec.Session {
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	return session
}

func getEnvOrError(env string) string {
	value := os.Getenv(env)
	if value == "" {
		Fail(fmt.Sprintf("Environment variable '%s' must be set", env))
	}

	return value
}
