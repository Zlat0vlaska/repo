package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"reliab-test/internal/contracts/domain_errors"
	"reliab-test/internal/contracts/dto"
	"reliab-test/internal/contracts/utils"
	"reliab-test/internal/services"
	"strings"
)

type UserHandler struct {
	log         *slog.Logger
	userService *services.UserService
}

func BuildUserHandler(userService *services.UserService, log *slog.Logger) *UserHandler {
	return &UserHandler{userService: userService, log: log}
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := uh.userService.GetAllUsers(ctx)

	if err != nil {
		uh.log.Error("Failed to get all users", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	usersDto := dto.MapUsersToDto(users)

	utils.BuildSuccessResponse(c, usersDto)
}

func (uh *UserHandler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()
	path := c.Query("path")
	uh.log.Info("Data got", path)

	searchParams, err := getFioAndEmailFromStr(path)
	if err != nil {
		uh.log.Error("Got wrong params", err)
		utils.BuildErrorResponse(c, http.StatusBadRequest, domain_errors.BadRequest)
		return
	}

	users, err := uh.userService.GetUsers(ctx, searchParams)

	if err != nil {
		uh.log.Error("Failed to get users with directory type of USER", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	usersDto := dto.MapUsersToDto(users)

	uh.log.Info("Data sent", usersDto)
	utils.BuildSuccessResponse(c, usersDto)
}

func (uh *UserHandler) GetApplicants(c *gin.Context) {
	ctx := c.Request.Context()
	path := c.Query("path")
	uh.log.Info("Data got", path)

	searchParams, err := getFioAndEmailFromStr(path)
	if err != nil {
		uh.log.Error("Got wrong params", err)
		utils.BuildErrorResponse(c, http.StatusBadRequest, domain_errors.BadRequest)
		return
	}

	users, err := uh.userService.GetApplicants(ctx, searchParams)

	if err != nil {
		uh.log.Error("Failed to get users with directory type of APPLICANT", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	usersDto := dto.MapUsersToDto(users)

	uh.log.Info("Data sent", usersDto)
	utils.BuildSuccessResponse(c, usersDto)
}

func getFioAndEmailFromStr(searchStr string) ([]string, error) {
	params := strings.Split(searchStr, " ")
	if len(params) > 2 {
		return nil, fmt.Errorf("got too much params %v", params)
	}
	for _, param := range params {
		if len(param) < 3 {
			return nil, fmt.Errorf("params should have length 3 or bigger %v", params)
		}
	}
	return params, nil
}
