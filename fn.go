package function

import (
	"context"
	"image/jpeg"
	"log"

	"cloud.google.com/go/storage"
	"github.com/nfnt/resize"
)

type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

func OnStorageFinalize(ctx context.Context, e GCSEvent) error {
	log.Printf("Processing file: %s", e.Name)

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil
	}
	defer client.Close()

	file, err := client.Bucket(e.Bucket).Object(e.Name).NewReader(ctx)
	if err != nil {
		return nil
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	thum := resize.Resize(300, 0, img, resize.Lanczos3)

	writeFile := client.Bucket("images-sample-thum").Object("thum_" + e.Name).NewWriter(ctx)
	defer writeFile.Close()

	err = jpeg.Encode(writeFile, thum, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return nil
}

