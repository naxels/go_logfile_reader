# Log file reader written in Golang

Simple Golang logfile reader written to help read in a logfile
and do a RegEx to for example highlight matches on a line

**TODO:**
- implement the option to read from tail
- create tests tail + lines read, lines to read larger than file (catch more lines than 'file' error without user error))


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

//loop over each line
for _, line := range result.Lines {
  fmt.Println(line.Data)
}
```
