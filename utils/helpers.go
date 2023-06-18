package utils

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os/exec"
)

func ReadJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func GenerateShortURL() string {
	num := rand.Int()
	hash := EncodeBase62(num)

	return hash
}

func OpenWebPage(url string) {
	cmd := exec.Command("open", url)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
