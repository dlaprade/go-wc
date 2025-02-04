package main 

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/tdewolff/argp"
	// Should I use
	// "https://github.com/spf13/pflag"
	// instead?
)

var (
	lengthFlag    bool = false
	bytesFlag     bool = false
	linesFlag     bool = false
	multibyteFlag bool = false
	wordFlag      bool = false
	filenames     []string

	args = argp.New("wc")
)

type Data struct {
  line       int
  bytes      int
  lines      int
  characters int
  words      int
  name       string
}

func init() {
	// libxoFlag     = flag.Bool("libxo", false, "Generate output via libxo(3) in a selection of different human and machine readable formats. See xo_parse-args(3) for details on command line arguments.")

	args.AddOpt(&lengthFlag   , "L"   , "max-line-length", "Write the length of the line containing the most bytes (default) or characters (when -m is provided) to standard output. When more than one file argument is specified, the longest input line of all files is reported as the value of the final \"total\".")
	args.AddOpt(&bytesFlag    , "c"   , "bytes"          , "The number of bytes in each input file is written to the standard output. This will cancel out any prior usage of the -m option.")
	args.AddOpt(&linesFlag    , "l"   , "lines"          , "The number of lines in each input file is written to the standard output.")
	args.AddOpt(&multibyteFlag, "m"   , "chars"          , "The number of characters in each input file is written to the standard output. If the current locale does not support multibyte characters, this is equivalent to the -c option. This will cancel out any prior usage of the -c option.")
	args.AddOpt(&wordFlag     , "w"   , "words"          , "The number of words in each input file is written to the standard output.")
	args.AddRest(&filenames   , "file", "The file name or complete path.")
}

// processStdin handles data from STD_IN
func processStdin() {
  data := Data{}
	
  // Read from STD_IN
  Process(os.Stdin, &data)

	showResults(&data)
}

// processFileOrInput handles data from a file
func processFile() {
  totals := Data{}

	for _, filename := range filenames {
		// Read the input file
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file:", err)
		}
		defer file.Close()

    data := Data{}
    data.name = filename

    Process(file, &data)

    showResults(&data)

    if len( filenames ) > 1 {
      totals.bytes      += data.bytes
      totals.words      += data.words

      if data.line > totals.line {
        totals.line = data.line
      }

      totals.characters += data.characters
      totals.lines      += data.lines
    }
	}

  if len( filenames ) > 1 {
    totals.name = "total"
    showResults(&totals)
  }
}

// Process uses processes data from the input Reader 
func Process(source io.Reader, d *Data) {
	// Continue to read from file until EOF
	reader := bufio.NewReader(source)

	for {
		  input, err := reader.ReadSlice('\n')
		  if err == io.EOF {
		  	break
		  } else if err != nil {
		  	fmt.Println("Error reading from stdin:", err)
        break
		  }

		  processLine( input, d )
  }
}

// processLine handles updating the data appropriate to the flags set
func processLine(str []byte, d *Data) {
	if lengthFlag {
    if utf8.RuneCountInString(string(str[:])) > d.line {
			d.line = utf8.RuneCountInString(string(str[:]))-1
		}
	}

	if linesFlag {
		d.lines += 1
	}

	if wordFlag {
		d.words += len(strings.Fields(string(str[:])))
	}

	if bytesFlag {
		d.bytes += len(str)
	}

	if multibyteFlag {
		d.characters += utf8.RuneCountInString(string(str[:])) // len(str)
	}
}

// showResults formats and displays the results of each count specified
func showResults( d *Data ) {
  if linesFlag {
		fmt.Printf("%8d", d.lines)
	}
	if wordFlag {
		fmt.Printf("%8d", d.words)
	}
	if bytesFlag && !multibyteFlag {
		fmt.Printf("%8d", d.bytes)
	}
	if multibyteFlag {
		fmt.Printf("%8d", d.characters)
	}
	if lengthFlag {
		fmt.Printf("%8d", d.line)
	}
  if ( d.name != "" ) {
    fmt.Printf(" %-8s", d.name)
  }
  fmt.Printf("\n")
}

// noFlagsSet determines if any flags were set
func noFlagsSet() (result bool) {
  if !lengthFlag && !linesFlag && !wordFlag && !bytesFlag && ! multibyteFlag {
    result = true
  }
  return
}

// setDefaultFlags sets the appropriate flags when non are passed
func setDefaultFlags() {
  wordFlag  = true
  bytesFlag = true
  linesFlag = true
}

func main() {
	// Process arguments
	args.Parse()

  if noFlagsSet() {
    setDefaultFlags()
  }

	// Determine if data is being passed via STD_IN
	stat, _ := os.Stdin.Stat()
	if (stat.Mode()&os.ModeCharDevice) == 0 || len(filenames) == 0 {
		processStdin()
	} else {
		processFile()
	}
}
