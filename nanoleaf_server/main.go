package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	nanoleafIP        = "192.168.1.5"
	nanoleafPort      = "16021"
	nanoleafAuthToken = "WjA9hXKzkagUqzy95ETF3RsBZYwTTTNN"
	nanoleafUrl       = "http://" + nanoleafIP + ":" + nanoleafPort
)

// Used to control brightness
type Brightness struct {
	Value    int `json:"value"`
	Duration int `json:"duration"`
}
type BrightnessState struct {
	Brightness Brightness `json:"brightness"`
}

// Used to control the color
type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}
type CT struct {
	CT int `json:"ct"`
}
type Hue struct {
	Hue int `json:"hue"`
}
type Sat struct {
	Sat int `json:"sat"`
}
type ColorComponent struct {
	Value int `json:"value"`
}
type CtValueState struct {
	Ct ColorComponent `json:"ct"`
}
type HueValueState struct {
	Hue ColorComponent `json:"hue"`
}
type SatValueState struct {
	Sat ColorComponent `json:"sat"`
}
type ColorState struct {
	Hue ColorComponent `json:"hue"`
	Sat ColorComponent `json:"sat"`
}

func brightnessHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		url := fmt.Sprintf("%s/api/v1/%s/state/brightness", nanoleafUrl, nanoleafAuthToken)
		fmt.Println(url)
		resp, err := http.Get(url)
		errCheck(err, w, "Failed to communicate with the light")
		defer resp.Body.Close()
		nanoleafErrCheck(resp, w)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		detectRespError(w, resp, "Failed to get brightness")
	case http.MethodPut:

		var reqBody BrightnessState
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		body, err := json.Marshal(reqBody)
		errCheck(err, w, "Unable to serialize brightness")

		url := fmt.Sprintf("%s/api/v1/%s/state/brightness", nanoleafUrl, nanoleafAuthToken)
		reqPut, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
		errCheck(err, w, "Failed to communicate with the light")
		reqPut.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(reqPut)
		errCheck(err, w, "Failed to update light brightness")
		defer resp.Body.Close()
		nanoleafErrCheck(resp, w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}
}

func lightsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		var reqBody struct {
			On bool `json:"on"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		payload := map[string]interface{}{
			"on": map[string]bool{
				"value": reqBody.On,
			},
		}
		body, err := json.Marshal(payload)
		errCheck(err, w, "Unable to serialize light state")

		url := fmt.Sprintf("%s/api/v1/%s/state", nanoleafUrl, nanoleafAuthToken)
		reqPut, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
		errCheck(err, w, "Failed to communicate with the light")
		reqPut.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(reqPut)
		errCheck(err, w, "Failed to update light state")
		defer resp.Body.Close()
		nanoleafErrCheck(resp, w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	case http.MethodGet:
		url := fmt.Sprintf("%s/api/v1/%s/state/on", nanoleafUrl, nanoleafAuthToken)
		resp, err := http.Get(url)

		errCheck(err, w, "Failed to communicate with the light")
		defer resp.Body.Close()
		nanoleafErrCheck(resp, w)

		w.Header().Set("Content-Type", "application/json")
		detectRespError(w, resp, "Failed to get light state")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ctHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		var ct CT
		if err := json.NewDecoder(r.Body).Decode(&ct); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		payload := CtValueState{
			Ct: ColorComponent{Value: ct.CT},
		}
		body, err := json.Marshal(payload)
		errCheck(err, w, "Unable to serialize color")
		url := fmt.Sprintf("%s/api/v1/%s/state", nanoleafUrl, nanoleafAuthToken)

		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
		errCheck(err, w, "Failed to create request")
		req.Header.Add("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		errCheck(err, w, "Failed to send request to Nanoleaf")
		defer resp.Body.Close()
		nanoleafErrCheck(resp, w)
	case http.MethodGet:
		url := fmt.Sprintf("%s/api/v1/%s/state/ct", nanoleafUrl, nanoleafAuthToken)
		resp, err := http.Get(url)
		errCheck(err, w, "Failed to communicate with the light")
		defer resp.Body.Close()
		nanoleafErrCheck(resp, w)
		w.Header().Set("Content-Type", "application/json")
		detectRespError(w, resp, "Failed to get light state")
	}
}

func hueHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		var hue Hue
		if err := json.NewDecoder(r.Body).Decode(&hue); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		payload := HueValueState{
			Hue: ColorComponent{Value: hue.Hue},
		}
		body, err := json.Marshal(payload)
		errCheck(err, w, "Unable to serialize color")
		url := fmt.Sprintf("%s/api/v1/%s/state", nanoleafUrl, nanoleafAuthToken)

		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
		errCheck(err, w, "Failed to create request")
		req.Header.Add("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		errCheck(err, w, "Failed to send request to Nanoleaf")
		defer resp.Body.Close()
		nanoleafErrCheck(resp, w)
	case http.MethodGet:
		url := fmt.Sprintf("%s/api/v1/%s/state/hue", nanoleafUrl, nanoleafAuthToken)
		resp, err := http.Get(url)
		errCheck(err, w, "Failed to communicate with the light")
		defer resp.Body.Close()
		nanoleafErrCheck(resp, w)
		w.Header().Set("Content-Type", "application/json")
		detectRespError(w, resp, "Failed to get light state")
	}
}

func satHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		var sat Sat
		if err := json.NewDecoder(r.Body).Decode(&sat); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		payload := SatValueState{
			Sat: ColorComponent{Value: sat.Sat},
		}
		body, err := json.Marshal(payload)
		errCheck(err, w, "Unable to serialize color")
		url := fmt.Sprintf("%s/api/v1/%s/state", nanoleafUrl, nanoleafAuthToken)

		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
		errCheck(err, w, "Failed to create request")
		req.Header.Add("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		errCheck(err, w, "Failed to send request to Nanoleaf")
		defer resp.Body.Close()
		nanoleafErrCheck(resp, w)
	case http.MethodGet:
		url := fmt.Sprintf("%s/api/v1/%s/state/sat", nanoleafUrl, nanoleafAuthToken)
		resp, err := http.Get(url)
		errCheck(err, w, "Failed to communicate with the light")
		defer resp.Body.Close()
		nanoleafErrCheck(resp, w)
		w.Header().Set("Content-Type", "application/json")
		detectRespError(w, resp, "Failed to get light state")
	}
}

func colorHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		var rgb RGB
		if err := json.NewDecoder(r.Body).Decode(&rgb); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		h, s, _ := rbhToHsv(float64(rgb.R), float64(rgb.G), float64(rgb.B))
		sPercent := int(s * 100)
		hueValue := int(h)

		payload := ColorState{
			Hue: ColorComponent{Value: hueValue},
			Sat: ColorComponent{Value: sPercent},
		}
		body, err := json.Marshal(payload)
		errCheck(err, w, "Unable to serialize color")
		url := fmt.Sprintf("%s/api/v1/%s/state", nanoleafUrl, nanoleafAuthToken)

		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
		errCheck(err, w, "Failed to create request")
		req.Header.Add("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		errCheck(err, w, "Failed to send request to Nanoleaf")
		defer resp.Body.Close()
		nanoleafErrCheck(resp, w)
	case http.MethodGet:
		url := fmt.Sprintf("%s/api/v1/%s/state", nanoleafUrl, nanoleafAuthToken)
		resp, err := http.Get(url)
		errCheck(err, w, "Failed to get state from Nanoleaf")
		defer resp.Body.Close()
		nanoleafErrCheck(resp, w)
		// Decode the response. We're only interested in hue and sat here.
		var state ColorState
		if err := json.NewDecoder(resp.Body).Decode(&state); err != nil {
			http.Error(w, "Failed to parse state", http.StatusInternalServerError)
			return
		}
		// Convert hue and saturation back to RGB.
		h := float64(state.Hue.Value)
		s := float64(state.Sat.Value) / 100.0
		// Assume full brightness (value=1). Adjust if you have a brightness value.
		rVal, gVal, bVal := hsvToRgb(h, s, 1.0)
		rgb := RGB{R: rVal, G: gVal, B: bVal}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rgb)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func rbhToHsv(r, g, b float64) (h, s, v float64) {
	r /= 255.0
	g /= 255.0
	b /= 255.0

	rMax := math.Max(r, math.Max(g, b))
	rMin := math.Min(r, math.Min(g, b))
	v = rMax
	delta := rMax - rMin
	if rMax != 0 {
		s = delta / rMax
	} else {
		s = 0
		h = 0
		return
	}

	switch {
	case delta == 0:
		h = 0
	case rMax == r:
		h = 60 * math.Mod((g-b)/delta, 6)
	case rMax == g:
		h = 60 * (((b - r) / delta) + 2)
	case rMax == b:
		h = 60 * (((r - g) / delta) + 4)
	}

	if h < 0 {
		h += 360
	}
	return
}

func hsvToRgb(h, s, v float64) (r, g, b int) {
	c := v * s
	x := c * (1 - math.Abs(math.Mod(h/60.0, 2)-1))
	m := v - c

	var rF, gF, bF float64
	switch {
	case h >= 0 && h < 60:
		rF, gF, bF = c, x, 0
	case h >= 60 && h < 120:
		rF, gF, bF = x, c, 0
	case h >= 120 && h < 180:
		rF, gF, bF = 0, c, x
	case h >= 180 && h < 240:
		rF, gF, bF = 0, x, c
	case h >= 240 && h < 300:
		rF, gF, bF = x, 0, c
	case h >= 300 && h < 360:
		rF, gF, bF = c, 0, x
	default:
		rF, gF, bF = 0, 0, 0
	}

	r = int((rF + m) * 255)
	g = int((gF + m) * 255)
	b = int((bF + m) * 255)
	return
}

func errCheck(err error, w http.ResponseWriter, message string) {
	if err != nil {
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
}

func nanoleafErrCheck(resp *http.Response, w http.ResponseWriter) {
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Nanoleaf API error: %d", resp.StatusCode), resp.StatusCode)
		return
	}
}

func detectRespError(w http.ResponseWriter, resp *http.Response, message string) {
	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/lights", lightsHandler)
	http.HandleFunc("/lights/brightness", brightnessHandler)
	http.HandleFunc("/lights/color", colorHandler)
	http.HandleFunc("/lights/ct", ctHandler)
	http.HandleFunc("/lights/hue", hueHandler)
	http.HandleFunc("/lights/sat", satHandler)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
