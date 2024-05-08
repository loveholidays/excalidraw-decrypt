package main

import (
	"excalidraw-decrypt/pkg/excalidrawdecrypt"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: excalidraw-decrypt <documentID,decryptionKey>")
		os.Exit(1)
	}

	decrypter := excalidrawdecrypt.CreateShareableExcalidrawDecrypter()

	shareableLinkParams := os.Args[1]
	decrypted, err := decrypter.Decrypt(shareableLinkParams)

	if err != nil {
		slog.Error("Error decrypting Excalidraw diagram", slog.Attr{
			Key:   "shareableID",
			Value: slog.StringValue(shareableLinkParams),
		}, slog.Attr{
			Key:   "err",
			Value: slog.StringValue(err.Error()),
		})
		os.Exit(1)
	}

	fmt.Print(decrypted)
}
