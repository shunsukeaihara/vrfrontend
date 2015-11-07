package vrfrontend

import (
	"container/list"
	"math"

	"github.com/mjibson/go-dsp/window"
)

type VAD interface {
	Update([]float64) (bool, error)
}

type PowerBasedVAD struct {
	powerThreshold float64
	widowSize      int
}

func NewPowerBasedVAD(threshold float64, windowSize int) *PowerBasedVAD {
	return &PowerBasedVAD{threshold, windowSize}
}

func (v *PowerBasedVAD) Update(signal []float64) (bool, error) {
	// signal must be normalized to [-1, 1]
	p := caluculatePower(signal)
	if p < v.powerThreshold {
		return true, nil
	} else {
		return false, nil
	}
}

type ZeroCrossVAD struct {
	powerThreshold     float64
	zerocrossThreshold int
	widowSize          int
}

func NewZeroCrossVAD(powerThreshold float64, zerocrossThreshold, windowSize int) *ZeroCrossVAD {
	return &ZeroCrossVAD{powerThreshold, zerocrossThreshold, windowSize}
}

func (v *ZeroCrossVAD) Update(signal []float64) (bool, error) {
	// signal must be normalized to [-1, 1]
	p := caluculatePower(signal)
	cross := caluculateZeroCross(signal)
	if p < v.powerThreshold && c > v.zerocrossThreshold {
		return true, nil
	} else {
		return false, nil
	}
}

type LTSDVAD struct {
	noiseSpectrum []float64
	order         int
	widowSize     int
	e0            float64
	e1            float64
	lambda0       float64
	lambda1       float64
	term          list.List
	window        []float64
}

func NewLTSDVAD(noiseFpectrum []float64, order, windowSize int, e0, e1, lambda0, lambda1 float64, winFunc func(int) []float64) {
	window := winFunc(windowSize)
	return &NewLTSDVAD{noiseFpectrum, order, windowSize, e0, e1, lambda0, lambda1, list.New(), window}
}

func (v *LTSDVAD) Update(signal []float64) (bool, error) {
	sig, error := ApplyWindow(signal, v.window)
	if error != nil {
		return false, error
	}

}

func caluculatePower(signal []float64) float64 {
	// signal must be normalized to [-1, 1]
	var sum float64
	for _, s := range signal {
		sum += math.Abs(s)
	}
	return 20.0 * math.Log10(sum/float64(len(signal)))
}

func caluculateZeroCross(signal []float64) int {
	// signal must be normalized to [-1, 1]
	var lastS float64
	lastS = 0.0
	count := 0
	for _, s := range signal {
		if (lastS * s) < 0 {
			// zero cross point
			count++
		}
		lastS = point
	}
	return count
}
