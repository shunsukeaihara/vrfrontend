package vrfrontend

import (
	"errors"

	"github.com/mjibson/go-dsp/fft"
	"github.com/mjibson/go-dsp/window"
)

func ApplyWindow(signal []float64, window []float64) ([]float64, error) {
	ret := make([]float64, len(signal))
	if len(signal) != len(window) {
		return ret, errors.New("signal size is not equal to window")
	}

	for i, w := range window {
		ret[i] = signal[i] * w
	}
	return ret, nil
}
