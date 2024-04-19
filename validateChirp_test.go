package main

import (
	"fmt"
	"testing"
)

func TestRmProfane(t *testing.T) {
	tests := []struct {
		chirp    string
		expected string
	}{
		{"Hello! are you a kerfuffle ?", "Hello! are you a **** ?"},
		{"This is a kerfuffle opinion I need to share with the world", "This is a **** opinion I need to share with the world"},
		{"I had something interesting for breakfast", "I had something interesting for breakfast"},
		{"I hear Mastodon is better than Chirpy. sharbert I need to migrate", "I hear Mastodon is better than Chirpy. **** I need to migrate"},
		{"I really need a kerfuffle to go to bed sooner, Fornax !", "I really need a **** to go to bed sooner, **** !"},
	}
	for _, test := range tests {
		if actual := rmProfane(test.chirp); actual != test.expected {
			t.Errorf("Test Failed: chirp: %s, expected: %s, got: %s\n", test.chirp, test.expected, actual)
		} else {
			fmt.Printf("Test Passed: chirp: %s, expected: %s, got: %s\n", test.chirp, test.expected, actual)
		}
	}
}
