package formula

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type TokenType int

const (
	TokEOF TokenType = iota
	TokErr
	TokOpen
	TokClose
	TokComma
	TokNumber
	TokOperator
	TokReference
	TokFunc
)

type Token struct {
	TokType TokenType
	value   any
}

func (t Token) String() string {
	val := fmt.Sprint(t.value)
	switch t.value.(type) {
	case rune:
		val = fmt.Sprintf("%c", t.value)
	}
	return fmt.Sprintf("{%d %s}", t.TokType, val)
}

// returns true if token belongs to one of the arguments
// argument can only be: "*", "/", "+", "-", "(", ")", ","
// "eof", "ref", "num", "func"
func (t *Token) oneOf(types ...string) bool {
	res := []bool{}
	for _, s := range types {
		res = append(res, t.is(s))
	}
	return slices.Contains(res, true)
}

func (t *Token) is(s string) bool {
	switch s {
	case "+":
		return t.TokType == TokOperator && t.value == '+'
	case "-":
		return t.TokType == TokOperator && t.value == '-'
	case "*":
		return t.TokType == TokOperator && t.value == '*'
	case "/":
		return t.TokType == TokOperator && t.value == '/'
	case "(":
		return t.TokType == TokOpen && t.value == '('
	case ")":
		return t.TokType == TokClose && t.value == ')'
	case ",":
		return t.TokType == TokComma && t.value == ','
	case "eof":
		return t.TokType == TokEOF && t.value == nil
	case "ref":
		return t.TokType == TokReference
	case "num":
		return t.TokType == TokNumber
	case "func":
		return t.TokType == TokFunc
	}
	return false
}

func (t *Token) EOF() bool {
	return t.TokType == TokEOF && t.value == nil
}

type Lexer struct {
	reader *strings.Reader
}

func NewLexer(expr string) *Lexer {
	return &Lexer{
		reader: strings.NewReader(expr),
	}
}

func (l *Lexer) NextToken() Token {
	for {
		s, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return Token{TokEOF, nil}
			}
			panic(err)
		}
		switch {
		case s == ',':
			return Token{TokComma, s}
		case s == '(':
			return Token{TokOpen, s}
		case s == ')':
			return Token{TokClose, s}
		case s == '+':
			return Token{TokOperator, s}
		case s == '-':
			return Token{TokOperator, s}
		case s == '*':
			return Token{TokOperator, s}
		case s == '/':
			return Token{TokOperator, s}
		case isDigit(s) || s == '.':
			l.reader.UnreadRune()
			n, err := l.ReadNumber()
			if err != nil {
				return Token{TokErr, err.Error()}
			}
			return Token{TokNumber, n}
		case isLetter(s):
			l.reader.UnreadRune()
			id := l.ReadIdentifier()
			if isRef(id) {
				return Token{TokReference, id}
			}
			return Token{TokFunc, id}
		case s == ' ' || s == '\t' || s == '\r':
			continue
		}
	}
}

func (l *Lexer) ReadWhile(f func(rune) bool) bytes.Buffer {
	var buf bytes.Buffer
	for {
		if s, _, err := l.reader.ReadRune(); err == io.EOF {
			break
		} else if !f(s) {
			l.reader.UnreadRune()
			break
		} else {
			buf.WriteRune(s)
		}
	}
	return buf
}

func (l *Lexer) ReadNumber() (float64, error) {
	buf := l.ReadWhile(func(r rune) bool {
		return isDigit(r) || r == '.'
	})
	return strconv.ParseFloat(buf.String(), 64)
}

func (l *Lexer) ReadIdentifier() string {
	buf := l.ReadWhile(func(r rune) bool {
		return isLetter(r) || isDigit(r)
	})
	return buf.String()
}

func isRef(id string) bool {
	res, _ := regexp.Match("^[A-Z]+\\d+$", []byte(id))
	return res
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}
