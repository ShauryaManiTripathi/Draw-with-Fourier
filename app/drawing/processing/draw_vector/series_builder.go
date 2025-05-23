package draw_vector

import (
	"math"
	"math/cmplx"

	"api/app/drawing/types"
)

func BuildSeries(originalPoints []types.OriginalPoint, maxVectors int) []types.DrawVector {
	n := 0
	vectors := []types.DrawVector{}
	vectorBuilder := VectorBuilder{}

	for (len(vectors) < maxVectors) && !vectorsAproximateOriginal(vectors, originalPoints) {
		vectors = append(vectors, vectorBuilder.Build(n, originalPoints))

		n = getNextN(n)
		print("n: ", n, " len: ", len(vectors), " average distance: ", getAverageDistance(originalPoints, vectors), "\n")
	}

	return vectors
}

func getNextN(n int) int {
	if n > 0 {
		return -1 * n
	}

	return (-1 * n) + 1
}

func vectorsAproximateOriginal(vectors []types.DrawVector, originalPoints []types.OriginalPoint) bool {
	if len(vectors) == 0 {
		return false
	}

	averageDistance := getAverageDistance(originalPoints, vectors)

	return averageDistance < 1
}

func getAverageDistance(originalPoints []types.OriginalPoint, vectors []types.DrawVector) float64 {
	distance := 0.00

	for i := 0; i < len(originalPoints); i++ {
		original := originalPoints[i]
		estimate := calculateOutput(originalPoints[i].Time, vectors)
		distance += math.Sqrt(math.Pow(real(estimate)-float64(original.X), 2) + math.Pow(imag(estimate)-float64(original.Y), 2))
	}

	return distance / float64(len(originalPoints))
}

func calculateOutput(time float64, vectors []types.DrawVector) complex128 {
	sum := complex(0, 0)

	for _, vector := range vectors {
		c := complex(vector.Real, vector.Imaginary)
		power := complex(0.00, float64(vector.N)*2.00*math.Pi*time)
		sum += c * cmplx.Exp(power)
	}

	return sum
}
