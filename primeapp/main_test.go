package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestIsPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func TestIntro(t *testing.T) {
	introTests := []struct {
		"Is it Prime?",
		"------------",
		"Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.",
	}
	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = w

	intro()

	w.Close()
	os.Stdout = oldStdout

	var output string
	io.Copy(&output, r)
	for _, e := range introTests {
		if !strings.Contains(output, e) {
			t.Errorf("%s is expected, but get a wrong answer", e)
		}
	}
}

func TestPrompt(t *testing.T) {
	promptTest := "-> "
	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = w

	prompt()

	w.Close()
	os.Stdout = oldStdout

	var output string
	io.Copy(&output, r)
	if !strings.Contains(output, e) {
		t.Errorf("%s is expected, but get a wrong answer", e)
	}
}

func TestCheckNumbers(t *testing.T) {
	testCases := []struct {
        input   string
        output  string
        returned bool
    }{
        {input: "2", output: "2 is a prime number!", returned: false},
        {input: "20", output: "20 is not a prime because it is divisible by 2!", returned: false},
        {input: "n", output: "Please enter a whole number!", returned: false},
        {input: "q", output: "", returned: true},
    }

    for _, e := range testCases {
        reader := strings.NewReader(e.input)
        scanner := bufio.NewScanner(reader)
        res, returned := checkNumbers(scanner)

        if res != e.output {
		t.Errorf("%s: %s expected, but got %s", e.input, e.output, res)
        }

        if done != e.returned {
		t.Errorf("%s: %v expected, but got %v", e.input, e.returned, returned)
        }
    }
}

func TestReadUserInput(t *testing.T) {
    doneChan := make(chan bool)

    go func() {
        defer close(doneChan)

        reader := strings.NewReader("2\n4\nq\n")
        readUserInput(reader, doneChan)
    }()

    select {
    case <-doneChan:
    case <-time.After(time.Second * 5):
        t.Error("Timed out waiting for doneChan to be closed")
    }
}
