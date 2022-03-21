package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/veliancreate/books-api/internal/entity"
)

func getBooks() ([]entity.Book, error) {
	var books []entity.Book

	jsonFile, err := os.Open("seed/books.json")

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

func writeFile(books []entity.Book) error {
	file, err := json.MarshalIndent(books, "", " ")
	if err != nil {
		return fmt.Errorf("error marshal indent json %w", err)
	}

	err = ioutil.WriteFile("seed/books.json", file, 0644)
	if err != nil {
		return fmt.Errorf("error writing file %w", err)
	}

	return nil
}

func pagination() {

}