package main

import "testing"

func TestInitialConfig(t *testing.T) {
	c := &Config{}
	result := c.Init([]string{"hello", "world"}, false, "json")

	expected := []string{"hello", "world"}
	found := c.args
	if len(found) != len(expected) || found[0] != expected[0] || found[1] != expected[1] {
		t.Fatalf("Got %#v, Expected %#v", found, expected)
	}

	var response error
	if result != response {
		t.Fatalf("Got [%s], Expected [%s]", result, response)
	}
}
