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

package main

import (
	"fmt"
	"github.com/loveholidays/excalidraw-decrypt/pkg/excalidrawdecrypt"
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
