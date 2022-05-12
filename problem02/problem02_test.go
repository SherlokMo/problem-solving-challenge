package problem02_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestChangeFileName(t *testing.T) {
	before, after := "original.txt", "changed.txt"
	_, err := exec.Command("/bin/sh", "script.sh", before+" "+after).Output()
	if err != nil {
		t.Fatalf("Got an error with script.sh file, err: (%v)", err)
	}

	_, err = os.Stat(before)
	if err == nil {
		t.Errorf("Expected %v to not exist but found it", before)
	}

	_, err = os.Stat(after)
	if err != nil {
		t.Errorf("Expected (%v) to exist but didn't found it", after)
	}
}
