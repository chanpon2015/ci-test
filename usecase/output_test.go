package usecase_test

import (
	"testing"

	"github.com/chanpon2015/ci-test/usecase"
)

func TestOut(t *testing.T) {
	o := usecase.NewOutput()
	if err := o.Out(); err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
