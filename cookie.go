package main

import "net/http"

const CookieName = "has_visited"

func visitedCookieExists(r *http.Request) bool {
	_, err := r.Cookie(CookieName)
	if err != nil {
		return false
	}
	return true
}

func writeVisitedCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:  CookieName,
		Value: "1",
	})
}
