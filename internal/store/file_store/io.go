package file_store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/veliancreate/books-api/internal/entity"
)

type FileIOManager interface {
	ReadFile() ([]entity.Book, error)
	WriteFile(books []entity.Book) error
}

type FileIO struct {
	path string
}

func NewFileIO(path string) *FileIO {
	return &FileIO{
		path,
	}
}

func (f *FileIO) ReadFile() ([]entity.Book, error) {
	var books []entity.Book

	jsonFile, err := os.Open(f.path)

	if err != nil {
		return books, fmt.Errorf("error opening json file %w", err)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return books, fmt.Errorf("error reading json file %w", err)
	}

	err = json.Unmarshal(byteValue, &books)
	if err != nil {
		return books, fmt.Errorf("error unmarshalling json file %w", err)
	}

	return books, nil
}

func (f *FileIO) WriteFile(books []entity.Book) error {
	file, err := json.MarshalIndent(books, "", " ")
	if err != nil {
		return fmt.Errorf("error marshal indent json %w", err)
	}

	err = ioutil.WriteFile(f.path, file, 0644)
	if err != nil {
		return fmt.Errorf("error writing file %w", err)
	}

	return nil
}
