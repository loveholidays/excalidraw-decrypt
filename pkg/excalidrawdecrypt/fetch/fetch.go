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
