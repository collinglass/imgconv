package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	toType *string

	converters = map[string]func(io.Writer, io.Reader) error{
		"png":  convertToPNG,
		"jpeg": convertToJPEG,
		"gif":  convertToGIF,
	}
)

type Elf struct{}

func (e Elf) ConvertImage(imagePath string, orderErr chan error, wg *sync.WaitGroup) {
	defer func() {
		time.Sleep(400 * time.Millisecond)
		wg.Done()
	}()

	fi, err := os.Open(imagePath)
	if err != nil {
		orderErr <- err
		return
	}

	z := strings.SplitN(imagePath, ".", 2)
	fo, err := os.Create(fmt.Sprintf("%v.%v", z[0], *toType))
	if err != nil {
		orderErr <- err
		return
	}

	if err := converters[*toType](fo, fi); err != nil {
		orderErr <- err
		return
	}

}

func main() {
	// flag definition and parsing
	toType = flag.String("to", "png", "To defines what type you'd like your images to be.")
	flag.Parse()

	// Elven feedback
	fmt.Println("The elves are processing your request...")
	time.Sleep(400 * time.Millisecond)

	// Processing

	// check parse errors
	if _, ok := converters[*toType]; !ok {
		fmt.Printf("Papa Elf: Sorry we don't convert to that type of image format. \nType entered: %v\n", *toType)
		return
	}

	// Elven feedback
	fmt.Println("Papa Elf is distributing your images...")
	time.Sleep(400 * time.Millisecond)

	elves := make([]Elf, len(flag.Args()))

	orderErr := make(chan error)

	var wg sync.WaitGroup
	for i, arg := range flag.Args() {
		wg.Add(1)
		filename := arg
		go elves[i].ConvertImage(filename, orderErr, &wg)
	}

	errorList := []error{}
	go func() {
		for err := range orderErr {
			errorList = append(errorList, err)
		}
	}()
	wg.Wait()

	if len(errorList) == 0 {
		// Elven feedback
		fmt.Printf("Papa Elf: Here's your order for %v %v images...\n", len(flag.Args()), *toType)
	} else {
		// Elven feedback
		fmt.Printf("Papa Elf: There were some hiccups with your order, we've delivered the %v images that we were able to complete...\n", *toType)
		fmt.Println("Here's the list of problems we encountered:")
		for _, err := range errorList {
			fmt.Println("\t- ", err)
		}
	}
}

// convertToGIF converts from any recognized format to GIF.
func convertToGIF(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return gif.Encode(w, img, &gif.Options{NumColors: 256})
}

// convertToJPEG converts from any recognized format to JPEG.
func convertToJPEG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, &jpeg.Options{Quality: 100})
}

// convertToPNG converts from any recognized format to PNG.
func convertToPNG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return png.Encode(w, img)
}
