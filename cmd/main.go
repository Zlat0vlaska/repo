package main

import (
	"os"
	"reliab-test/internal/config"
	"reliab-test/internal/contracts/api"
	"reliab-test/internal/contracts/api/handlers"
	"reliab-test/internal/infrastructure/datastore"
	"reliab-test/internal/infrastructure/log"
	"reliab-test/internal/repositories"
	"reliab-test/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	defaultLog := log.InitDefaultLogger()

	cfg, err := config.Load(defaultLog)
	if err != nil {
		defaultLog.Error("Terminate execution", err)
	}

	logger := log.InitLogger()
	db := datastore.InitDB(cfg.DB, logger)
	r := gin.Default()

	userRepository := repositories.BuildUserRepository(db, logger)
	meetingRepository := repositories.BuildMeetingRepository(db)
	vacancyRepository := repositories.BuildVacancyRepository(db, logger)
	resumeRepository := repositories.BuildResumeRepository(db, logger)
	resumeVacancyRepository := repositories.BuildResumeVacancyRepository(db, logger)

	userService := services.BuildUserService(userRepository, logger)
	meetingService := services.BuildMeetingService(logger, meetingRepository, userRepository)
	vacancyService := services.BuildVacancyService(logger, vacancyRepository)
	resumeService := services.BuildResumeService(logger, resumeRepository)
	resumeVacancyService := services.BuildResumeVacancyService(logger, resumeVacancyRepository)

	userHandler := handlers.BuildUserHandler(userService, logger)
	meetingHandler := handlers.BuildMeetingHandler(meetingService, logger)
	vacancyHandler := handlers.BuildVacancyHandler(vacancyService, logger)
	resumeHandler := handlers.BuildResumeHandler(resumeService, logger)
	resumeVacancyHandler := handlers.BuildResumeVacancyHandler(resumeVacancyService, logger)

	server := api.BuildServer(r, logger, userHandler, meetingHandler, vacancyHandler, resumeHandler, resumeVacancyHandler)

	if err := server.Start(); err != nil {
		os.Exit(1)
	}
}
