package factory

import (
	"errors"
	"myapi/internal/app/domain"
	"myapi/internal/app/shared"
	pkgurl "net/url"

	"github.com/google/uuid"
)

func NewShortenerFactory(url string, hash string, user uuid.UUID) (domain.Shortener, error) {

	verify := urlVerify{}
	err := verify.execute(url)
	if err != nil {
		return domain.Shortener{}, err
	}

	return domain.Shortener{
		Url:  url,
		Hash: hash,
		User: user,
	}, nil
}

type ShortenerValidator interface {
	execute(url string)
	setNext(ShortenerValidator)
}

//

type urlVerify struct {
	next ShortenerValidator
}

func (r *urlVerify) setNext(next ShortenerValidator) {
	r.next = next
}

func (v *urlVerify) execute(url string) error {
	_, err := pkgurl.ParseRequestURI(url)
	if err != nil {
		return errors.New(shared.URL_INVALID_ERROR)
	}

	return nil
}
