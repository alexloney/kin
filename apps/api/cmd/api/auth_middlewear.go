package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	fbauth "firebase.google.com/go/v4/auth"
)

type ctxKey string

const (
	ctxFirebaseUID ctxKey = "firebase_uid"
)

func firebaseUIDFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(ctxFirebaseUID)
	uid, ok := v.(string)
	return uid, ok && uid != ""
}

func authMiddleware(firebaseAuth *fbauth.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authz := r.Header.Get("Authorization")
			if authz == "" || !strings.HasPrefix(authz, "Bearer ") {
				http.Error(w, "missing bearer token", http.StatusUnauthorized)
				return
			}

			idToken := strings.TrimSpace(strings.TrimPrefix(authz, "Bearer "))
			if idToken == "" {
				http.Error(w, "missing bearer token", http.StatusUnauthorized)
				return
			}

			ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
			defer cancel()

			tok, err := firebaseAuth.VerifyIDToken(ctx, idToken)
			if err != nil {
				http.Error(w, "invalid bearer token", http.StatusUnauthorized)
				return
			}

			// Attach UID to request context
			req := r.WithContext(context.WithValue(r.Context(), ctxFirebaseUID, tok.UID))
			next.ServeHTTP(w, req)
		})
	}
}