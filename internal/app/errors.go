package app

import "fmt"

type ValidationError struct {
	url string
	err error
}

func NewValidationError(url string, err error) ValidationError {
	return ValidationError{url: url, err: err}
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("short url %q not valid: %v", e.url, e.err)
}

type DecodingError struct {
	url string
	err error
}

func NewDecodingError(url string, err error) DecodingError {
	return DecodingError{url: url, err: err}
}

func (e DecodingError) Error() string {
	return fmt.Sprintf("can't decode url %q: %v", e.url, e.err)
}

type GettingOriginalError struct {
	id  uint64
	err error
}

func NewGettingOriginalError(id uint64, err error) GettingOriginalError {
	return GettingOriginalError{id: id, err: err}
}

func (e GettingOriginalError) Error() string {
	return fmt.Sprintf("can't get original url from repo by id = %v: %v", e.id, e.err)
}

type GettingShortenError struct {
	url string
	err error
}

func NewGettingShortenError(url string, err error) GettingShortenError {
	return GettingShortenError{url: url, err: err}
}

func (e GettingShortenError) Error() string {
	return fmt.Sprintf("can't get shorten url from repo by original url %q: %v", e.url, e.err)
}

type SavingNewURLError struct {
	id    uint64
	encID string
	orig  string
	err   error
}

func NewSavingNewURLError(id uint64, encID, orig string, err error) SavingNewURLError {
	return SavingNewURLError{
		id:    id,
		encID: encID,
		orig:  orig,
		err:   err,
	}
}

func (e SavingNewURLError) Error() string {
	return fmt.Sprintf("can't save new url to repo with id = %q, encID = %q, orig = %q: %v",
		e.id,
		e.encID,
		e.orig,
		e.err,
	)
}
