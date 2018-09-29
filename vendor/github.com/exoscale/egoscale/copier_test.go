package egoscale

import (
	"testing"
)

func TestCopier(t *testing.T) {
	type Foo struct {
		Name string
	}

	type Bar struct {
		Name string
	}

	f := Foo{"Eggs"}
	b := Bar{"Spam"}

	if err := Copy(&f, b); err != nil {
		t.Error(err)
	}
	if f.Name != b.Name {
		t.Errorf("copy failed %q != %q", f.Name, b.Name)
	}

	b.Name = "Yo!"
	if err := Copy(&f, &b); err != nil {
		t.Error(err)
	}
	if f.Name != b.Name {
		t.Errorf("copy failed %q != %q", f.Name, b.Name)
	}

	if err := Copy(f, b); err == nil {
		t.Errorf("an error was expected")
	}
}

func TestCopierNonConvertible(t *testing.T) {
	type Foo struct {
		Name string
	}

	type Bar struct {
		Name string
		Age  int
	}

	f := Foo{"Eggs"}
	b := Bar{"Spam", 42}

	if err := Copy(&f, b); err == nil {
		t.Errorf("an error was expected")
	}

	if err := Copy(&b, f); err == nil {
		t.Errorf("an error was expected")
	}
}
