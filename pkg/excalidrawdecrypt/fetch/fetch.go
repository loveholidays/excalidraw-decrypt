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

package fetch

//go:generate mockgen -destination=../mocks/fetch/mock_fetch.go -source=./fetch.go

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type DocumentIDFetcher interface {
	FetchEncryptedDiagram(documentID string) ([]byte, error)
}

type ExcalidrawURLFetcher struct {
	v2Endpoint string
}

func CreateNewExcalidrawURLFetcher(v2Endpoint string) ExcalidrawURLFetcher {
	return ExcalidrawURLFetcher{v2Endpoint: v2Endpoint}
}

func (e ExcalidrawURLFetcher) FetchEncryptedDiagram(documentID string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s", e.v2Endpoint, documentID)
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
