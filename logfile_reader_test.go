package logfilereader

import (
	"os"
	"testing"
)

const (
	testDataDir    = "testdata"
	testFile       = "sample"
	testFileSuffix = ".txt"
)

type testpair struct {
	fileName          string
	readFrom          string
	linesToRead       int
	regExPattern      string
	expectedlineCount int
	expectedMatches   int
}

var tests = []testpair{
	{testFile + testFileSuffix, "", 0, "", 13, 0},              //read fulltext, all lines, no RegEx test
	{testFile + testFileSuffix, "", 0, "ERROR", 13, 1},         //read fulltest, all lines, RegEx test
	{testFile + testFileSuffix, "", 0, "ERROR|E=[1-9]", 13, 2}, //read fulltest, all lines, advanced RegEx test
	// {testFile + testFileSuffix, "", 0, "", 0, 0}, //template
}

func TestRead(t *testing.T) {
	for _, pair := range tests {
		result, err := Read(testDataDir+string(os.PathSeparator)+pair.fileName, pair.readFrom, pair.linesToRead, pair.regExPattern)
		if err != nil {
			t.Fatalf("Read(%q) err = %v, expected nil", testFile, err)
		}

		//test fileName match
		if result.Filename != pair.fileName {
			t.Fatalf("Read(%q) fileName = %v, expected %v", testFile, result.Filename, pair.fileName)
		}

		//test Lines in result
		if len(result.Lines) != pair.expectedlineCount {
			t.Fatalf("Read(%q) resultLines = %v, expected %v", testFile, len(result.Lines), pair.expectedlineCount)
		}

		resultMatches := 0
		for _, line := range result.Lines {
			if line.Matches != nil {
				resultMatches++
			}
		}
		//test Match count in result
		if resultMatches != pair.expectedMatches {
			t.Fatalf("Read(%q) resultMatches = %v, expected %v", testFile, resultMatches, pair.expectedMatches)
		}

		t.Logf("result: %+v", result)
	}
}
