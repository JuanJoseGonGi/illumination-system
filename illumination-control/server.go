package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

var mux = http.NewServeMux()

var server = &http.Server{
	Addr:    ":3001",
	Handler: mux,
}

var errInvalidState = errors.New("invalid state")

const (
	respInvalidMethod     = "method not allowed"
	respInvalidInput      = "invalid input"
	respInvalidInputState = "invalid input state"
	respInternalError     = "internal error"
)

type state string

var (
	stateOn  state = "on"
	stateOff state = "off"

	validStates = map[state]bool{
		stateOn:  true,
		stateOff: true,
	}
)

type putInput struct {
	State state `json:"state"`
}

func decodePutInputJSON(body io.ReadCloser) (putInput, error) {
	input := putInput{}

	if err := json.NewDecoder(body).Decode(&input); err != nil {
		return putInput{}, fmt.Errorf("error decoding json: %w", err)
	}

	state, ok := validStates[input.State]
	if !state || !ok {
		return putInput{}, fmt.Errorf("%w: %s", errInvalidState, input.State)
	}

	return input, nil
}

func handleLightsModePut(w http.ResponseWriter, r *http.Request) {
	input, err := decodePutInputJSON(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(respInvalidInput))

		log.Error(respInvalidInput, "err", err)
		return
	}

	err = setLightsModeState(input.State)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(respInternalError))

		log.Error(respInternalError, "err", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("lights are on automatic state"))
}

func handleLightsModeGet(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("lights are on automatic state"))
}

func lightsModeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "PUT, GET, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodPut && r.Body != nil {
			handleLightsModePut(w, r)
			return
		}

		if r.Method == http.MethodGet {
			handleLightsModeGet(w, r)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(respInvalidMethod))

		log.Error(respInvalidMethod, "method", r.Method)
	}
}

func handleLightsPut(w http.ResponseWriter, r *http.Request) {
	input, err := decodePutInputJSON(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(respInvalidInput))

		log.Error(respInvalidInput, "err", err)
		return
	}

	err = setLightsState(input.State)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(respInternalError))

		log.Error(respInternalError, "err", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("lights are on"))
}

func handleLightsGet(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("lights are on"))
}

func lightsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodPut && r.Body != nil {
			handleLightsPut(w, r)
			return
		}

		if r.Method == http.MethodGet {
			handleLightsGet(w, r)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(respInvalidMethod))

		log.Error(respInvalidMethod, "method", r.Method)
	}
}

func handleCurtainsPut(w http.ResponseWriter, r *http.Request) {
	input := putInput{}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(respInvalidInput))

		log.Error(respInvalidInput, "err", err)

		return
	}

	err := setCurtainsState(input.State)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(respInternalError))

		log.Error(respInternalError, "err", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("curtains are open"))
}

func handleCurtainsGet(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("curtains are open"))
}

func curtainsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "PUT, GET, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodPut && r.Body != nil {
			handleCurtainsPut(w, r)
			return
		}

		if r.Method == http.MethodGet {
			handleCurtainsGet(w, r)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(respInvalidMethod))

		log.Error(respInvalidMethod, "method", r.Method)
	}
}

func handleBlindsPut(w http.ResponseWriter, r *http.Request) {
	input := putInput{}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(respInvalidInput))

		log.Error(respInvalidInput, "err", err)

		return
	}

	err := setBlindsState(input.State)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(respInternalError))

		log.Error(respInternalError, "err", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("blinds are open"))
}

func handleBlindsGet(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("blinds are open"))
}

func blindsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "PUT, GET, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodPut && r.Body != nil {
			handleBlindsPut(w, r)
			return
		}

		if r.Method == http.MethodGet {
			handleBlindsGet(w, r)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(respInvalidMethod))

		log.Error(respInvalidMethod, "method", r.Method)
	}
}

func initServer() {
	mux.Handle("/lights/mode", lightsModeHandler())
	mux.Handle("/lights", lightsHandler())
	mux.Handle("/curtains", curtainsHandler())
	mux.Handle("/blinds", blindsHandler())

	log.Info("starting server", "port", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
