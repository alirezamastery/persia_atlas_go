package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"persia_atlas/server/models"
	"strings"
)

type ProfileRequestBody struct {
	FirstName *string `json:"first_name" binding:"omitempty,good-name"`
	LastName  *string `json:"last_name"`
	Avatar    *string `json:"avatar"`
}

var goodName validator.Func = func(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	fmt.Println("validator good-name:", value, "ok:", ok)
	if ok {
		if value == "chaikin" {
			return false
		}
	}
	return true
}

func (h handler) UpdateProfile() gin.HandlerFunc {
	const profileImgDir = "profile"

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("good-name", goodName, false)
		if err != nil {
			log.Fatalln("error in registering validator:", err.Error())
			return nil
		}
	}

	return func(c *gin.Context) {
		user, ex := c.Get("user")
		if !ex {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get request user"})
		}

		oldProfile := user.(models.User).Profile
		newProfile := models.Profile{}
		contentType := c.Request.Header.Get("Content-Type")

		if strings.Contains(contentType, "multipart/form-data") {
			file, err := c.FormFile("avatar")
			if file != nil && err != nil {
				c.String(http.StatusBadRequest, "form file data err: %s", err.Error())
				return
			}
			if file != nil {
				fileExtension := filepath.Ext(file.Filename)
				filename := fmt.Sprintf("%s%s", uuid.New().String(), fileExtension)
				destination := filepath.Join(os.Getenv("MEDIA_PATH"), profileImgDir, filename)
				if err := c.SaveUploadedFile(file, destination); err != nil {
					c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
					return
				}
				newProfile.Avatar = fmt.Sprintf("/%s/%s/%s", os.Getenv("MEDIA_DIR"), profileImgDir, filename)
			} else {
				newProfile.Avatar = oldProfile.Avatar
			}

			firstName, exists := c.GetPostForm("first_name")
			if exists {
				newProfile.FirstName = firstName
			} else {
				newProfile.FirstName = oldProfile.FirstName
			}

			lastName, exists := c.GetPostForm("last_name")
			if exists {
				newProfile.LastName = lastName
			} else {
				newProfile.LastName = oldProfile.LastName
			}

			h.DB.Model(&oldProfile).Updates(newProfile)

		} else {
			var body = ProfileRequestBody{}
			jsonErr := c.BindJSON(&body)
			if jsonErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
				return
			}

			if body.FirstName == nil {
				newProfile.FirstName = oldProfile.FirstName
			} else {
				newProfile.FirstName = *body.FirstName
			}

			if body.LastName == nil {
				newProfile.LastName = oldProfile.LastName
			} else {
				newProfile.LastName = *body.LastName
			}

			if body.Avatar == nil {
				newProfile.Avatar = oldProfile.Avatar
			} else {
				newProfile.Avatar = *body.Avatar
			}

			h.DB.Model(&oldProfile).Updates(newProfile)
		}

		c.JSON(http.StatusOK, &newProfile)
	}
}
