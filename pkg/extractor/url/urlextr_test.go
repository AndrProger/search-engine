package url

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestExtractUrl(t *testing.T) {
	// Создаем тестовый сервер, который возвращает заранее определенный HTML-ответ
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<html><body><p>Hello world! This is a <b>test</b>.</p><p>Another paragraph.</p></body></html>"))
	}))
	defer ts.Close()

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "Test right test",
			args:    args{url: ts.URL},
			want:    []string{"Hello", "world!", "This", "is", "a", "test", ".", "Another", "paragraph."},
			wantErr: false,
		},
		{
			name:    "Test right test",
			args:    args{url: ts.URL + "error"},
			want:    []string{"Hello", "world!", "This", "is", "a", "test", ".", "Another", "paragraph."},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractUrl(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, v := range got {
				if !reflect.DeepEqual(v, tt.want[i]) {
					t.Errorf("ExtractUrl()[%v] = [%v], want [%v]", i, v, tt.want[i])
				}
			}
		})
	}
}

func Test_extractText(t *testing.T) {
	type args struct {
		body io.Reader
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test right test",
			args: args{strings.NewReader(
				"<html><body><p>Hello world! This is a <b>test</b>.</p><p>Another paragraph.</p></body></html>",
			)},
			want: []string{"Hello", "world!", "This", "is", "a", "test", ".", "Another", "paragraph."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractText(tt.args.body)
			for i, v := range got {
				if !reflect.DeepEqual(v, tt.want[i]) {
					t.Errorf("extractText()[%v] = [%v], want [%v]", i, v, tt.want[i])
				}
			}
		})
	}
}
