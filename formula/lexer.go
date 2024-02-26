package data

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
	TokColon
	TokNumber
	TokOperator
	TokReference
	TokFunc
	TokBool
)

type Token struct {
	T       TokenType
	literal string
}

func (t Token) String() string {
	return fmt.Sprintf("tok:{%#v %s}", t.T, t.literal)
}

// returns true if token belongs to one of the arguments
// argument can only be: "*", "/", "+", "-", "(", ")", ",", ":",
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
		return t.T == TokOperator && t.literal == s
	case "-":
		return t.T == TokOperator && t.literal == s
	case "*":
		return t.T == TokOperator && t.literal == s
	case "/":
		return t.T == TokOperator && t.literal == s
	case "(":
		return t.T == TokOpen
	case ")":
		return t.T == TokClose
	case ",":
		return t.T == TokComma
	case ":":
		return t.T == TokColon
	case "eof":
		return t.T == TokEOF
	case "ref":
		return t.T == TokReference
	case "num":
		return t.T == TokNumber
	case "func":
		return t.T == TokFunc
	case "bool":
		return t.T == TokBool
	}
	return false
}

func (t *Token) EOF() bool {
	return t.T == TokEOF
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
				return Token{TokEOF, ""}
			}
			panic(err)
		}
		switch {
		case s == ',':
			return Token{TokComma, string(s)}
		case s == ':':
			return Token{TokColon, string(s)}
		case s == '(':
			return Token{TokOpen, string(s)}
		case s == ')':
			return Token{TokClose, string(s)}
		case s == '+':
			return Token{TokOperator, string(s)}
		case s == '-':
			return Token{TokOperator, string(s)}
		case s == '*':
			return Token{TokOperator, string(s)}
		case s == '/':
			return Token{TokOperator, string(s)}
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
			} else if id == "TRUE" || id == "FALSE" {
				return Token{TokBool, id}
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

func (l *Lexer) ReadNumber() (string, error) {
	buf := l.ReadWhile(func(r rune) bool {
		return isDigit(r) || r == '.'
	})
	if _, err := strconv.ParseFloat(buf.String(), 64); err != nil {
		return buf.String(), err
	}
	return buf.String(), nil
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
