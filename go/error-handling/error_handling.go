package erratum

func Use(opener ResourceOpener, input string) (errP error) {
	ressource, errO := opener()

	for _, ok := errO.(TransientError); ok; _, ok = errO.(TransientError) {
		ressource, errO = opener()
	}
	if errO != nil {
		return errO
	}

	defer ressource.Close()
	defer func() {
		if r := recover(); r != nil {
			if fe, ok := r.(FrobError); ok {
				ressource.Defrob(fe.defrobTag)
			}
			errP = r.(error)
		}
	}()

	ressource.Frob(input)
	return
}
