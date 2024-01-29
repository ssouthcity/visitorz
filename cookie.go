package main

import (
	"net/http"
	"strconv"
)

const CookieName = "visitorNumber"

func readVisitorCookie(r *http.Request) (int, bool) {
	c, err := r.Cookie(CookieName)
	if err != nil {
		return 0, false
	}

	visitorNumber, err := strconv.Atoi(c.Value)
	if err != nil {
		return 0, false
	}

	return visitorNumber, true
}

func writeVisitorCookie(w http.ResponseWriter, visitorNumber int) {
	http.SetCookie(w, &http.Cookie{
		Name:  CookieName,
		Value: strconv.Itoa(visitorNumber),
	})
}
