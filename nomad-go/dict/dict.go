package dict

import "errors"

// key(string) : value(string)으로 선언한 map에 type을 사용해서 alias를 부여할 수 있다.
// 그리고 객체만이 아니라 alias가 부여된 타입에도 메소드를 작성할 수 있다.
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errKeyExists = errors.New("Already Exists")

// 위 Map의 활용을 위해 메소드를 정의할 수 있다.
func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case errNotFound:
		d[key] = value
	case nil:
		return errKeyExists
	}
	return nil
}