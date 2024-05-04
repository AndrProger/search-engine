package revindex

import (
	"reflect"
	"testing"
)

func TestAddIndexes(t *testing.T) {
	type args struct {
		strArr []string
		url    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test right test",
			args: args{strArr: []string{"Hello", "world!", "This", "is", "a", "test", ".", "Another", "paragraph."}, url: "https://habr.com/ru/articles/812199/"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddIndexes(tt.args.strArr, tt.args.url)
		})
	}
}

func TestGetUrls(t *testing.T) {
	strArr := []string{"Hello", "world!", "This", "is", "a", "test", ".", "Another", "paragraph."}
	url := "url"

	AddIndexes(strArr, url)
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test right  1",
			args: args{str: "Hello"},
			want: []string{"url"},
		},
		{
			name: "Test right  2",
			args: args{str: "world!"},
			want: []string{"url"},
		},
		{
			name: "Test lose ",
			args: args{str: "Error!"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUrls(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUrls() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
