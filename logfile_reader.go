package logfilereader

import "os"

var (
	l Logfile
)

//Logfile struct for file
type Logfile struct {
	Filename string `json:"filename"`
	Lines    []Line `json:"lines"`
}

//Line struct for file lines
type Line struct {
	Data string `json:"data"`
}

//Read returns a Logfile struct after opening file and processing it
func Read(fileLocation string, readFrom string, linesToRead int, regExPattern string) (*Logfile, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return &l, err
	}

	fileInfo, _ := file.Stat()
	l.Filename = fileInfo.Name()

	//TODO process file

	file.Close()

	return &l, nil
}
