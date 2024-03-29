package mttools

import (
	"reflect"
	"testing"
)

func TestSplitArgumentsString(t *testing.T) {
	strToParse := "/test 10 мяу fghfgh \"sdfg ert\" AAA-BBB AAA_BBB !ha"
	result := SplitArgumentsString(strToParse)

	expectedResult := []string{"/test", "10", "мяу", "fghfgh", "sdfg ert", "AAA-BBB", "AAA_BBB", "!ha"}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("EXPECTED: %#v\nGOT:%#v", expectedResult, result)
	}
}
