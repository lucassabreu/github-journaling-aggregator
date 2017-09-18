// based on: https://blog.gopheracademy.com/advent-2014/parsers-lexers/
package filterparser

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/lucassabreu/github-journaling-aggregator/filter"
)

var tokenName = map[Token]string{
	FIELD:             "FIELD",
	VALUE:             "VALUE",
	AND:               "AND",
	OR:                "OR",
	NOT:               "NOT",
	NOT_LIKE:          "NOT LIKE",
	LIKE:              "LIKE",
	EQUALS:            "EQUALS",
	NOT_EQUALS:        "NOT EQUALS",
	OPEN_PARENTHESES:  "OPEN PARENTHESES",
	CLOSE_PARENTHESES: "CLOSE PARENTHESES",
}

// Parser represents a parser
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	for {
		tok, lit = p.scan()
		if tok != WS {
			return
		}
	}
}

// Parse will read the content and return a Filter instance with it
func (p *Parser) Parse() (filter.Filter, error) {
	fg, err := p.parseUntil(EOF)
	return fg, err
}

func (p *Parser) parseUntil(endToken Token) (fg *filter.FilterGroup, err error) {
	fg = filter.NewFilterGroup()
	var f filter.Filter

	for {
		tok, lit := p.scan()

		switch tok {
		case OPEN_PARENTHESES:
			f, err = p.parseUntil(CLOSE_PARENTHESES)
			if err != nil {
				fg = nil
				return
			}
		case FIELD, VALUE:
			p.unscan()
			f, err = p.parseClause()
		}

		if f == nil {
			tokName, ok := tokenName[tok]
			if !ok {
				tokName = "UNKNOWN"
			}
			err = fmt.Errorf("Expected (, field or value, received: %s (type: %s)", lit, tokName)
		}

		fg.Append(f)

		if tok == endToken {
			break
		}
	}

	return
}

func (p *Parser) parseClause() (filter.Filter, error) {
	lTok, lLit := p.scan()
	tokOperator, opLit := p.scan()
	rTok, rLit := p.scan()

	if tokOperator != EQUALS || tokOperator != NOT_EQUALS || tokOperator != NOT_LIKE || tokOperator != LIKE {
		tokName, ok := tokenName[tokOperator]
		if !ok {
			tokName = "UNKNOWN"
		}
		return nil, fmt.Errorf("Expected operator received: %s (type: %s)", opLit, tokName)
	}

	if lTok == rTok {
		return nil, errors.New("You can not compare two fields. Must be a field and a value in a clause !")
	}

	var field string
	var value string

	switch lTok {
	case FIELD:
		field, value = lLit, rLit
	case VALUE:
		field, value = rLit, lLit
	}

	value = value[1:][:len(value)-2]

	switch strings.ToLower(field) {
	case "repo.name", "repo", "repository":
		switch tokOperator {
		case EQUALS:
			return filter.NewEqualsRepository(value), nil
		case NOT_EQUALS:
			return filter.NewNot(filter.NewEqualsRepository(value)), nil
		case LIKE:
			return filter.NewRepositoryNameRegExpFilter(regexp.MustCompile(field)), nil
		case NOT_LIKE:
			return filter.NewNot(filter.NewRepositoryNameRegExpFilter(regexp.MustCompile(value))), nil
		}
	default:
		return nil, fmt.Errorf("Unknown field: %s", field)
	}

	return nil, nil
}
