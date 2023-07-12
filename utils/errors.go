package utils

import "errors"

var (
	ErrConversion = errors.New("error while converting")
	ErrOpenDB     = errors.New("error while opening db")
	ErrGetDB      = errors.New("error while getting db")
	ErrGetKarma   = errors.New("error while getting karmaCmd")
)
