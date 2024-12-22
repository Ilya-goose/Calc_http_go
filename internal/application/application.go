package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Ilya-goose/Calc_http_go/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

var ErrInternalServer = errors.New("internal server error")

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	var request Request
	var buffer bytes.Buffer

	encoder := json.NewEncoder(&buffer)

	err := decoder.Decode(&request)
	if err != nil {
		responce := Response{
			Error: ErrInternalServer.Error(),
		}
		encoder.Encode(responce)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, buffer.String())
		return
	}

	res_of_cal, err := calculation.Calc(request.Expression)
	if err != nil {
		responce := Response{
			Error: err.Error(),
		}
		encoder.Encode(responce)
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintln(w, buffer.String())
	} else {
		response := Response{
			Result: res_of_cal,
		}
		encoder.Encode(response)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, buffer.String())
	}
}

func (app *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+app.config.Addr, nil)
}
