package virtualfilesystem

import (
	"testing"
)

func TestIsValidateFileName(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidateFileName(tt.args.filename); got != tt.want {
				t.Errorf("IsValidateFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
