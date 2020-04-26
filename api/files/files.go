package files

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sthorer/api/plans"
)

func Upload(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	for _, files := range form.File {
		for _, file := range files {
			if file.Size > plans.FreeUploadLimit {
				return c.JSON(http.StatusUnauthorized, "file size is superior to limit")
			}

			fmt.Println(file.Filename, file.Size)
		}
	}

	return c.NoContent(http.StatusOK)
}
