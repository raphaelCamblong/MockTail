package dynamicCore

import (
	"errors"
	"mockTail/pkg/tokenize"
	"os"
	"strconv"
	"strings"
)

type Parser interface {
	SetCurrentDirEntry(file os.DirEntry) error
	GetEntry() os.DirEntry
	Tokenize(res *FileResource) error
	ParseParams(res *FileResource) error
	ParseBody(res *FileResource) error
	ParseDynamic(res *FileResource) error
}

const (
	paramIndex        = 1
	paramDelimiter    = ","
	keyValueDelimiter = ":"
	openBracket       = "["
	closeBracket      = "]"
)

var ErrParamMissing = errors.New("parameter token is missing")

type parser struct {
	entry os.DirEntry
}

func NewParser(entry os.DirEntry) Parser {
	return &parser{entry: entry}
}

func (p *parser) SetCurrentDirEntry(file os.DirEntry) error {
	if file == nil {
		return errors.New("file entry is nil")
	}
	p.entry = file
	return nil
}

func (p *parser) GetEntry() os.DirEntry {
	return p.entry
}

func (p *parser) Tokenize(r *FileResource) error {
	if p.entry == nil {
		return errors.New("entry is not set")
	}
	r.Token = strings.Split(p.entry.Name(), ".")
	return nil
}

func (p *parser) ParseParams(r *FileResource) error {
	if err := p.validateRawParam(r); err != nil {
		if errors.Is(err, ErrParamMissing) {
			return nil
		}
		return err
	}
	paramString := &r.Token[paramIndex]

	if *paramString == "" {
		return nil
	}

	paramPairs, err := tokenize.TokenizeParams(*paramString, paramDelimiter, keyValueDelimiter)
	if err != nil {
		return err
	}
	r.Params = paramPairs
	return nil
}

func (p *parser) validateRawParam(r *FileResource) error {
	if r.Token == nil {
		return errors.New("invalid token: nil")
	}
	if len(r.Token) <= 3 {
		return ErrParamMissing
	}
	paramString := &r.Token[paramIndex]

	if !strings.HasPrefix(*paramString, openBracket) || !strings.HasSuffix(*paramString, closeBracket) {
		return errors.New("invalid parameter format: missing enclosing brackets")
	}

	*paramString = strings.Trim(*paramString, openBracket+closeBracket)
	return nil
}

func (p *parser) ParseBody(r *FileResource) error {
	httpCode, err := strconv.Atoi(r.Token[len(r.Token)-2])
	if err != nil {
		return errors.New("file name must have a valid HTTP code at len-1")
	}
	r.Method = r.Token[0]
	r.HttpCode = httpCode
	r.Extension = r.Token[len(r.Token)-1]
	return nil
}

func (p *parser) ParseDynamic(r *FileResource) error {
	return nil
}
