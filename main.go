package main // import "4d63.com/dnsovertlsproxy"

import (
	"encoding/csv"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
)

var version = "<not set>"

func main() {
	flagPrintHelp := flag.Bool("help", false, "print this help")
	flagPrintVersion := flag.Bool("version", false, "print version")
	flagFile := flag.String("f", "", "csv file")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "csv extracts columns from a CSV file using Go templates.\n")
		fmt.Fprintf(os.Stderr, "Usage: csv -f=file.csv '{{index . 3}}'\n")
		fmt.Fprintf(os.Stderr, "       cat file.csv | csv '{{index . 3}}'\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *flagPrintHelp {
		flag.Usage()
		return
	}

	if *flagPrintVersion {
		fmt.Println("csv", "v"+version)
		return
	}

	var reader io.Reader
	if *flagFile == "" {
		reader = os.Stdin
	} else {
		file, err := os.Open(*flagFile)
		if err != nil {
			fmt.Println("Error opening csv file:", err)
			os.Exit(1)
		}
		reader = file
	}

	query := flag.Arg(0)

	err := transpose(os.Stdout, reader, query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func transpose(w io.Writer, r io.Reader, query string) error {
	tmpl, err := template.New("tmpl").Parse(query)
	if err != nil {
		return fmt.Errorf("error parsing query: %s", err)
	}

	csvReader := csv.NewReader(r)
	csvReader.ReuseRecord = true
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err == csv.ErrFieldCount {
			// this is okay
		} else if err != nil {
			return fmt.Errorf("Error parsing csv: %s", err)
		}

		err = tmpl.Execute(w, record)
		if err != nil {
			return fmt.Errorf("Error transposing: %s", err)
		}
	}

	return nil
}
