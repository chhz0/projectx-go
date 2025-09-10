package main

// import (
// 	"bufio"
// 	"context"
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/chhz0/projectx-go/pkg/app"
// 	"github.com/spf13/cobra"
// )

// var (
// 	uppercase bool
// 	lowercase bool
// 	titlecase bool
// 	wordcount bool
// 	charcount bool
// )

// func main() {
// 	cmd := &app.Command{
// 		Use:   "textfmt [text]",
// 		Short: "A text formatter",
// 		Long: `textfmt is a CLI tool for formatting and analyzing text.

// You can provide text as arguments or pipe it in via stdin.
// Multiple formatting options can be applied simultaneously.`,
// 		Args: cobra.ArbitraryArgs,
// 		Init: func(cmd *cobra.Command) {
// 			cmd.Flags().BoolVarP(&uppercase, "upper", "u", false, "Convert text to uppercase")
// 			cmd.Flags().BoolVarP(&lowercase, "lower", "l", false, "Convert text to lowercase")
// 			cmd.Flags().BoolVarP(&titlecase, "title", "t", false, "Convert text to title case")
// 			cmd.Flags().BoolVar(&wordcount, "words", false, "Count words in text")
// 			cmd.Flags().BoolVar(&charcount, "chars", false, "Count characters in text")
// 		},
// 		PreRun: func(cmd *cobra.Command, args []string) {},
// 		Run:    runTextFormatter,
// 	}

// 	if err := cmd.Exec(context.Background()); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }

// func runTextFormatter(cmd *cobra.Command, args []string) {
// 	var text string

// 	// Get input text from args or stdin
// 	if len(args) > 0 {
// 		text = strings.Join(args, " ")
// 	} else {
// 		// Read from stdin
// 		scanner := bufio.NewScanner(os.Stdin)
// 		var lines []string
// 		for scanner.Scan() {
// 			lines = append(lines, scanner.Text())
// 		}
// 		text = strings.Join(lines, "\n")
// 	}

// 	if text == "" {
// 		fmt.Println("No input text provided")
// 		return
// 	}

// 	result := text

// 	// Apply formatting transformations
// 	if uppercase {
// 		result = strings.ToUpper(result)
// 	}
// 	if lowercase {
// 		result = strings.ToLower(result)
// 	}
// 	if titlecase {
// 		result = strings.Title(result)
// 	}

// 	// Output the formatted text
// 	fmt.Println(result)

// 	// Show analysis if requested
// 	if wordcount {
// 		words := len(strings.Fields(text))
// 		fmt.Printf("Words: %d\n", words)
// 	}
// 	if charcount {
// 		chars := len(text)
// 		fmt.Printf("Characters: %d\n", chars)
// 	}
// }
