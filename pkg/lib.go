package bookinfo

import (
	"errors"
	"github.com/kapmahc/epub"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

type Book struct {
	Title       string
	Language    string
	Author      string
	Description string
	Publisher   string
	Date        string
}

func Open(fileName string) (*Book, error) {
	var mBook *Book = nil
	// Detect file type
	if strings.ToLower(filepath.Ext(fileName)) == ".epub" {
		// EPUB
		book, err := epub.Open(fileName)
		if err != nil {
			return nil, err
		}
		logrus.Debugf("book: %v", book)

		mBook = &Book{
			Title:       book.Opf.Metadata.Title[0],
			Language:    book.Opf.Metadata.Language[0],
			Author:      book.Opf.Metadata.Creator[0].Data,
			Description: book.Opf.Metadata.Description[0],
			Publisher:   book.Opf.Metadata.Publisher[0],
			Date:        book.Opf.Metadata.Date[0].Data,
		}
		return mBook, nil
	}

	return nil, errors.New("unrecognized book format")
}
