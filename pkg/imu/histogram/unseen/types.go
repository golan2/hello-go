package unseen

import "math"

type AllGainsHistogramResponseV2 struct {
	Data []AllGainsHistogramDataV2 `json:"data"`
}

type AllGainsHistogramDataV2 struct {
	DependentIndependentDataV2
	GainsHistogramDataV2
}

type DependentIndependentDataV2 struct {
	DependentName   string `json:"dependent_name"`
	IndependentName string `json:"independent_name"`
	DependentType   string `json:"dependent_type"`
	IndependentType string `json:"independent_type"`
}

type GainsHistogramDataV2 struct {
	UnseenAndSeen[*HistogramV2[float32]]
	ConfidenceRate *float32 `json:"confidence_rate,omitempty"`
	Opacity        *float32 `json:"opacity,omitempty"`
}

type UnseenAndSeen[T any] struct {
	Unseen  T `json:"unseen,omitempty"`
	Seen    T `json:"seen,omitempty"`
	Overall T `json:"overall,omitempty"`
}

type HistogramV2[T any] struct {
	MinValue    float32  `json:"min_value"`
	MaxValue    float32  `json:"max_value"`
	BinScale    float32  `json:"bin_scale"`
	MinBinSize  int      `json:"min_bin_size"`
	MaxBinSize  int      `json:"max_bin_size"`
	ValuesCount int      `json:"values_count"`
	Mean        float32  `json:"mean,omitempty"`
	Median      float32  `json:"median,omitempty"`
	StdDev      float32  `json:"std_dev,omitempty"`
	Bins        []Bin[T] `json:"bins"`
}

type Bin[T any] struct {
	Min        float32 `json:"min"`
	Max        float32 `json:"max"`
	Count      int     `json:"count"`
	DataPoints *[]T    `json:"data_points,omitempty"`
}

func (h *HistogramV2[T]) Clone() *HistogramV2[T] {
	if h == nil {
		return nil
	}

	clone := &HistogramV2[T]{
		MinValue:    h.MinValue,
		MaxValue:    h.MaxValue,
		BinScale:    h.BinScale,
		MinBinSize:  h.MinBinSize,
		MaxBinSize:  h.MaxBinSize,
		ValuesCount: h.ValuesCount,
		Mean:        h.Mean,
		Median:      h.Median,
		StdDev:      h.StdDev,
	}

	// Deep copy the Bins slice
	clone.Bins = make([]Bin[T], len(h.Bins))
	for i, bin := range h.Bins {
		clone.Bins[i] = Bin[T]{
			Min:   bin.Min,
			Max:   bin.Max,
			Count: bin.Count,
		}
		if bin.DataPoints != nil {
			dataPoints := make([]T, len(*bin.DataPoints))
			copy(dataPoints, *bin.DataPoints)
			clone.Bins[i].DataPoints = &dataPoints
		}
	}

	return clone
}

func (h *HistogramV2[T]) CloneShift() *HistogramV2[T] {
	if h == nil {
		return nil
	}

	singleBinWidth := (h.MaxValue - h.MinValue) * h.BinScale
	newMax := h.MaxValue + singleBinWidth

	clone := &HistogramV2[T]{
		MinValue:    h.MinValue,
		MaxValue:    newMax,
		BinScale:    h.BinScale,
		MinBinSize:  h.MinBinSize,
		MaxBinSize:  h.MaxBinSize,
		ValuesCount: h.ValuesCount,
		Mean:        h.Mean,
		Median:      h.Median,
		StdDev:      h.StdDev,
	}

	// Create a new Bins slice with one additional bin
	clone.Bins = make([]Bin[T], len(h.Bins)+1)

	// Create the first bin with zero count
	binWidth := (h.MaxValue - h.MinValue) / float32(len(h.Bins))
	clone.Bins[0] = Bin[T]{
		Min:   h.MinValue,
		Max:   h.MinValue + binWidth,
		Count: 0,
	}

	// Shift all other bins one position to the right
	for i := 0; i < len(h.Bins); i++ {
		clone.Bins[i+1] = Bin[T]{
			Min:   h.Bins[i].Min + binWidth,
			Max:   h.Bins[i].Max + binWidth,
			Count: h.Bins[i].Count,
		}
		if h.Bins[i].DataPoints != nil {
			dataPoints := make([]T, len(*h.Bins[i].DataPoints))
			copy(dataPoints, *h.Bins[i].DataPoints)
			clone.Bins[i+1].DataPoints = &dataPoints
		}
	}

	return clone
}

func NewEmptyHistogramV2(minValue, maxValue, binScale float32) *HistogramV2[float32] {
	if minValue >= maxValue || binScale <= 0 || binScale >= 1 {
		return nil // Invalid input parameters
	}

	totalRange := maxValue - minValue
	binWidth := totalRange * binScale
	numBins := int(math.Ceil(float64(totalRange / binWidth)))

	histogram := &HistogramV2[float32]{
		MinValue:    minValue,
		MaxValue:    maxValue,
		BinScale:    binScale,
		MinBinSize:  0, // Set to 0 as all bins are empty
		MaxBinSize:  0, // Set to 0 as all bins are empty
		ValuesCount: 0,
		Mean:        0,
		Median:      0,
		StdDev:      0,
		Bins:        make([]Bin[float32], numBins),
	}

	for i := 0; i < numBins; i++ {
		binMin := minValue + float32(i)*binWidth
		binMax := binMin + binWidth
		if i == numBins-1 {
			binMax = maxValue // Ensure the last bin reaches exactly to maxValue
		}

		histogram.Bins[i] = Bin[float32]{
			Min:   binMin,
			Max:   binMax,
			Count: 0,
		}
	}

	return histogram
}
