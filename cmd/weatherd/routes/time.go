package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	c "github.com/luke-jj/go-weather-api/internal/config"
)

type Time struct {
	ID                    string `json:"$id"`
	CurrentDateTime       string `json:"currentDateTime"`
	UTCOffset             string `json:"utcOffset"`
	IsDayLightSavingsTime bool   `json:"isDayLightSavingsTime"`
	DayOfTheWeek          string `json:"dayOfTheWeek"`
	TimeZoneName          string `json:"UTC""`
	CurrentFileTime       int64  `json:"currentFileTime"`
	OrdinalDate           string `json:"ordinalDate"`
}

func Times() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getTime)

	return r
}

func getTime(w http.ResponseWriter, r *http.Request) {
	var t Time
	con, ok := r.Context().Value("config").(*c.Config)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	uri := &url.URL{
		Scheme: "http",
		Host:   con.TIME_URI,
		Path:   "/api/json/utc/now",
	}
	response, err := http.Get(uri.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "Problem fetching time from external api."}`))
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	err = json.Unmarshal(data, &t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	// convert file time to unix time
	datestring := strconv.FormatInt(t.CurrentFileTime, 10)
	fileTime, _ := strconv.Atoi(datestring[:len(datestring)-7])
	epoch := time.Date(1601, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	unixSeconds := int(epoch) + fileTime
	w.Write([]byte(`{ "currentUnixTime": ` + strconv.Itoa(unixSeconds) + ` }`))
}
