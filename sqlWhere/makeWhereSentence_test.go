package sqlWhere

import (
	"testing"
)

var rightTestString = "$filter=(name eq 'Milk' or name eq 'Eggs') and price lt 2.55"

func TestOk(t *testing.T) {
	_, err := MakeWhereSentence(rightTestString)
	if err != nil {
		t.Errorf("TestOk failed")
	}
}

var wrongTestString = "$filter =(name eq 'Milk' or name eq 'Eggs') and price lt 2.55"

func TestError(t *testing.T) {
	_, err := MakeWhereSentence(wrongTestString)
	if err == nil {
		t.Errorf("TestError failed")
	}
}
