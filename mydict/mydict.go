package mydict

import "errors"

/*
Dictionary 의 경우 map[string]string의 가명(alias)이다.
아래와 같은 것을 type이라고 하는데
type의 경우 메서드를 가질 수 있다.
*/
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("Word Exists")

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

/*
Add 이곳에서 리서버가 포인터 리시버가 되어야 할 줄 알았으나
일반 리시버를 사용해도 잘 동작한다.
왜 그러하냐면 Dictionary 타입이 담고 있는 map타입은 기본적으로 포인터를 이용하기 때문이다.
*/
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}
