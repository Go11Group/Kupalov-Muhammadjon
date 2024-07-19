package main

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	imagePath := "image.jpg"
	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Example: Write imageBytes to a binary file
	err = ioutil.WriteFile("output.jpg", imageBytes, 0644)
	if err != nil {
		log.Fatal("Error writing file:", err)
	}

	saveImage(imageBytes, "output.jpg") // Corrected file extension to .png
}

func saveImage(profileImage []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(bytes.NewReader(profileImage))
	if err != nil {
		return err
	}

	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}
