package revindex

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"testing"
)

func Test_contains(t *testing.T) {
	type args struct {
		arr []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test right  1",
			args: args{str: "world!", arr: []string{"Hello", "world!", "This", "is", "a", "test", ".", "Another", "paragraph."}},
			want: true,
		},
		{
			name: "Test right  1",
			args: args{str: "error!", arr: []string{"Hello", "world!", "This", "is", "a", "test", ".", "Another", "paragraph."}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.arr, tt.args.str); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeIndexesToJson(t *testing.T) {
	// Тест для успешного случая
	indexes := make(map[string][]string)
	indexes["Hello"] = []string{"url"}
	t.Run("Test success case", func(t *testing.T) {
		wantW := `{"Hello":["url"]}`
		w := &bytes.Buffer{}
		err := WriteIndexesToJson(indexes, w)

		if err != nil {
			t.Errorf("writeIndexesToJson() error = %v, wantErr false", err)
		}

		if gotW := w.String(); gotW != wantW {
			t.Errorf("writeIndexesToJson() gotW = %v, want %v", gotW, wantW)
		}
	})

	t.Run("Test error writer case ", func(t *testing.T) {
		wantErr := true
		w := &errorWriter{}
		err := WriteIndexesToJson(indexes, w)

		if (err != nil) != wantErr {
			t.Errorf("writeIndexesToJson() error = %v, wantErr %v", err, wantErr)
		}
	})
}

type errorWriter struct{}

func (ew *errorWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("mock error")
}

func TestAddIndexes1(t *testing.T) {
	indexes := make(map[string][]string)
	indexes["Hello"] = []string{"url"}
	indexes["world"] = []string{"url"}
	type args struct {
		indexes map[string][]string
		strArr  []string
		url     string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "Test right",
			args: args{
				indexes: make(map[string][]string),
				strArr:  []string{"Hello", "world"},
				url:     "url",
			},
			want: indexes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddIndexes(tt.args.indexes, tt.args.strArr, tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddIndexes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readIndexesFromJson(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]string
		wantErr bool
	}{
		{
			name: "Test right",
			args: args{
				r: bytes.NewReader(
					[]byte(`{"Hello":["url"]}`),
				),
			},
			want: map[string][]string{
				"Hello": {"url"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadIndexesFromJson(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("readIndexesFromJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readIndexesFromJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
