package processimage

import (
	"testing"
)

func TestGetImageFileName(t *testing.T) {
	type args struct {
		imageURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetImageFileName(tt.args.imageURL); got != tt.want {
				t.Errorf("GetImageFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownloadCompressAndStoreImage1(t *testing.T) {
	type args struct {
		imageURL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DownloadCompressAndStoreImage1(tt.args.imageURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadCompressAndStoreImage1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DownloadCompressAndStoreImage1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownloadCompressAndStoreImage(t *testing.T) {

}

func TestProcessProductimage(t *testing.T) {
	tests := []struct {
		name    string
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcessProductimage(2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessProductimage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ProcessProductimage() = %v, want %v", got, tt.want)
			}
		})
	}
}
