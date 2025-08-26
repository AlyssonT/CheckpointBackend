package repositories

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type FileRepository struct {
	BasePath string
}

func NewFileRepository(basePath string) *FileRepository {
	os.MkdirAll(basePath, os.ModePerm)
	return &FileRepository{BasePath: basePath}
}

func (fr *FileRepository) SaveAvatar(file multipart.File, userID uint) (string, error) {
	defer file.Close()
	filename := fmt.Sprintf("avatar_%d_%d.png", userID, time.Now().UnixNano())
	filepath := filepath.Join(fr.BasePath, filename)
	dst, err := os.Create(filepath)

	if err != nil {
		return "", err
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return "/uploads/avatars/" + filename, nil
}
