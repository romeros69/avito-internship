package app

import (
	"avito-internship/configs"
	_ "avito-internship/docs"
	balanceHttp "avito-internship/internal/myapp/balance/delivery/http/v1"
	balanceRepository "avito-internship/internal/myapp/balance/repository"
	balanceUseCase "avito-internship/internal/myapp/balance/usecase"
	historyHttp "avito-internship/internal/myapp/history/delivery/http/v1"
	historyRepository "avito-internship/internal/myapp/history/repository"
	historyUseCase "avito-internship/internal/myapp/history/usecase"
	reportRepository "avito-internship/internal/myapp/report/repository"
	reportUseCase "avito-internship/internal/myapp/report/usecase"
	reserveHttp "avito-internship/internal/myapp/reserve/delivery/http/v1"
	reserveRepository "avito-internship/internal/myapp/reserve/repository"
	reserveUseCase "avito-internship/internal/myapp/reserve/usecase"
	serviceRepository "avito-internship/internal/myapp/service/repository"
	serviceUseCase "avito-internship/internal/myapp/service/usecase"
	"avito-internship/internal/pkg/httpserver"
	"avito-internship/internal/pkg/postgres"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *configs.Config) {
	pg, err := postgres.New(cfg)

	if err != nil {
		log.Fatal("Error in creating postgres instance")
	}

	// http server
	handler := gin.New()

	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Init repositories
	balanceRepo := balanceRepository.NewBalanceRepo(pg)
	historyRepo := historyRepository.NewHistoryRepo(pg)
	reserveRepo := reserveRepository.NewReserveRepo(pg)
	reportRepo := reportRepository.NewReportRepo(pg)
	serviceRepo := serviceRepository.NewServiceRepo(pg)
	// Init useCases
	historySimpleUC := historyUseCase.NewSimpleHistoryUseCase(historyRepo)
	reportUC := reportUseCase.NewReportUseCase(reportRepo)
	serviceUC := serviceUseCase.NewServiceUseCase(serviceRepo)
	reserveSimpleUC := reserveUseCase.NewSimpleReserveUseCase(reserveRepo, historySimpleUC, reportUC, serviceUC)
	balanceUC := balanceUseCase.NewBalanceUseCase(balanceRepo, historySimpleUC, reserveSimpleUC)
	historyUC := historyUseCase.NewHistoryUseCase(historyRepo, balanceUC, serviceUC)
	reserveUC := reserveUseCase.NewReserveUseCase(reserveRepo, balanceUC, historyUC, reportUC, serviceUC)

	// Init handlers
	balanceHandlers := balanceHttp.NewBalanceHandlers(balanceUC)
	reserveHandlers := reserveHttp.NewReserveHandlers(reserveUC)
	historyHandlers := historyHttp.NewHistoryHandlers(historyUC)

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := handler.Group("/api/v1")

	balanceGroup := v1.Group("balance")
	reserveGroup := v1.Group("reserve")
	historyGroup := v1.Group("history")

	reserveHttp.MapReserveRoutes(reserveGroup, reserveHandlers)
	balanceHttp.MapBalanceRoutes(balanceGroup, balanceHandlers)
	historyHttp.MapHistoryRoutes(historyGroup, historyHandlers)

	serv := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err = <-serv.Notify():
		log.Printf("Notify from http server")
	}

	err = serv.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}
}
