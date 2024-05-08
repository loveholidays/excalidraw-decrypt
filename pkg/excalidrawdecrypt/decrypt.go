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
