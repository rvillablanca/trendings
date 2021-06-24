package command

import "errors"

var (
	mockErr = errors.New("just a mock error")
)

type storeMock struct {
}

func (s storeMock) AddSearchResult(location string, hashtags []string) {}

func (s storeMock) Clean() {}

type engineMock struct {
}

func (e engineMock) Search(location string) ([]string, error) {
	return []string{"hello", "world"}, nil
}

type badReaderMock struct {
}

func (b badReaderMock) Read(p []byte) (n int, err error) {
	return 0, mockErr
}
