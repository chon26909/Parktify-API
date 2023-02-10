package utils

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/spf13/viper"
)

func UploadImage(file string) (interface{}, error) {

	cld, err := cloudinary.NewFromParams(
		viper.GetString("cloudinary.name"),
		viper.GetString("cloudinary.apiKey"),
		viper.GetString("cloudinary.apiSecret"))

	if err != nil {
		return nil, err
	}

	res, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{Folder: viper.GetString("cloudinary.folder")})

	return res, err
}
