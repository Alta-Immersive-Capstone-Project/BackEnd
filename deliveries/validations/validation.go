package validations

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type validation struct {
	v *validator.Validate
}

func NewValidation(v *validator.Validate) *validation {
	return &validation{
		v: v,
	}
}

func (v *validation) Validation(request interface{}) error {
	v.v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	v.v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	err := v.v.Struct(request)
	if err != nil {
		return err
	}

	return nil
}

func ValidationImage(files []*multipart.FileHeader) (string, error) {
	for i, file := range files {
		if file.Size >= 1000*1000 {
			return "max file image 1 MB", errors.New("max file image 1 MB")
		}
		fmt.Println(file.Size)
		src, err := file.Open()
		defer src.Close()
		if err != nil {
			return "eror open file", err
		}
		fileByte, _ := ioutil.ReadAll(src)
		fileType := http.DetectContentType(fileByte)
		if fileType != "image/png" && fileType != "image/jpeg" {
			return "type file image " + strconv.Itoa(i+1) + " not PNG or JPEG", errors.New("file image " + strconv.Itoa(i+1) + " type not PNG or JPEG")
		}
	}
	return "", nil
}

func ValidationAvatar(file *multipart.FileHeader) (string, error) {
	if file.Size >= 1000*1000 {
		return "max file image 1 MB", errors.New("max file image 1 MB")
	}
	fmt.Println(file.Size)
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		return "eror open file", err
	}
	fileByte, _ := ioutil.ReadAll(src)
	fileType := http.DetectContentType(fileByte)
	if fileType != "image/png" && fileType != "image/jpeg" {
		return "type file image not PNG or JPEG", errors.New("file image type not PNG or JPEG")
	}
	return "", nil
}
