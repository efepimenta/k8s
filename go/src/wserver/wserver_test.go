package main

import "testing"

func TestGreeting(t *testing.T) {
	total := greeting("Code.education Rocks!");
	if total != "<b>Code.education Rocks!</b>" {
		t.Errorf("Saída incorreta, recebi: %s, esperado: %s.", total, "<b>Code.education Rocks!</b>")
	}
}
