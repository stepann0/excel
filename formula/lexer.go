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
	TokColon
	TokNumber
	TokOperator
	TokReference
	TokFunc
	TokBool
)

func (t TokenType) String() string {
	switch t {
	case TokEOF:
		return "eof"
	case TokErr:
		return "err"
	case TokOpen:
		return "lparen"
	case TokClose:
		return "rparen"
	case TokComma:
		return "comma"
	case TokColon:
		return "colon"
	case TokNumber:
		return "num"
	case TokOperator:
		return "op"
	case TokReference:
		return "ref"
	case TokFunc:
		return "func"
	case TokBool:
		return "bool"
	}
	return ""
}

type Token struct {
	T       TokenType
	literal string
}

func (t Token) String() string {
	return fmt.Sprintf("tok:{%s '%s'}", t.T, t.literal)
}

// returns true if token belongs to one of the arguments
// argument can only be: "*", "/", "+", "-", "(", ")", ",", ":",
// ">", "<", "=", ">=", "<=", "<>",
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
	case ">", "<", "=", ">=", "<=", "<>":
		return t.T == TokOperator && t.literal == s
	case "+", "-", "*", "/":
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
		case s == '>', s == '<', s == '=':
			return l.comparisonOperator(s)
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

func (l *Lexer) comparisonOperator(first_half rune) Token {
	second_half, _, err := l.reader.ReadRune()
	if err == io.EOF || second_half != '<' && second_half != '>' && second_half != '=' {
		op := Token{TokOperator, string(first_half)}
		l.reader.UnreadRune()
		return op
	}
	// err != EOF and second rune is <, > or =
	if err != nil {
		panic(err)
	}

	op := string(first_half) + string(second_half)
	switch op {
	case ">=", "<=", "<>":
		return Token{TokOperator, op}
	}
	panic("invalid comparison operator")
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
