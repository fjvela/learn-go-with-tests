package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("the key does not exist")
	ErrWordExists = errors.New("word exists")
)

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Search(key string) (string, error) {
	result, exits := d[key]

	if !exits {
		return "", ErrNotFound
	}

	return result, nil
}
