package xmlpath

import (
	"errors"
	"strings"
)

// FindAllString finds all matching node's string values
// Func names inspired by the std regexp package names
func FindAllString(xml, xpath string) ([]string, error) {
	path, err := Compile(xpath)
	if err != nil {
		return nil, err
	}

	root, err := Parse(strings.NewReader(xml))
	if err != nil {
		return nil, err
	}

	ss := []string{}

	iter := path.Iter(root)
	for iter.Next() {
		s := iter.Node().String()
		ss = append(ss, strings.TrimSpace(s))
	}

	return ss, nil
}

// Finds first matching node's string value
// Use when there's only one expected matching node
func FindString(xml, xpath string) (string, error) {
	ss, err := FindAllString(xml, xpath)
	if err != nil {
		return "", err
	}

	if len(ss) == 0 {
		return "", errors.New("No Strings found for given xpath")
	}

	return ss[0], nil
}
