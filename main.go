package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

var (
	app         string = "uuid-gen"
	version     string = "dev"
	description string = "UUID Generator (Version1 / Version4)"
	site        string = "https://github.com/dassump/uuid-gen"
	info        string = "%s (%s)\n\n%s\n%s\n\n"
	usage       string = "Usage of %s:\n"

	count       int
	count_name  string = "count"
	count_value int    = 1
	count_usage string = "How many to generate"
	count_error string = "Fail to generate UUIDs, %s\n"

	output       string
	output_name  string = "type"
	output_value string = "v4"
	output_usage string = "UUID Version1 (v1) / Version4 (v4)"
	output_error string = "Type %s is not valid\n"

	quotation       string
	quotation_name  string = "quotation"
	quotation_value string = ""
	quotation_usage string = "Quotation character"

	separator       string
	separator_name  string = "separator"
	separator_value string = "\n"
	separator_usage string = "Separator character"

	format = map[string]func() (uuid.UUID, error){
		"v1": uuid.NewUUID,
		"v4": uuid.NewRandom,
	}
)

func init() {
	log.SetFlags(log.Lmsgprefix)
	log.SetPrefix("Error: ")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), info, app, version, description, site)
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}

	flag.IntVar(&count, count_name, count_value, count_usage)
	flag.StringVar(&output, output_name, output_value, output_usage)
	flag.StringVar(&quotation, quotation_name, quotation_value, quotation_usage)
	flag.StringVar(&separator, separator_name, separator_value, separator_usage)
	flag.Parse()

	if _, ok := format[output]; !ok {
		flag.Usage()
		log.Fatalf(output_error, output)
	}
}

func main() {
	for x := 1; x <= count; x++ {
		id, err := format[output]()
		if err != nil {
			log.Fatalf(count_error, err)
		}

		fmt.Print(quotation + id.String() + quotation)

		if x < count {
			fmt.Print(separator)
		}
	}
}
