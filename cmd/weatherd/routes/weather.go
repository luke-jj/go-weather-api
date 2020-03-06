package routes

import (
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	c "github.com/luke-jj/go-weather-api/internal/config"
)

func Weather() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getWeather)

	return r
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	con, ok := r.Context().Value("config").(*c.Config)
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

	// TODO: reg ex test city parameter
	// if (!/^[0-9A-Za-z .,]{1,42}$/.test(req.query['city'])) {
	// return res.status(400).send('Illegal city name.');
	// }

	// TODO: escape the city query parameter

	weatherUri := "https://" + con.WEATHER_URI + "/data/2.5/forecast" + "?city=" + city
	response, err := http.Get(weatherUri)

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
