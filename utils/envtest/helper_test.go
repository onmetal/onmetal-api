package envtest_test

import (
	"os"

	. "github.com/onmetal/onmetal-api/utils/envtest"
	. "github.com/onmetal/onmetal-api/utils/testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helper", func() {
	ctx := SetupContext()

	Describe("GetPath", func() {
		It("should retrieve remote requested paths into the requested directory", func() {
			dst := GinkgoT().TempDir()
			Expect(GetPath(ctx, dst, "github.com/onmetal/onmetal-api//config/apiserver/apiservice/bases")).To(Equal(dst))
		})

		It("should return the path to local dependencies", func() {
			tempDir := GinkgoT().TempDir()

			dst, err := os.MkdirTemp(tempDir, "dst")
			Expect(err).NotTo(HaveOccurred())

			path, err := os.MkdirTemp(tempDir, "path")
			Expect(err).NotTo(HaveOccurred())

			Expect(GetPath(ctx, dst, path)).To(Equal(path))
		})
	})
})
