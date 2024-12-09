package tokenize

import (
	"errors"
	"fmt"
	"strings"
)

func TokenizeSlug(params string, paramDelimiter string, keyValueDelimiter string) (map[string]string, error) {
	tokenParams := make(map[string]string)
	params = strings.TrimSpace(params)
	tokens := strings.Split(params, paramDelimiter)
	for _, token := range tokens {
		keyValue := strings.Split(token, keyValueDelimiter)
		if len(keyValue) != 2 {
			return nil, errors.New("invalid parameter format: must be key:value")
		}
		key := keyValue[0]
		value := keyValue[1]
		if key == "" || value == "" {
			return nil, errors.New(fmt.Sprintf("invalid parameter: key or value is empty - %s", token))
		}
		tokenParams[key] = value
	}
	return tokenParams, nil
}
