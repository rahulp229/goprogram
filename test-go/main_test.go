package main

import (
	"reflect"
	"testing"
)

func Test_removeElem(t *testing.T) {
	type args struct {
		country []string
	}
	actual := []string{"india", "india"}
	expected := []string{"india"}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{name: "remove", args: args{country: actual}, want: expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElem(tt.args.country); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("actual() = %v, want %v", got, tt.want)
			}
		})
	}
}
