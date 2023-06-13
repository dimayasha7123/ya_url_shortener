package app

import "context"

func (sh *Shortener) Shorten(orig string) (string, error) {
	err := sh.validateOrig(orig)
	if err != nil {
		return "", NewValidationError(orig, err)
	}

	short, err := sh.repo.GetShortenByOrigIfExists(context.Background(), orig)
	if err != nil {
		return "", NewGettingShortenError(orig, err)
	}

	if short != "" {
		return short, nil
	}

	id := sh.idGen.Generate()
	encID := sh.encoder.Encode(id)

	err = sh.repo.SaveNewURL(context.Background(), id, encID, orig)
	if err != nil {
		return "", NewSavingNewURLError(id, encID, orig, err)
	}

	return encID, nil
}
