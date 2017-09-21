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
	ILLEGAL:           "ILLEGAL",
	EOF:               "EOF",
	WS:                "WS",
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
	return p.parseUntil(EOF)
}

func (p *Parser) parseUntil(endToken Token) (last filter.Filter, err error) {

	for {
		tok, lit := p.scanIgnoreWhitespace()

		switch tok {
		case EOF, endToken:
			return
		case AND, OR:
			if last == nil {
				return nil, unexpectedToken(tok, lit, "(, field or value")
			}

			p.unscan()
			last, err = p.parseGroup(last)
		case NOT:
			last, err = p.parseNot()
		case OPEN_PARENTHESES, FIELD, VALUE:
			p.unscan()
			last, err = p.parseCommon()
		default:
			err = unexpectedToken(tok, lit, "(, AND, OR, field or value")
		}

		if err != nil {
			return
		}
	}

	return last, nil
}

func unexpectedToken(tok Token, lit, expected string) error {
	tokName, ok := tokenName[tok]
	if !ok {
		tokName = "UNKNOWN"
	}
	return fmt.Errorf("Expected %s, received: %s (type: %s)", expected, lit, tokName)
}

func (p *Parser) parseGroup(prev filter.Filter) (filter.Filter, error) {
	var fg filter.FilterGroup
	switch tok, lit := p.scanIgnoreWhitespace(); tok {
	case AND:
		fg = filter.NewAndGroup()
	case OR:
		fg = filter.NewOrGroup()
	default:
		return nil, unexpectedToken(tok, lit, "AND or OR")
	}

	next, err := p.parseCommon()
	if err != nil {
		return nil, err
	}

	fg.Append(prev, next)
	return fg, nil
}

func (p *Parser) parseCommon() (filter.Filter, error) {
	switch tok, lit := p.scanIgnoreWhitespace(); tok {
	case FIELD, VALUE:
		p.unscan()
		return p.parseClause()
	case OPEN_PARENTHESES:
		return p.parseUntil(CLOSE_PARENTHESES)
	default:
		return nil, unexpectedToken(tok, lit, "(, field or value")
	}
}

func (p *Parser) parseNot() (filter.Filter, error) {
	f, err := p.parseCommon()
	if err != nil {
		return nil, err
	}

	return filter.NewNot(f), nil
}

func (p *Parser) parseClause() (filter.Filter, error) {
	lTok, lLit := p.scanIgnoreWhitespace()
	tokOperator, opLit := p.scanIgnoreWhitespace()
	rTok, rLit := p.scanIgnoreWhitespace()

	if tokOperator != EQUALS && tokOperator != NOT_EQUALS && tokOperator != NOT_LIKE && tokOperator != LIKE {
		return nil, unexpectedToken(tokOperator, opLit, "operator")
	}

	if lTok == rTok {
		return nil, errors.New("You can not compare two fields or values. Must be a field and a value in a clause !")
	}

	var field, value string

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
			return filter.NewRepositoryNameRegExpFilter(regexp.MustCompile(value)), nil
		case NOT_LIKE:
			return filter.NewNot(filter.NewRepositoryNameRegExpFilter(regexp.MustCompile(value))), nil
		}
	default:
		return nil, fmt.Errorf("Unknown field: %s", field)
	}
	return nil, fmt.Errorf("Could not undestand field %s", field)
}
