// based on the post: https://blog.gopheracademy.com/advent-2014/parsers-lexers/

package filterparser

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Token int

const (
	// Special Tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	FIELD

	// String values between " or '
	VALUE

	// Group operators
	AND
	OR
	NOT

	// Operators
	LIKE
	NOT_LIKE
	EQUALS
	NOT_EQUALS

	// Group statement
	OPEN_PARENTHESES
	CLOSE_PARENTHESES
)

// Scanner represents a lexical scanner
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new Scanner instance
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

const eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isOperator(ch rune) bool {
	return ch == '!' || ch == '='
}

func isQuote(ch rune) bool {
	return ch == '"' || ch == '\''
}

// Scan returns the next Token and literal value.
func (s *Scanner) Scan() (Token, string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	}

	if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	if isOperator(ch) {
		s.unread()
		return s.scanOperator()
	}

	if isQuote(ch) {
		s.unread()
		return s.scanValue()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	}

	return ILLEGAL, string(ch)
}

func (s *Scanner) scanWhitespace() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		ch := s.read()
		if ch == eof {
			break
		}

		if !isWhitespace(ch) {
			s.unread()
			break
		}

		buf.WriteRune(ch)
	}

	return WS, buf.String()
}

func (s *Scanner) scanIdent() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		ch := s.read()
		if ch == eof {
			break
		}

		if !isLetter(ch) && !isDigit(ch) && ch != '_' && ch != '.' {
			s.unread()
			break
		}

		buf.WriteRune(ch)
	}

	switch strings.ToUpper(buf.String()) {
	case "NOT":
		return s.scanNotLike(buf)
	case "LIKE":
		return LIKE, buf.String()
	case "AND":
		return AND, buf.String()
	case "OR":
		return OR, buf.String()
	}

	return FIELD, buf.String()
}

func (s *Scanner) scanNotLike(buf bytes.Buffer) (Token, string) {
	_, ws := s.scanWhitespace()
	if ws == "" {
		return FIELD, buf.String()
	}

	tok, lit := s.scanIdent()
	buf.Write([]byte(lit))

	if tok != LIKE {
		return ILLEGAL, buf.String()
	}

	return NOT_LIKE, buf.String()
}

func (s *Scanner) scanOperator() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		ch := s.read()
		if ch == eof {
			break
		}

		if !isOperator(ch) {
			s.unread()
			break
		}

		buf.WriteRune(ch)
	}

	lit := buf.String()
	switch lit {
	case "!=":
		return NOT_EQUALS, lit
	case "==":
		return EQUALS, lit
	case "!":
		return NOT, lit
	}

	return ILLEGAL, lit
}

func (s *Scanner) scanValue() (Token, string) {
	var buf bytes.Buffer
	quote := s.read()
	buf.WriteRune(quote)

	for {
		ch := s.read()
		if ch == eof {
			return ILLEGAL, buf.String()
		}

		if ch == '\\' {
			ch = s.read()
			if ch == eof {
				buf.WriteRune('\\')
				return ILLEGAL, buf.String()
			}
			writeEscapedRune(ch, buf)
			continue
		}

		buf.WriteRune(ch)

		if ch == quote {
			break
		}
	}

	return VALUE, buf.String()
}

func writeEscapedRune(ch rune, buf bytes.Buffer) {
	switch ch {
	case 'n':
		buf.WriteRune('\n')
	case 't':
		buf.WriteRune('\t')
	case '\'':
		buf.WriteRune('\'')
	case '"':
		buf.WriteRune('"')
	case '\n':
		return
	default:
		buf.WriteRune('\\')
		buf.WriteRune(ch)
	}
}
