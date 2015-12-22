# Log file reader written in Golang

Simple Golang logfile reader written to help read in a logfile
and do a RegEx to highlight words on a line

**TODO:**
- input path + filename
- parse logfile (default: complete file)
- option to tell how many lines to read + head or tail
- option to input a RegEx pattern which is used to highlight text in each line
- create tests (head + lines, tail + lines, fullfile, RegEx pattern)


You can use a go routine to input a volume of log files and get lines back

## Installation
```
go get github.com/naxels/go_logfile_reader
```

## Basic Usage
```
import "github.com/naxels/go_logfile_reader"
// be sure to import fmt package for this example

TBD...
```
