package maps

import "errors"

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

type Dictionary map[string]string

func (d Dictionary) SearchA(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
