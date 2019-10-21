package acceptance_test

import (
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Apps", func() {

	var (
		tmpDir     string
		testRunDir string

		session *Session
		cmd     *exec.Cmd
	)

	BeforeEach(func() {
		tmpDir = createTempDir()
		testRunDir = getCWD()

		Expect(os.Chdir(tmpDir)).To(Succeed())
	})

	AfterEach(func() {
		Expect(os.Chdir(testRunDir)).To(Succeed())
		Expect(os.RemoveAll(tmpDir)).To(Succeed())
	})

	BeforeEach(func() {
		cmd = exec.Command(storyPath)
		cmd.Args = append(cmd.Args, "apps")
		cmd.Env = append(os.Environ(), getEnvOrError("STORYSCRIPT_TOKEN"))
		cmd.Stdout = GinkgoWriter
		cmd.Stderr = GinkgoWriter
	})

	Describe("Listing", func() {
		BeforeEach(func() {
			cmd.Args = append(cmd.Args, "list")
		})

		JustBeforeEach(func() {
			session = execStory(cmd)
		})

		It("exits successfully", func() {
			Eventually(session).Should(Exit(0))
		})

		When("there are no apps", func() {
			It("displays an informative message", func() {
				Eventually(session).Should(Say("No apps found."))
				Eventually(session).Should(Say("Create your first app using `story apps create`"))
			})
		})

		When("an app has been created", func() {
			var appName string

			BeforeEach(func() {
				appName = createRandomApp()
			})

			AfterEach(func() {
				destroyApp(appName)
			})

			It("displays that app in the list", func() {
				Eventually(session).Should(Say(`NAME`))
				Eventually(session).Should(Say(appName))
			})
		})
	})
})

func story(args ...string) *Session {
	cmd := exec.Command("story", args...)
	session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return session
}

func createRandomApp() string {
	appName := generateRandomAppName()
	Eventually(story("apps", "create", appName)).Should(Exit(0))
	return appName
}

func destroyApp(appName string) {
	Eventually(story("apps", "destroy", "-y", "-a", appName)).Should(Exit(0))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func generateRandomAppName() string {
	b := make([]rune, 20)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func createTempDir() string {
	tmpDir, e := ioutil.TempDir("", "")
	Expect(e).NotTo(HaveOccurred())
	return tmpDir
}

func getCWD() string {
	cwd, e := os.Getwd()
	Expect(e).NotTo(HaveOccurred())
	return cwd
}
