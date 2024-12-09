package dynamicCore

import (
	"errors"
)

type FileResource struct {
	Method       string
	Name         string
	CompletePath string
	Token        []string
	HttpCode     int
	Params       map[string]string
	Extension    string
	Error        *error
	IsValid      bool
}

func NewResource(parser Parser) (*FileResource, error) {
	res := FileResource{
		Name:         parser.GetEntry().Name(),
		CompletePath: parser.GetEntry().Name(),
		IsValid:      false,
	}

	if err := parser.Tokenize(&res); err != nil {
		return nil, err
	}
	if err := parser.ParseBody(&res); err != nil {
		return nil, err
	}
	if err := parser.ParseParams(&res); err != nil {
		res.Error = &err
	}
	if err := parser.ParseDynamic(&res); err != nil {
		return nil, err
	}
	if err := res.ValidateResource(); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *FileResource) ValidateResource() error {
	if r.Method == "" {
		return errors.New("method is missing")
	}
	if r.Name == "" {
		return errors.New("name is missing")
	}
	if r.HttpCode == 0 {
		return errors.New("http code is missing")
	}
	if r.Extension == "" {
		return errors.New("extension is missing")
	}
	r.IsValid = true
	if r.Error != nil {
		r.IsValid = false
	}
	return nil
}

func (r *FileResource) Is(request Request) error {
	if r.Method != request.Method {
		return errors.New("method mismatch")
	}
	return nil
}
