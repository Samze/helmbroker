package main_test

import (
	"fmt"
	"net/http"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Helmbroker", func() {
	var session *gexec.Session
	var cmd *exec.Cmd
	port := 8080

	BeforeEach(func() {
		broker, err := gexec.Build("github.com/samze/helmbroker/cmd")
		Expect(err).NotTo(HaveOccurred())
		cmd = exec.Command(broker)
	})

	JustBeforeEach(func() {
		var err error
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Eventually(session.Terminate()).Should(gexec.Exit())
	})

	It("starts a web server", func() {
		Eventually(func() error {
			_, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
			return err
		}).Should(BeNil())
	})

	Context("when configuring username and password", func() {
		BeforeEach(func() {
		})

		It("responds with 401 with invalid auth credentials", func() {
		})

		It("responds with 404 with valid auth credentials", func() {
		})
	})
})
