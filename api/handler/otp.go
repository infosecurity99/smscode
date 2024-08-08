package handler

import (
	"context"
	"net/http"
	"test/api/models"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateOTP godoc
// @Router       /auth/otp [POST]
// @Summary      Create OTP
// @Description  Create OTP for user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user_id body models.CreateOTPRequest true "User ID"
// @Success      201  {object}  models.OTP
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateOTP(c *gin.Context) {
	request := models.CreateOTPRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := c.ShouldBindJSON(&request); err != nil {
		handleResponse(c, h.log, "Invalid input data", http.StatusBadRequest, err.Error())
		return
	}

	otp, err := h.services.OTPService().CreateOTP(ctx, request)
	if err != nil {
		handleResponse(c, h.log, "Error while creating OTP", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "OTP created successfully", http.StatusCreated, otp)
}

// VerifyOTP godoc
// @Router       /auth/otp/verify [POST]
// @Summary      Verify OTP
// @Description  Verify OTP for user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user_id body models.VerifyOTPRequest true "User ID and OTP Code"
// @Success      200  {object}  bool
// @Failure      400  {object}  models.Response
// @Failure      401  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) VerifyOTP(c *gin.Context) {
	request := models.VerifyOTPRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := c.ShouldBindJSON(&request); err != nil {
		handleResponse(c, h.log, "Invalid input data", http.StatusBadRequest, err.Error())
		return
	}

	valid, err := h.services.OTPService().VerifyOTP(ctx, request)
	if err != nil {
		handleResponse(c, h.log, "Error while verifying OTP", http.StatusInternalServerError, err.Error())
		return
	}

	if !valid {
		handleResponse(c, h.log, "Invalid OTP", http.StatusUnauthorized, "Invalid OTP")
		return
	}

	handleResponse(c, h.log, "OTP is valid", http.StatusOK, "OTP is valid")
}
