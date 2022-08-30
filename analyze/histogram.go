package analyze

import "fmt"

type Hist struct {
	counts      map[int]int
	min         int
	max         int
	sum         int
	obsCount    int
	bucketCount int
	BarChar     string
}

func NewHist(bucketCount int) *Hist {
	return &Hist{
		counts:      make(map[int]int),
		bucketCount: bucketCount,
		BarChar:     "◼",
	}
}

func (h *Hist) Add(value int) {
	h.sum += value

	if h.obsCount == 0 {
		h.counts[value] = 1

		h.min = value
		h.max = value

		h.obsCount = 1

		return
	}

	if _, ok := h.counts[value]; !ok {
		h.counts[value] = 1
	} else {
		h.counts[value]++
	}

	if value < h.min {
		h.min = value
	}
	if value > h.max {
		h.max = value
	}

	h.obsCount++
}

func (h *Hist) Print(width int) {
	bucketSize := (h.max - h.min) / h.bucketCount

	maxBucket := h.bucketCount * bucketSize
	maxBucketLen := len(fmt.Sprint(maxBucket))
	labelFmt := fmt.Sprintf("%%0%dd", maxBucketLen)

	labels := make([]string, h.bucketCount)
	for barIndex := 0; barIndex < h.bucketCount; barIndex++ {
		labels[barIndex] = fmt.Sprintf(labelFmt, barIndex*bucketSize)
	}

	bars := make([]int, h.bucketCount)
	for countIndex := h.min; countIndex <= h.max; countIndex++ {
		if count, ok := h.counts[countIndex]; ok {
			barIndex := countIndex / bucketSize
			if barIndex == h.bucketCount {
				barIndex--
			}

			bars[barIndex] += count
		}
	}

	maxBar := 0
	for _, bar := range bars {
		if bar > maxBar {
			maxBar = bar
		}
	}

	pcts := make([]float64, h.bucketCount)
	for index, bar := range bars {
		pcts[index] = float64(bars[index]) / float64(h.obsCount)

		adjBar := int(float64(bars[index]) / float64(maxBar) * float64(width))
		if bar > 0 && adjBar == 0 {
			adjBar = 1
		}
		bars[index] = adjBar
	}

	for index, bar := range bars {
		fmt.Printf("%s | ", labels[index])
		for j := 0; j < bar; j++ {
			fmt.Print(h.BarChar)
		}

		if bar > 0 {
			fmt.Printf(" %0.2f%%", pcts[index]*100)
		}

		fmt.Println()
	}
}

func (h *Hist) Mean() float64 {
	return float64(h.sum) / float64(h.obsCount)
}

func (h *Hist) Min() int {
	return h.min
}

func (h *Hist) Max() int {
	return h.max
}

func (h *Hist) Obs() int {
	return h.obsCount
}
