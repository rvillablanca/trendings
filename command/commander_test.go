package command

import (
	"io"
	"strings"
	"testing"
)

func TestCommander_Listen(t *testing.T) {

	tests := []struct {
		name    string
		reader  io.Reader
		wantErr bool
	}{
		{"invalid", badReaderMock{}, true},
		{"valid country", strings.NewReader("chile"), false},
		{"comando exit", strings.NewReader("exit"), false},
		{"comando clean", strings.NewReader("clean"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCommander(storeMock{}, engineMock{}, tt.reader)
			if err := c.Listen(); (err != nil) != tt.wantErr {
				t.Errorf("Listen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
