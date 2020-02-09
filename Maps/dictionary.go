package main

const (
	// ErrNotFound means the definition could not be found
	ErrNotFound = DictionaryErr("I couldn't find the word you were looking for")

	// ErrWordExists means that you are trying to add a repeated word
	ErrWordExists = DictionaryErr("You cannot add this word because it already exists")

	// ErrWordDoesNotExist occurs when trying to update a word that does not exist in the dictionary
	ErrWordDoesNotExist = DictionaryErr("You cannot update this word because it does not exist")
)

// DictionaryErr are errors that can happen when interacting with the dictionary
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary store definitions of words
type Dictionary map[string]string

// Search find a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add inserts a word and definition into the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		return ErrWordExists
	case ErrNotFound:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Update changes the definition of a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err

	}

	return nil
}

// Delete removes a word of the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func main() {}
