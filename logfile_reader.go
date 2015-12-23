package logfilereader

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

//Logfile struct for file
type Logfile struct {
	Filename string `json:"filename"`
	Lines    []Line `json:"lines"`
}

//Add Line to Logfile
func (l *Logfile) Add(data Line) {
	l.Lines = append(l.Lines, data)
}

//Line struct for file lines
type Line struct {
	Data    string  `json:"data"`
	Matches [][]int `json:"matches"`
}

//Read returns a Logfile struct after opening file and processing it
// readFrom can be: "" (defaults to fulltext), "head", "tail"
// linesToRead can be: 0 (defaults to entire file), or > 0
// regExPattern must the a RegEx parsable string, and looks for all matches in a line
func Read(fileLocation string, readFrom string, linesToRead int, regExPattern string) (*Logfile, error) {
	var l Logfile

	//set configuration
	if readFrom == "" {
		readFrom = "fulltext"
	}

	//test RegEx Compilation
	r, err := regexp.Compile(regExPattern)
	if err != nil {
		return &l, err
	}

	//open file
	file, err := os.Open(fileLocation)
	if err != nil {
		return &l, err
	}

	fileInfo, _ := file.Stat()
	l.Filename = fileInfo.Name()

	reader := bufio.NewReader(file)

	for {
		fileline, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return &l, err
		}

		//ensuring an empty string doesn't match everything
		if regExPattern != "" {
			l.Add(Line{fileline, r.FindAllStringIndex(fileline, -1)})
		} else {
			l.Add(Line{fileline, nil})
		}
	}

	file.Close()

	return &l, nil
}
