package readFile

import (
	"io"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func Test_openFile(t *testing.T) {
	r := strings.NewReader(`test,test,test`)
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    io.Reader
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				fileName: "test.csv",
			},
			want:    r,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpenFile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				b, _ :=ioutil.ReadAll(got)
				w, _ := ioutil.ReadAll(tt.want)
				t.Errorf("exp = %v want = %v", b, w)
				//t.Errorf("openFile() = %v, wantErr %v", false, tt.wantErr)
			}
		})
	}
}
