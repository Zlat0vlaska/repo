package api

import (
	"log/slog"
	"net/http"
	"reliab-test/internal/contracts/api/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router               *gin.Engine
	log                  *slog.Logger
	userHandler          *handlers.UserHandler
	meetingHandler       *handlers.MeetingHandler
	vacancyHandler       *handlers.VacancyHandler
	resumeHandler        *handlers.ResumeHandler
	resumeVacancyHandler *handlers.ResumeVacancyHandler
}

func BuildServer(
	r *gin.Engine,
	log *slog.Logger,
	userHandler *handlers.UserHandler,
	meetingHandler *handlers.MeetingHandler,
	vacancyHandler *handlers.VacancyHandler,
	resumeHandler *handlers.ResumeHandler,
	resumeVacancyHandler *handlers.ResumeVacancyHandler,
) *Server {
	s := Server{
		Router:               r,
		log:                  log,
		userHandler:          userHandler,
		meetingHandler:       meetingHandler,
		vacancyHandler:       vacancyHandler,
		resumeHandler:        resumeHandler,
		resumeVacancyHandler: resumeVacancyHandler,
	}

	api := r.Group("/api")
	{
		api.GET("/users", s.userHandler.GetAllUsers)
		api.GET("/users/suggestions", s.userHandler.GetUsers)
		api.GET("/directories/suggestions", s.userHandler.GetApplicants)

		api.GET("/meetings", s.meetingHandler.GetMeetings)
		api.GET("/meetings/:id", s.meetingHandler.GetMeetingByID)
		api.DELETE("/meetings/:id", s.meetingHandler.DeleteMeetingByID)
		api.POST("/meetings", s.meetingHandler.CreateMeeting)

		api.GET("/vacancies", s.vacancyHandler.GetVacancies)
		api.GET("/vacancies/:id", s.vacancyHandler.GetVacancyByID)
		api.GET("/vacancies/search", s.vacancyHandler.GetVacancies)
		api.GET("/vacancies/filter", s.vacancyHandler.GetVacancies)
		api.GET("/vacancies/favorites", s.vacancyHandler.GetVacancies)

		api.GET("/resumes", s.resumeHandler.GetAllResumes)
		api.GET("/resumes/:id", s.resumeHandler.GetResumeByID)
		api.GET("/resumes/search", s.resumeHandler.GetAllResumes)

		api.GET("resume/:vacancy_id/vacancies", s.resumeVacancyHandler.GetResumesByVacancyHandler)
		api.POST("resume/:resume_id/vacancies/:vacancy_id", s.resumeVacancyHandler.UpdateLinkStatusHandler)
		api.POST("resume/:resume_id/vacancies/:vacancy_id/status", s.resumeVacancyHandler.CreateLinkHandler)

		api.GET("vacancy/:resume_id/resumes", s.resumeVacancyHandler.GetVacanciesByResumeHandler)

		api.POST("vacancy/:vacancy_id/resumes/:resume_id", s.resumeVacancyHandler.UpdateLinkStatusHandler)
		api.POST("vacancy/:vacancy_id/resumes/:resume_id/status", s.resumeVacancyHandler.CreateLinkHandler)

	}
	return &s
}

func (s *Server) Start() error {
	s.log.Info("Starting server...")
	return http.ListenAndServe(":8080", s.Router)
}
