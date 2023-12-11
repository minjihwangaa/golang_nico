package mydict

import (
	"errors"
)

// Dictionary type => map[string]string 에 대한 alias(가명)
type Dictionary map[string]string

// 2.4 Search for word
var errNotFound = errors.New("not found")
func (d Dictionary) Search(word string) (string, error) {
    value, exists := d[word]


    if exists {
        return value, nil // return string, error
    }
    return "", errNotFound  // return string, error
}

// 2.5 Add Method - Add word to dictionary
var errWordExists = errors.New("this word already exist")
func (d Dictionary) Add(word, def string) error {
    _, err := d.Search(word)

    // If 문
    // if err == errNotFound {
    //  d[word] = def
    // }else if err == nil {
    //  return errWordExists
    // }
    // return nil
    switch err {
    case errNotFound:
        d[word] = def
    case nil:
        return errWordExists
    }
    return nil
}

// 2.6 Update 
var errCantUpdate = errors.New("this word can't update")
func (d Dictionary) Update(word, defination string) error {
    _, err := d.Search(word)
    switch err {
    case nil: 
        d[word] = defination
    case errNotFound:
        return errCantUpdate
    }
    return nil
}

// Delete
func (d Dictionary) Delete(word string) error  {
    _, err := d.Search(word)
    //delete(d, word)
    if err == nil {
        delete(d, word)
    }
    return errNotFound
    
}
