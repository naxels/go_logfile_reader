# Log file reader written in Golang

Simple Golang logfile reader written to help read in a logfile
and do a RegEx to highlight words on a line

**TODO:**
- option to tell how many lines to read
- option to read from head or tail
- option to input a RegEx pattern which is used to highlight text in each line
- create tests (fulltext read, head + lines read, tail + lines read, RegEx pattern, lines to read larger than file (catch too many lines))


You can use a go routine to input a volume of log files and get lines back

## Installation
```
go get github.com/naxels/go_logfile_reader
```

## Basic Usage
```
import "github.com/naxels/go_logfile_reader"
// be sure to import fmt package for this example

//single file:
fileLocation := "filelocation/filename.extension"

//read file, fulltext, all of the file, no RegEx pattern
result, err := Read(fileLocation, "", 0, "")
if err != nil {
  fmt.Println(err)
}

fmt.Println(result.Filename)

..TODO.. process each line
```
