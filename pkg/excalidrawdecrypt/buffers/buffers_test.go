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

package buffers_test

import (
	"github.com/loveholidays/excalidraw-decrypt/pkg/excalidrawdecrypt/buffers"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBuffers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Buffers Suite")
}

var _ = Describe("Buffers Suite", func() {

	Describe("Should Split buffers", func() {
		It("errors for version > 1", func() {

			extracted, err := buffers.SplitBuffers([]byte{0, 0, 0, 2})

			Expect(err.Error()).To(Equal("cannot split this version of buffers"))
			Expect(extracted).To(BeNil())
		})
		It("errors for position 4 if there are not enough bytes", func() {

			extracted, err := buffers.SplitBuffers([]byte{0, 0, 0, 1, 0})

			Expect(err.Error()).To(Equal("trying to scan further than buffer"))
			Expect(extracted).To(BeNil())
		})
		It("Read a valid buffer of size 1 after version", func() {

			extracted, err := buffers.SplitBuffers([]byte{0, 0, 0, 1, 0, 0, 0, 1, 5})

			Expect(err).To(Not(HaveOccurred()))
			Expect(extracted).To(Equal([][]byte{{5}}))
		})
	})

	Describe("Should extract numbers from the position", func() {
		It("returns 1 for position 0", func() {

			extracted, err := buffers.SeekFrom([]byte{0, 0, 0, 1}, 0)

			Expect(err).To(Not(HaveOccurred()))
			Expect(extracted).To(Equal(uint32(1)))
		})
		It("errors for position 1", func() {

			extracted, err := buffers.SeekFrom([]byte{0, 0, 0, 1}, 1)

			Expect(err.Error()).To(Equal("trying to scan further than buffer"))
			Expect(extracted).To(Equal(uint32(0)))
		})
	})

})
