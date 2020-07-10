// Package to test build.go file
package thing_test

import (
	"testing"
	"thing"
)

func TestCommand(t *testing.T) {
	t.Parallel()
	want := "You successfully ran a Linux command from Go!!!\n"
	got, err := thing.Command("echo", "You successfully ran a Linux command from Go!!!")
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestShell(t *testing.T) {
	t.Parallel()
	want := "You successfully ran a Linux command from Go!!!\n"
	got, err := thing.Shell("testdata/echo.sh")
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
