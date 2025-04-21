package services

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func SaveImage(image *multipart.FileHeader) (string, error) {
	uploadPath := "./images"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.MkdirAll(uploadPath, os.ModePerm)
	}

	ext := filepath.Ext(image.Filename)

	// Crear nombre único
	newFileName := uuid.New().String() + ext

	// Ruta final
	fullPath := filepath.Join(uploadPath, newFileName)

	src, err := image.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Crear archivo destino
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copiar contenido
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return newFileName, nil
}

func DeleteImage(imageName string) error {
	uploadPath := "./images"
	fullPath := filepath.Join(uploadPath, imageName)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return nil // El archivo no existe, no se puede eliminar
	}

	err := os.Remove(fullPath)
	if err != nil {
		return err
	}

	return nil
}