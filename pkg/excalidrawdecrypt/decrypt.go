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

package excalidrawdecrypt

import (
	"bytes"
	"compress/zlib"
	"excalidraw-decrypt/pkg/excalidrawdecrypt/buffers"
	"excalidraw-decrypt/pkg/excalidrawdecrypt/crypto"
	"excalidraw-decrypt/pkg/excalidrawdecrypt/fetch"
	"io"
	"strings"
)

type EcalidrawDecrypter interface {
	Decrypt(shareableID string) (string, error)
}

type ShareableExcalidrawDecrypter struct {
	Fetcher fetch.DocumentIDFetcher
}

func CreateShareableExcalidrawDecrypter() ShareableExcalidrawDecrypter {
	return ShareableExcalidrawDecrypter{
		Fetcher: fetch.CreateNewExcalidrawURLFetcher("https://json.excalidraw.com/api/v2"),
	}
}

func (decrypter *ShareableExcalidrawDecrypter) Decrypt(shareableID string) (string, error) {

	splittedShareableID := strings.Split(shareableID, ",")
	documentID := splittedShareableID[0]
	privateKey := splittedShareableID[1]

	encryptedDiagram, err := decrypter.Fetcher.FetchEncryptedDiagram(documentID)

	if err != nil {
		return "", err
	}

	splitBuffers, err := buffers.SplitBuffers(encryptedDiagram)
	if err != nil {
		return "", err
	}

	// currently unused
	// encodingMetadataBuffer := splitBuffers[0]
	iv := splitBuffers[1]
	buffer := splitBuffers[2]

	decrypted, err := crypto.Decrypt(buffer, iv, privateKey)

	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(decrypted)
	gzreader, err := zlib.NewReader(reader)
	if err != nil {
		return "", err
	}

	unzipped, err := io.ReadAll(gzreader)
	if err != nil {
		return "", err
	}

	// There seem to be 16 bytes of garbage at the beginning
	stripped := unzipped[16:]

	return string(stripped), nil
}
