package main

import (
	"fmt"
	"net/http"
)

// Error handler
func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	switch status {
	case http.StatusNotFound:
		// Handle 404 Not Found
		fmt.Fprint(w, "Page not found (404)")
	case http.StatusMethodNotAllowed:
		// Handle 405 Method Not Allowed
		fmt.Fprint(w, "Method not allowed (405)")
	case http.StatusInternalServerError:
		// Handle 500 Internal Server Error
		fmt.Fprint(w, "Internal server error (500)")
	default:
		// Handle other statuses generically
		fmt.Fprintf(w, "Error %d", status)
	}
}

// setError sets the error message and redirects to the homepage.
func setError(w http.ResponseWriter, r *http.Request, errorMessage string) {
	res.Err = errorMessage // Set the error message
	res.Res = ""           // Clear previous values
	res.Res1 = ""

	// Force a reload of the page to show the error
	http.Redirect(w, r, "/", http.StatusFound)
}
