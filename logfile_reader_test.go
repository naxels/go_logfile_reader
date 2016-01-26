package logfilereader

import (
	"os"
	"reflect"
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
	compareResult     int
	expectedResult    []Line
}

var tests = []testpair{
	{testFile + testFileSuffix, "", 0, "", 13, 0, 0, []Line{}},                                                                 //read fulltext, all lines, no RegEx test
	{testFile + testFileSuffix, "", 0, "ERROR", 13, 1, 0, []Line{}},                                                            //read fulltest, all lines, RegEx test
	{testFile + testFileSuffix, "", 0, "ERROR|E=[1-9]", 13, 2, 0, []Line{}},                                                    //read fulltest, all lines, advanced RegEx test
	{testFile + testFileSuffix, "", 10, "", 10, 0, 0, []Line{}},                                                                //read first 10 lines
	{testFile + testFileSuffix, "head", 1, "", 1, 0, 1, []Line{Line{Data: "2015/12/22 07:38:06 - Kitchen - Start of run.\n"}}}, //read 1 line from the head
	// {testFile + testFileSuffix, "", 0, "", 0, 0, 0, []Line{}}, //template
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
		if result.Count() != pair.expectedlineCount {
			t.Fatalf("Read(%q) resultLines = %v, expected %v", testFile, result.Count(), pair.expectedlineCount)
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

		//test exact result
		if pair.compareResult == 1 {
			if reflect.DeepEqual(result.Lines, pair.expectedResult) == false {
				t.Fatalf("Read(%q) got = %v, expected %v", testFile, result.Lines, pair.expectedResult)
			}
		}

		t.Logf("result: %+v", result)
	}
}
