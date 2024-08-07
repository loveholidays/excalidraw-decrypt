/*
excalidraw-decrypt
Copyright (C) 2023 loveholidays

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program; if not, write to the Free Software Foundation,
Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
*/

package excalidrawdecrypt_test

import (
	_ "embed"
	"github.com/loveholidays/excalidraw-decrypt/pkg/excalidrawdecrypt"
	mock_fetch "github.com/loveholidays/excalidraw-decrypt/pkg/excalidrawdecrypt/mocks/fetch"
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
