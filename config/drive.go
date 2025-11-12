package config

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func UploadToDrive(file io.Reader, fileName string) (string, error) {
	ctx := context.Background()

	srv, err := drive.NewService(ctx, option.WithCredentialsFile("config/credentials.json"))
	if err != nil {
		return "", fmt.Errorf("gagal membuat service Drive: %v", err)
	}

	f := &drive.File{Name: fileName, Parents: []string{"root"}} // bisa ganti folder ID
	fileDrive, err := srv.Files.Create(f).Media(file).Do()
	if err != nil {
		return "", fmt.Errorf("gagal upload ke Drive: %v", err)
	}

	return fileDrive.Id, nil
}
