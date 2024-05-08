package excalidrawdecrypt_test

import (
	_ "embed"
	"excalidraw-decrypt/pkg/excalidrawdecrypt"
	mock_fetch "excalidraw-decrypt/pkg/excalidrawdecrypt/mocks/fetch"
	"go.uber.org/mock/gomock"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//go:embed response.bytes
var mockedResponseInBytes []byte

//go:embed decryptedDocument.txt
var decryptedDocument string

func TestDecrypt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Decrypt Suite")
}

var (
	mockCtrl          *gomock.Controller
	documentIDFetcher *mock_fetch.MockDocumentIDFetcher
)

var _ = Describe("Decrypt Suite", func() {
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		documentIDFetcher = mock_fetch.NewMockDocumentIDFetcher(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Should Decrypt excalidraw document using shareable link parameters", func() {

		It("decrypts without errors and matches the content", func() {

			documentIDFetcher.EXPECT().FetchEncryptedDiagram("pJK6JcJMr7LGOuy1NbCKP").Return(mockedResponseInBytes, nil)

			decrypter := excalidrawdecrypt.ShareableExcalidrawDecrypter{Fetcher: documentIDFetcher}

			decrypt, err := decrypter.Decrypt("pJK6JcJMr7LGOuy1NbCKP,YneEARvxllEU6vlDQfz81A")

			Expect(err).To(Not(HaveOccurred()))
			Expect(decrypt).To(Equal(decryptedDocument))
		})
	})
})
