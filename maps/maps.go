package maps

type Dictionary map[string]string
type DictionaryErr string

// https://dave.cheney.net/2016/04/07/constant-errors
var (
	ErrNotFound   = DictionaryErr("the key does not exist")
	ErrWordExists = DictionaryErr("word exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(key string, value string) {
	d[key] = value
}
