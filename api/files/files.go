package files

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sthorer/api/ent"

	"github.com/sthorer/api/api/types"

	"github.com/sthorer/api/config"

	"github.com/labstack/echo/v4"
)

func Upload(c echo.Context) error {
	cc := c.(*types.Context)
	user := cc.Get(types.UserKey).(*ent.User)
	form, err := cc.MultipartForm()
	if err != nil {
		return err
	}

	var files []*ent.File
	for _, formFiles := range form.File {
		for _, formFile := range formFiles {
			if formFile.Size > config.FreeUploadLimit {
				return cc.JSON(http.StatusUnauthorized, "file size is superior to limit")
			}

			f, err := formFile.Open()
			if err != nil {
				return err
			}

			hash, err := cc.Shell.Add(f)
			if err != nil {
				return err
			}

			fmt.Println("Headers:", formFile.Header)

			file, err := cc.Client.File.Create().
				SetHash(hash).
				SetUser(user).
				SetSize(formFile.Size).
				SetMetadata(map[string]interface{}{
					"name": formFile.Filename,
					"size": formFile.Size,
				}).
				Save(context.Background())
			if err != nil {
				return err
			}

			log.Printf("successfuly added %s (hash: %s)\n", file.Metadata["name"], file.Hash)

			files = append(files, file)
		}
	}

	return cc.JSON(http.StatusOK, files)
}
