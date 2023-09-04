package processimage

import (
	"testing"
)

func TestGetImageFileName(t *testing.T) {
	type args struct {
		imageURL string
	}
	res := GetImageFileName(args.imageURL)
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				"https://images.unsplash.com/5/unsplash-kitsune-4.jpg?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=bc01c83c3da0425e9baa6c7a9204af81",
			},
			want: res,
		},
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
			got, err := DownloadCompressAndStoreImage(tt.args.imageURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadCompressAndStoreImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DownloadCompressAndStoreImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessProductimage(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcessProductimage(tt.args.a)
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
