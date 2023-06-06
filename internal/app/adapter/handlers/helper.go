package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func validateRequest[T any](c *gin.Context, requestStruct *T) error {

	err := c.ShouldBindJSON(requestStruct)
	if err != nil {
		return err
	}

	fmt.Println(requestStruct)

	err = validator.New().Struct(requestStruct)
	if err != nil {
		return err
	}

	return nil
}
