package v1

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFilePathJoin(t *testing.T) {
	p1 := "."
	name := "fa.txt"
	path := filepath.Join(p1, name)
	fmt.Printf("%s\n", path)
	path, _ = filepath.Abs("v1_test.go")
	fmt.Printf("%s\n", path)
}

func TestAbs(t *testing.T) {
	path, err := filepath.Abs("./assets")
	if err != nil {
		t.Errorf("%s\n", err)
	}
	fmt.Printf("%s\n", path)

}
