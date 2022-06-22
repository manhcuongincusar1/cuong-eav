package httphandler

import (
	e "cuong-eav/constants/entity"
	"cuong-eav/core/port"
	"cuong-eav/domain"
	"cuong-eav/domain/user"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterUser(r *gin.Engine, userService port.KycService) {
	apiv1 := r.Group("/api/v1")
	apiv1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Pong")
	})

	apiv1.GET("/users", func(ctx *gin.Context) {
		users, err := userService.GetUsers()
		if err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, ctx.Errors)
			return
		}
		ctx.JSON(http.StatusOK, users)
	})

	apiv1.POST("/users", func(ctx *gin.Context) {
		body := user.User{}
		if err := ctx.BindJSON(&body); err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, ctx.Errors)
			return
		}

		jsons, _ := json.MarshalIndent(body, "", "\t")
		fmt.Printf("Request: %s\n", jsons)

		if err := userService.CreateUser(&body, e.AttributeSetDefault, e.EntityTypeUser); err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, ctx.Errors)
			return
		}
		ctx.JSON(http.StatusAccepted, "User Created")
	})

	apiv1.GET("/attributes", func(ctx *gin.Context) {
		if attributes, err := userService.GetAttributes(e.AttributeSetDefault, e.EntityTypeUser); err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, ctx.Errors)
			return
		} else {
			ctx.JSON(http.StatusOK, attributes)
		}
	})

	apiv1.POST("/attributes", func(ctx *gin.Context) {
		body := domain.EavAttribute{}
		if err := ctx.BindJSON(&body); err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, ctx.Errors)
			return
		}
		if err := userService.CreateAttribute(&body); err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, ctx.Errors)
			return
		}
		ctx.JSON(http.StatusAccepted, "Attribute Created")
	})

}
