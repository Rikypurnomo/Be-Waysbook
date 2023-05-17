package middleware

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
)

const iMg = "png"
const pDf = "pdf"

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("photo")
		if err != nil {
			file, err = c.FormFile("thumbnail")
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
		}

		data, err := handleUpload(file, iMg)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		c.Set("dataFile", data)
		return next(c)
	}
}

func UploadPdf(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("content")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		data, err := handleUpload(file, pDf)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		c.Set("dataPdf", data)
		return next(c)
	}
}

func handleUpload(file *multipart.FileHeader, ext string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	tempFile, err := ioutil.TempFile("uploads", "file-*."+ext)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	if _, err = io.Copy(tempFile, src); err != nil {
		return "", err
	}

	data := tempFile.Name()
	return data, nil
}
