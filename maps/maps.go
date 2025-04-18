package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	result, exits := d[key]

	if !exits {
		return "", errors.New("the key does not exist")
	}
	return result, nil
}
