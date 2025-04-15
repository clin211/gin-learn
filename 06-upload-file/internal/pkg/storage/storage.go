package storage

import "mime/multipart"

type Storage interface {
	Save(fileHeader *multipart.FileHeader, dstPath string) error
	GetURL(objectKey string) (string, error)
	Delete(objectKey string) error
}
