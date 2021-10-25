// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
)

func NewSearchHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := decodeSearchParams(r)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.Search(r.Context(), params)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeSearchResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewSearchByTagIDHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := decodeSearchByTagIDParams(r)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.SearchByTagID(r.Context(), params)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeSearchByTagIDResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewGetBookHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := decodeGetBookParams(r)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.GetBook(r.Context(), params)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeGetBookResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, writeErr := json.Marshal(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
	if writeErr == nil {
		w.Write(data)
	}
}