package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"log"
	"math"
	"os"
	"strings"
)

const VERSION = "1.0.0"

func main() {
	// Parse command line arguments
	inputPath := flag.String("input", "", "Path to input PNG image")
	outputPath := flag.String("output", "animated.gif", "Path for output GIF")
	override := flag.Bool("y", false, "Override Y/N prompts with Y response")
	showVersion := flag.Bool("v", false, "Show version")
	flag.Parse()

	if *showVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if *inputPath == "" {
		log.Fatal("User Error: Please provide an input image path using -input flag")
	}

	existsErr := handleFileOverwrite(*outputPath, *override)
	if existsErr != nil {
		log.Fatal(existsErr)
	}

	// Read the source image
	sourceFile, err := os.Open(*inputPath)
	if err != nil {
		log.Fatal("Error opening source image:", err)
	}
	defer func(sourceFile *os.File) {
		err := sourceFile.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error closing source file: %v", err)
		}
	}(sourceFile)

	sourcePNG, err := png.Decode(sourceFile)
	if err != nil {
		log.Fatal("Error decoding PNG:", err)
	}

	// Verify image dimensions
	bounds := sourcePNG.Bounds()
	if bounds.Dx() != bounds.Dy() || bounds.Dy() < 144 {
		log.Fatal("Input image must be squared at least 144x144 pixels")
	}

	// Create a palette for the GIF that preserves colors
	colorSet := make(map[color.Color]bool)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			colorSet[sourcePNG.At(x, y)] = true
		}
	}

	// Convert unique colors to palette
	palette := make([]color.Color, 0, 256)
	palette = append(palette, color.Transparent) // Keep transparent as first color
	for c := range colorSet {
		if len(palette) < 256 {
			palette = append(palette, c)
		}
	}

	// Fill remaining palette slots if needed
	for len(palette) < 256 {
		palette = append(palette, color.Transparent)
	}

	// Define rotation angles for 24 frames
	angles := []float64{
		0,                 // 0 degrees
		math.Pi / 12,      // 15 degrees
		math.Pi / 6,       // 30 degrees
		math.Pi / 4,       // 45 degrees
		math.Pi / 3,       // 60 degrees
		5 * math.Pi / 12,  // 75 degrees
		math.Pi / 2,       // 90 degrees
		7 * math.Pi / 12,  // 105 degrees
		2 * math.Pi / 3,   // 120 degrees
		3 * math.Pi / 4,   // 135 degrees
		5 * math.Pi / 6,   // 150 degrees
		11 * math.Pi / 12, // 165 degrees
		math.Pi,           // 180 degrees
		13 * math.Pi / 12, // 195 degrees
		7 * math.Pi / 6,   // 210 degrees
		5 * math.Pi / 4,   // 225 degrees
		4 * math.Pi / 3,   // 240 degrees
		17 * math.Pi / 12, // 255 degrees
		3 * math.Pi / 2,   // 270 degrees
		19 * math.Pi / 12, // 285 degrees
		5 * math.Pi / 3,   // 300 degrees
		7 * math.Pi / 4,   // 315 degrees
		11 * math.Pi / 6,  // 330 degrees
		23 * math.Pi / 12, // 345 degrees
	}

	// Initialize animation parameters
	frames := make([]*image.Paletted, len(angles))
	delays := make([]int, len(angles))
	disposals := make([]byte, len(angles))

	// Set frame parameters
	for i := range angles {
		frames[i] = image.NewPaletted(bounds, palette)
		delays[i] = 17                      // 100 centiseconds = 1 second per frame
		disposals[i] = gif.DisposalPrevious // Clear previous frame before rendering next
	}

	// Generate the frames
	for i, angle := range angles {
		transformFrame(sourcePNG, frames[i], angle)
	}

	// Create output file
	outFile, err := os.Create(*outputPath)
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error closing output file: %v", err)
		}
	}(outFile)

	// Encode the GIF with disposal methods
	err = gif.EncodeAll(outFile, &gif.GIF{
		Image:    frames,
		Delay:    delays,
		Disposal: disposals,
	})
	if err != nil {
		log.Fatal("Error encoding GIF:", err)
	}

	fmt.Printf("Animation saved to %s\n", *outputPath)
}

func transformFrame(src image.Image, dst *image.Paletted, angle float64) {
	bounds := src.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Clear the destination frame
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dst.Set(x, y, color.Transparent)
		}
	}

	centerX := width / 2

	// Use reverse mapping to prevent gaps
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Calculate reverse projection
			relX := float64(x - centerX)

			// Apply perspective and rotation in reverse
			scale := math.Cos(angle)
			if math.Abs(scale) < 0.01 {
				scale = 0.01 // Prevent division by zero
			}

			// Calculate source x position
			srcX := int(relX/scale) + centerX

			// Perform anti-aliasing by sampling multiple points
			if srcX >= 0 && srcX < width {
				c := src.At(srcX, y)

				// Only process non-transparent pixels
				if _, _, _, a := c.RGBA(); a > 0 {
					// Edge darkening when viewing the coin's side
					if math.Abs(math.Sin(angle)) > 0.7 && math.Abs(relX) > float64(width)*0.48 {
						r, g, b, a := c.RGBA()
						c = color.RGBA{
							R: uint8(float64(r>>8) * 0.7),
							G: uint8(float64(g>>8) * 0.7),
							B: uint8(float64(b>>8) * 0.7),
							A: uint8(a >> 8),
						}
					}

					// Apply additional interpolation for smoother transitions
					if srcX > 0 && srcX < width-1 {
						// Get neighboring pixels for interpolation
						c1 := src.At(srcX-1, y)
						c2 := src.At(srcX+1, y)

						// Only interpolate if neighbors are not transparent
						if _, _, _, a1 := c1.RGBA(); a1 > 0 {
							if _, _, _, a2 := c2.RGBA(); a2 > 0 {
								r1, g1, b1, _ := c1.RGBA()
								r2, g2, b2, _ := c2.RGBA()
								r, g, b, a := c.RGBA()

								// Blend colors for smoother transition
								blend := math.Abs(math.Sin(angle)) * 0.5
								c = color.RGBA{
									R: uint8(float64(r>>8)*(1-blend) + float64((r1+r2)>>9)*blend),
									G: uint8(float64(g>>8)*(1-blend) + float64((g1+g2)>>9)*blend),
									B: uint8(float64(b>>8)*(1-blend) + float64((b1+b2)>>9)*blend),
									A: uint8(a >> 8),
								}
							}
						}
					}

					dst.Set(x, y, c)
				}
			}
		}
	}
}

// handleFileOverwrite checks if the output file exists and handles overwrite protection.
// It returns an error if the file shouldn't be overwritten or if there's a problem.
func handleFileOverwrite(outputPath string, override bool) error {
	fileInfo, err := os.Stat(outputPath)

	// If file doesn't exist, we can proceed
	if os.IsNotExist(err) {
		return nil
	}

	// Handle other stat errors
	if err != nil {
		return fmt.Errorf("error checking output file: %v", err)
	}

	// Check if it's a directory
	if fileInfo.IsDir() {
		return fmt.Errorf("output path points to a directory: %s", outputPath)
	}

	// If override flag is set, proceed with overwrite
	if override {
		return nil
	}

	// If file exists and has content, prompt for confirmation
	if fileInfo.Size() > 0 {
		// Create a buffered reader for user input
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("File %s already exists. Overwrite? [y/N]: ", outputPath)
		response, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error reading user input: %v", err)
		}

		// Clean up the response string
		response = strings.ToLower(strings.TrimSpace(response))

		// Only proceed if user explicitly confirms with 'y' or 'yes'
		if response != "y" && response != "yes" {
			return fmt.Errorf("operation cancelled by user")
		}
	}

	return nil
}
