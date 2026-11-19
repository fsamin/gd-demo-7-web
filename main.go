// gd-demo-7-web is step 7 of the git-deploy demo: web exposures. The same
// small application as step 1, which also shows who is signed in when it is
// published behind the platform's authenticating proxy.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// version is bumped on each demo commit to make the rolling update visible.
const version = "v2"

func main() {
	log.Printf("gd-demo-7-web %s starting on :8080", version)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		hostname, _ := os.Hostname()
		printf(w, "Hello from gd-demo-7-web %s\n", version)
		printf(w, "\npod:     %s\n", hostname)
		printf(w, "served:  %s\n", time.Now().Format(time.RFC3339))

		// Set by the exposure's authenticating proxy. Behind it, the application
		// is reachable through the proxy only, so these come from it; on a
		// public exposure anyone can send them, which is fine for a display.
		if user := r.Header.Get("X-Forwarded-Preferred-Username"); user != "" {
			printf(w, "\nsigned in as: %s\n", user)
			// The platform puts the identity provider in the proxy's email field.
			if idp := r.Header.Get("X-Forwarded-Email"); idp != "" {
				printf(w, "through:      %s\n", idp)
			}
			if groups := r.Header.Get("X-Forwarded-Groups"); groups != "" {
				printf(w, "groups:       %s\n", groups)
			}
		} else {
			printf(w, "\nnot signed in: this exposure is public\n")
		}

		printf(w, "\nStep 7 of the git-deploy demo: web exposures.\n")
		printf(w, "Publish it on a hostname of your choice, then put a sign-in in front of it\n")
		printf(w, "with `git-deploy expose set --proxy --idp ...`: no code change needed here.\n")

		log.Printf("%s %s from %s (%s)", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start).Round(time.Millisecond))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// printf writes one line of the page, discarding the write error on purpose: a
// client hanging up mid-response is this server's normal weather, and there is
// nothing left to send it. Discarding it explicitly keeps the quality gate green
// if this repository ever declares one.
func printf(w io.Writer, format string, args ...any) {
	_, _ = fmt.Fprintf(w, format, args...)
}
