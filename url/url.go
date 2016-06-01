package url

import "github.com/asaskevich/govalidator"

type URL interface {
	URLString() string
	String() string
}

type url string

func (u url) URLString() string {
	return string(u)
}

func (u url) String() string {
	return string(u)
}

func NewOrNil(urlString string) URL {
	url, _ := New(urlString)
	return url
}

func New(urlString string) (URL, error) {
	if !govalidator.IsURL(urlString) {
		return nil, ErrInvalidURL
	}

	return url(urlString), nil
}
