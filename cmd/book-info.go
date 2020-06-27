package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	bookinfo "github.com/denysvitali/book-info/pkg"
	"github.com/sirupsen/logrus"
	"os"
)

var args struct {
	FileName string `arg:"positional" help:"path to the file that you want to analyze"`
}

func main(){
	arg.MustParse(&args)

	if _, err := os.Stat(args.FileName); os.IsNotExist(err) {
		logrus.Fatalf("file %s does not exist!", err)
	}

	book, err := bookinfo.Open(args.FileName)

	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("%-10s\t%-50s\n", "Title:", book.Title)
	fmt.Printf("%-10s\t%-50s\n", "Author:", book.Author)
	fmt.Printf("%-10s\t%-50s\n", "Date:", book.Date)
	fmt.Printf("%-10s\t%-50s\n", "Publisher:", book.Publisher)
	fmt.Printf("%-10s\t%-50s\n", "Language:", book.Language)
}