//go:build !examples
// +build !examples

// Unit tests can be run with `go test -v` and require access to the API. Many of
// these tests are actually integration tests that call the weather.gov API and
// parse responses accordingly to confirm expected responses are returned.
//
// Thus, in the future if weather.gov changes the endpoints or responses, these
// tests should alert users of this wrapper SDK accordingly.
package nws_test

import (
	"testing"

	"github.com/jacaudi/nws"
)

func TestBlank(t *testing.T) {
	point, err := nws.Points("", "")
	if point == nil && err != nil {
		return
	}
	t.Error("nws.Points() should return a 404 error for a blank lat, lon.")
}

func TestBlankLat(t *testing.T) {
	point, err := nws.Points("", "-147.7390417")
	if point == nil && err != nil {
		return
	}
	t.Error("nws.Points() should return a 404 error for a blank lat.")
}

func TestBlankLon(t *testing.T) {
	point, err := nws.Points("64.828421", "")
	if point == nil && err != nil {
		return
	}
	t.Error("nws.Points() should return a 404 error for a blank lon.")
}

func TestZero(t *testing.T) {
	point, err := nws.Points("0", "0")
	if point == nil && err != nil {
		return
	}
	t.Error("nws.Points() should return a 404 error for a zero lat, lon.")
}

func TestInternational(t *testing.T) {
	point, err := nws.Points("48.85660", "2.3522") // Paris, France
	if point == nil && err != nil {
		return
	}
	t.Error("nws.Points() should return a 404 error for lat, lon outside the U.S. territories.")
}

func TestAlaska(t *testing.T) {
	point, err := nws.Points("64.828421", "-147.7390417")
	if point != nil && err == nil {
		return
	}
	t.Error("nws.Points() should return valid points for parts of Alaska.")
}

func TestMetricUnits(t *testing.T) {
	nws.SetUnits("si")
	forecast, err := nws.Forecast("41.837", "-87.685")
	if err != nil || forecast == nil {
		t.Error("nws.Forecast() should return valid data for Chicago.")
		return
	}
	if forecast.Units != "si" {
		t.Error("nws.Forecast() should return valid data for Chicago in metric.")
	}
}

func TestUSUnits(t *testing.T) {
	nws.SetUnits("us")
	forecast, err := nws.Forecast("41.837", "-87.685")
	if err != nil {
		t.Error("nws.Forecast() should return valid data for Chicago.")
	}
	if forecast.Units != "us" {
		t.Error("nws.Forecast() should return valid data for Chicago in standard units.")
	}
}

func TestChicagoOffice(t *testing.T) {
	office, err := nws.Office("LOT")
	if office != nil && err == nil {
		if office.Name == "Chicago, IL" {
			return
		}
	}
	t.Error("nws.Office(\"LOT\") should return valid office information.")
}

func TestChicagoHourly(t *testing.T) {
	hourly, err := nws.HourlyForecast("41.837", "-87.685")
	if err != nil {
		t.Error("nws.HourlyForecast() should return valid data for Chicago.")
	}
	if len(hourly.Periods) == 0 {
		t.Error("expected at least one period")
	}
}
