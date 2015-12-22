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

func TestRead(t *testing.T) {
	//standard read fulltext, all lines, no RegEx test
	result, err := Read(testDataDir+string(os.PathSeparator)+testFile+testFileSuffix, "", 0, "")
	if err != nil {
		t.Fatalf("Read(%q) err = %v, expected nil", testFile, err)
	}

	//test fileName match
	if result.Filename != testFile+testFileSuffix {
		t.Fatalf("Read(%q) fileName = %v, expected %v", testFile, result.Filename, testFile+testFileSuffix)
	}

	t.Logf("result: %+v", result)
}
