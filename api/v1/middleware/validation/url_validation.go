package validation

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/ccallazans/url-shortener/internal/domain/models"
	"github.com/ccallazans/url-shortener/internal/utils"
)

type UrlRequestValidation interface {
	Execute(urlRequest *models.UrlRequest) error
}

//

type UrlValidation struct {
	Next UrlRequestValidation
}

func (v *UrlValidation) Execute(urlRequest *models.UrlRequest) error {

	if urlRequest.Url == "" {
		return errors.New(utils.URL_EMPTY_ERROR)
	}

	u, err := url.Parse(urlRequest.Url)
	if err != nil {
		return errors.New(utils.URL_INVALID_ERROR)
	}

	fmt.Println(u)

	if !u.IsAbs() || (u.Scheme != "http" && u.Scheme != "https") {
		return errors.New(utils.URL_INVALID_ERROR)
	}

	return nil
}
