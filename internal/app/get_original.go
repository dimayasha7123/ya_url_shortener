package app

import "context"

func (sh *Shortener) GetOrig(shorten string) (string, error) {
	err := sh.validateShort(shorten)
	if err != nil {
		return "", NewValidationError(shorten, err)
	}

	id, err := sh.encoder.Decode(shorten)
	if err != nil {
		return "", NewDecodingError(shorten, err)
	}

	orig, err := sh.repo.GetOrigByID(context.Background(), id)
	if err != nil {
		return "", NewGettingOriginalError(id, err)
	}

	return orig, nil
}
