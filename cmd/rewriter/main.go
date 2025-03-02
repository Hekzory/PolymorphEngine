package main

import (
	"flag"
	"fmt"
	"github.com/Hekzory/polymorphengine/internal/rewriter"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	
	// Define command-line flags
	inputFile := flag.String("input", "", "Path to the Go file to rewrite")
	outputFile := flag.String("output", "", "Path to save the rewritten file (defaults to <input>.rewritten.go)")
	
	// Parse flags
	flag.Parse()
	
	// Create a new rewriter
	r := rewriter.NewRewriter()
	
	// Handle non-flag arguments as input files
	if flag.NArg() > 0 && *inputFile == "" {
		*inputFile = flag.Arg(0)
	}
	
	// Validate input
	if *inputFile == "" {
		fmt.Println("Error: No input file specified")
		fmt.Println("Usage: rewriter [options] -input <file.go>")
		flag.PrintDefaults()
		os.Exit(1)
	}
	
	// Set default output file if not specified
	if *outputFile == "" {
		*outputFile = *inputFile + ".rewritten.go"
	}
	
	// Perform the rewriting
	fmt.Printf("Rewriting %s to %s...\n", *inputFile, *outputFile)
	
	// Rewrite the file
	rewritten, err := r.RewriteFile(*inputFile)
	if err != nil {
		fmt.Printf("Error rewriting file: %v\n", err)
		os.Exit(1)
	}
	
	// Save the rewritten content
	err = r.SaveRewrittenFile(*outputFile, rewritten)
	if err != nil {
		fmt.Printf("Error saving rewritten file: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("Rewriting completed successfully!")
} 