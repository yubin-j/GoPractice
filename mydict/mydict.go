package mydict

import "errors"

/*
Dictionary 의 경우 map[string]string의 가명(alias)이다.
아래와 같은 것을 type이라고 하는데
type의 경우 메서드를 가질 수 있다.
*/
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

/*
Search struct의 메소드와 동일하게 리시버를 받아줘야 한다.
*/
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound

}
