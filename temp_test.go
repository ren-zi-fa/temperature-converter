package temperature

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	result := CelsiusToFahrenheit(0)
	expected := 32.0
	if result != expected {
		t.Errorf("CelsiusToFahrenheit(0) = %v; want %v", result, expected)
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	result := FahrenheitToCelsius(32)
	expected := 0.0
	if result != expected {
		t.Errorf("FahrenheitToCelsius(32) = %v; want %v", result, expected)
	}
}
