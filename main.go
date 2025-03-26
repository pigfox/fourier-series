package main

import (
	"fmt"
	"math"
)

// Point represents a 2D point
type Point struct {
	X, Y float64
}

// FourierSeries holds the series parameters
type FourierSeries struct {
	Period       float64
	NumHarmonics int
}

// CalculateCoefficient computes the nth coefficient for a square wave
func (fs *FourierSeries) CalculateCoefficient(n int) float64 {
	if n == 0 {
		return 0 // DC component is 0 for square wave
	}
	// For square wave, only odd harmonics exist
	if n%2 == 0 {
		return 0
	}
	return 4.0 / (float64(n) * math.Pi)
}

// Evaluate computes the series value at point t
func (fs *FourierSeries) Evaluate(t float64) float64 {
	sum := 0.0
	for n := 1; n <= fs.NumHarmonics; n++ {
		coeff := fs.CalculateCoefficient(n)
		sum += coeff * math.Sin(2.0*math.Pi*float64(n)*t/fs.Period)
	}
	return sum
}

// GeneratePoints creates a slice of points for plotting
func GeneratePoints(fs *FourierSeries, samples int) []Point {
	points := make([]Point, samples)
	dt := fs.Period / float64(samples)

	for i := 0; i < samples; i++ {
		t := float64(i) * dt
		points[i] = Point{
			X: t,
			Y: fs.Evaluate(t),
		}
	}
	return points
}

// Simple console visualization
func Visualize(points []Point) {
	const height = 20
	const width = 70

	// Find min and max Y values
	minY, maxY := points[0].Y, points[0].Y
	for _, p := range points {
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	// Create display grid
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	// Plot points
	for i, p := range points {
		x := int(float64(i) * float64(width) / float64(len(points)))
		normalizedY := (p.Y - minY) / (maxY - minY)
		y := height - 1 - int(normalizedY*float64(height-1))
		if y >= 0 && y < height && x >= 0 && x < width {
			grid[y][x] = '*'
		}
	}

	// Draw zero line
	zeroY := height - 1 - int((0-minY)/(maxY-minY)*float64(height-1))
	if zeroY >= 0 && zeroY < height {
		for x := 0; x < width; x++ {
			if grid[zeroY][x] == ' ' {
				grid[zeroY][x] = '-'
			}
		}
	}

	// Print the grid
	for i := range grid {
		fmt.Println(string(grid[i]))
	}
}

func main() {
	// Create a Fourier Series with 7 harmonics
	fs := FourierSeries{
		Period:       2.0, // Period of 2 units
		NumHarmonics: 7,   // Using 7 harmonics for approximation
	}

	// Configuration values
	samples := 200
	numExamplePoints := 5 // Changed from 8 to 5 to stay within bounds (200/40=5)

	// Generate points
	points := GeneratePoints(&fs, samples)

	// Print example values
	fmt.Println("Example values:")
	for i := 0; i < numExamplePoints; i++ {
		fmt.Printf("t=%.2f, f(t)=%.4f\n", points[i*40].X, points[i*40].Y)
	}

	// Visualize the result
	fmt.Println("\nFourier Series Approximation of Square Wave:")
	Visualize(points)

	fmt.Printf("\nUsing %d harmonics\n", fs.NumHarmonics)
	fmt.Println("This approximates a square wave using sine terms")
}
