package routes

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	"github.com/go-chi/chi"
	c "github.com/luke-jj/go-weather-api/internal/config"
)

func Weather() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getWeather)

	return r
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	config, ok := r.Context().Value("config").(*c.Config)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	city := r.URL.Query().Get("city")
	if city == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{ "message": "Query string param 'city' is required."}`))
		return
	}
	if !regexp.MustCompile(`^[0-9A-Za-z .,]{1,42}$`).MatchString(city) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{ "message": "Illegal city name."}`))
		return
	}
	params := url.Values{}
	params.Add("q", city)
	params.Add("units", "metric")
	params.Add("APPID", config.WEATHER_KEY)
	uri := &url.URL{
		Scheme:   "https",
		Host:     config.WEATHER_URI,
		Path:     "/data/2.5/forecast",
		RawQuery: params.Encode(),
	}
	response, err := http.Get(uri.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "Problem fetching weather from external api."}`))
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	w.Write([]byte(data))
}
