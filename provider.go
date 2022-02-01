package idf

import (
	"fmt"
	"strings"
)

type IDProvider struct {
	length     int
	padChar    string
	normalized bool
}

func New(opts ...func(*IDProvider)) *IDProvider {
	p := &IDProvider{}
	for _, o := range opts {
		o(p)
	}
	return p
}

func WithNormalized() func(*IDProvider) {
	return func(p *IDProvider) {
		p.normalized = true
	}
}

func WithPadding(char string, length int) func(*IDProvider) {
	return func(p *IDProvider) {
		p.padChar = char
		p.length = length
	}
}

func (p *IDProvider) ToID(v string) (string, error) {
	if p.normalized {
		v = strings.TrimSpace(strings.ToLower(v))
	}

	// if len not set or already equal to the target size just return
	if p.length == 0 || len(v) == p.length {
		return v, nil
	}

	// if longer than the desired length, return substring
	if len(v) > p.length {
		runes := []rune(v)
		return string(runes[:p.length]), nil
	}

	// if shorter than pad to the the desired length and return that
	return fmt.Sprintf("%s%s", v, strings.Repeat(p.padChar, p.length-len(v))), nil
}
