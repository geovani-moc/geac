package action

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

//LoadImage load the image base for certificate
func LoadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if nil != err {
		return nil, err
	}

	image, _, err := image.Decode(file)
	return image, nil
}

//StoreImageJPEG save the image in jpeg format
func StoreImageJPEG(nameImage string, base image.Image) {
	file, err := os.Create(nameImage)
	if nil != err {
		log.Print("Erro ao criar local para salvar imagem, ", err)
		return
	}
	defer file.Close()

	qualityImage := 90
	options := jpeg.Options{
		Quality: qualityImage,
	}
	err = jpeg.Encode(file, base, &options)
	if nil != err {
		log.Print("Erro ao salvar a imagem, ", err)
	}
}
