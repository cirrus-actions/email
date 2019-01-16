package internal

import (
	"testing"
)

func Test_SimpleTemplateRendering(t *testing.T) {
	data := struct {
		Foo string
	}{
		Foo: "bar",
	}
	result, err := Render("Foo={{.Foo}}", data)
	if err != nil {
		t.Error(err)
	}
	if result != "Foo=bar" {
		t.Errorf("Wrong render output: %s", result)
	}
}
