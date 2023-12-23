package jScan

import (
	"bufio"
	"log"
	"os"
)

type Scanner struct {
	scanner *bufio.Scanner
}

func NewScanner() *Scanner {
	m := &Scanner{
		scanner: bufio.NewScanner(os.Stdin),
	}
	return m
}

func (m *Scanner) ReadLine() string {
	m.scanner.Scan()
	err := m.scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return m.scanner.Text()
}
