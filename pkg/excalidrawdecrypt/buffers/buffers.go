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

package buffers

import (
	"encoding/binary"
	"errors"
)

const VersionOrChunkSizeField = 4
const BuffersVersion = 1

func SplitBuffers(concatenatedBuffer []byte) ([][]byte, error) {
	cursor := 0
	var outputBuffers [][]byte

	// First 4 bytes is the version
	version, err := SeekFrom(concatenatedBuffer, cursor)

	if err != nil {
		return nil, err
	}

	if version > BuffersVersion {
		return nil, errors.New("cannot split this version of buffers")
	}

	cursor += VersionOrChunkSizeField

	for {
		chunkSize, err := SeekFrom(concatenatedBuffer, cursor)
		if err != nil {
			return nil, err
		}

		cursor += VersionOrChunkSizeField

		endPos := cursor + int(chunkSize)
		if endPos > len(concatenatedBuffer) {
			endPos = len(concatenatedBuffer)
		}
		outputBuffers = append(outputBuffers, concatenatedBuffer[cursor:endPos])
		cursor = endPos
		if cursor >= len(concatenatedBuffer) {
			break
		}
	}

	return outputBuffers, nil
}

func SeekFrom(concatenatedBuffer []byte, pos int) (uint32, error) {

	if pos+VersionOrChunkSizeField > cap(concatenatedBuffer) {
		return 0, errors.New("trying to scan further than buffer")
	}

	return binary.BigEndian.Uint32(concatenatedBuffer[pos : pos+VersionOrChunkSizeField]), nil
}
