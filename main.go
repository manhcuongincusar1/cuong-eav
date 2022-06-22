package main

import (
	service "cuong-eav/core/service/kyc_service"
	"cuong-eav/domain"
	"cuong-eav/handler/httphandler"
	"cuong-eav/repository"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	domain.Setup()
	domain.Migrate()
}

func main() {

	kycRepo := repository.NewKycRepository()
	kycService := service.NewKycService(kycRepo)
	routersInit := httphandler.InitRouter()

	httphandler.RouterUser(routersInit, kycService)

	gin.SetMode(gin.DebugMode)

	endPoint := fmt.Sprintf(":%d", 8888)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		MaxHeaderBytes: maxHeaderBytes,
	}

	fmt.Printf("[info] currently server time %s\n", time.Now())
	fmt.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
