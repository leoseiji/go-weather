package dto

type CEPWeatherResponse struct {
	TemperatureInCelcius    float64 `json:"temp_C"`
	TemperatureInFahrenheit float64 `json:"temp_F"`
	TemperatureInKelvin     float64 `json:"temp_K"`
}

func NewCEPWeatherResponse(weather *Weather) *CEPWeatherResponse {
	return &CEPWeatherResponse{
		TemperatureInCelcius:    weather.Current.TempC,
		TemperatureInFahrenheit: weather.Current.TempF,
		TemperatureInKelvin:     weather.Current.TempC + 273.15,
	}
}
