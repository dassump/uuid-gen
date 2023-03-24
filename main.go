package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

var (
	app         = "uuid-gen"
	version     = "dev"
	description = "UUID Generator (Version1 / Version4)"
	site        = "https://github.com/dassump/uuid-gen"

	count     = flag.Int("count", 1, "How many UUIDs to generate")
	output    = flag.String("type", "v4", "UUID Version1 (v1) / Version4 (v4)")
	quotation = flag.String("quotation", "", "Quotation character(s)")
	separator = flag.String("separator", "\n", "Separator character(s)")

	format = map[string]func() (uuid.UUID, error){
		"v1": uuid.NewUUID,
		"v4": uuid.NewRandom,
	}

	replacer = strings.NewReplacer(
		`\b`, "\b", // backspace
		`\f`, "\f", // form feed
		`\n`, "\n", // line feed or newline
		`\r`, "\r", // carriage return
		`\t`, "\t", // horizontal tab
		`\v`, "\v", // vertical tab
	)

	logger = log.New(os.Stderr, "(ERROR) ", log.Lmsgprefix+log.LstdFlags)
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s (%s)\n\n%s\n%s\n\n", app, version, description, site)
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if _, ok := format[*output]; !ok {
		flag.Usage()
		logger.Fatalf("Type %s is not valid", *output)
	}

	*quotation = replacer.Replace(*quotation)
	*separator = replacer.Replace(*separator)
}

func main() {
	for i := 1; i <= *count; i++ {
		id, err := format[*output]()
		if err != nil {
			logger.Fatalf("Failed to generate UUIDs: %s", err)
		}

		fmt.Print(*quotation, id.String(), *quotation)

		if i < *count {
			fmt.Print(*separator)
		}
	}
}
