package handlers

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/workradar/server/internal/services"
	"github.com/workradar/server/pkg/utils"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register godoc
type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Sanitize inputs (lowercase + trim)
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Username = strings.TrimSpace(req.Username)

	// Validate required fields
	if req.Email == "" || req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email, username, and password are required",
		})
	}

	// Max length validation
	if len(req.Password) > 128 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Password must not exceed 128 characters",
		})
	}
	if len(req.Username) > 50 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username must not exceed 50 characters",
		})
	}

	// Validate Gmail format
	validator := utils.NewInputValidator()
	validator.ValidateGmail(req.Email, "email")

	// Validate username - XSS prevention
	validator.ValidateNoXSS(req.Username, "username")
	validator.ValidateNoSQLInjection(req.Username, "username")

	if !validator.IsValid() {
		errors := validator.GetErrors()
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors[0].Message,
		})
	}

	// Log OTP request activity
	log.Printf("📝 OTP_REQUEST: action=register email=%s ip=%s", req.Email, c.IP())

	user, code, err := h.authService.Register(req.Email, req.Username, req.Password)
	if err != nil {
		log.Printf("❌ REGISTER_FAILED: email=%s ip=%s error=%s", req.Email, c.IP(), err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// REVISED: Return response without token - user must verify email first
	response := fiber.Map{
		"message":               "User registered successfully. Please verify your email.",
		"user":                  user.ToResponse(),
		"requires_verification": true,
	}

	// Only include code in development mode (when SMTP not configured)
	if code != "" {
		response["code"] = code
		response["dev_mode"] = true
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// Login godoc
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	MFACode  string `json:"mfa_code,omitempty"` // Optional MFA code
}

type LoginResponse struct {
	Message     string      `json:"message"`
	User        interface{} `json:"user,omitempty"`
	Token       string      `json:"token,omitempty"`
	RequiresMFA bool        `json:"requires_mfa,omitempty"`
	MFAToken    string      `json:"mfa_token,omitempty"`
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Sanitize email
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	// Max length validation for password
	if len(req.Password) > 128 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Password must not exceed 128 characters",
		})
	}

	// Use MFA-aware login
	result, err := h.authService.LoginWithMFA(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// If MFA is required, return partial response
	if result.RequiresMFA {
		return c.Status(fiber.StatusOK).JSON(LoginResponse{
			Message:     "MFA verification required",
			RequiresMFA: true,
			MFAToken:    result.MFAToken,
		})
	}

	return c.Status(fiber.StatusOK).JSON(LoginResponse{
		Message:     "Login successful",
		User:        result.User.ToResponse(),
		Token:       result.Token,
		RequiresMFA: false,
	})
}

// ForgotPassword godoc
type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

func (h *AuthHandler) ForgotPassword(c *fiber.Ctx) error {
	var req ForgotPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Sanitize email
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	if req.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email is required",
		})
	}

	// Log OTP request
	log.Printf("📝 OTP_REQUEST: action=forgot_password email=%s ip=%s", req.Email, c.IP())

	code, err := h.authService.ForgotPassword(req.Email)
	if err != nil {
		log.Printf("❌ FORGOT_PASSWORD_FAILED: email=%s ip=%s error=%s", req.Email, c.IP(), err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Response depends on whether email was sent or not
	// If code is empty, email was sent successfully (production mode)
	// If code is returned, SMTP not configured (development mode)
	response := fiber.Map{
		"message": "Verification code sent to email",
	}

	// Only include code in development mode (when SMTP not configured)
	if code != "" {
		response["code"] = code
		response["dev_mode"] = true
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// ResetPassword godoc
type ResetPasswordRequest struct {
	Code        string `json:"code"`
	NewPassword string `json:"new_password"`
}

func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	var req ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Code == "" || req.NewPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Code and new password are required",
		})
	}

	if err := h.authService.ResetPassword(req.Code, req.NewPassword); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password reset successful",
	})
}

// GetProfile mendapatkan data user yang sedang login
func (h *AuthHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	// TODO: Fetch user from repository
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id": userID,
	})
}

// UpdateProfile memperbarui profile user
type UpdateProfileRequest struct {
	Username       string  `json:"username"`
	ProfilePicture *string `json:"profile_picture"`
}

func (h *AuthHandler) UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.authService.UpdateProfile(userID, req.Username, req.ProfilePicture)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Profile updated successfully",
		"user":    user.ToResponse(),
	})
}

// ChangePassword mengubah password (untuk edit profile)
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Old and new password are required",
		})
	}

	if err := h.authService.ChangePassword(userID, req.OldPassword, req.NewPassword); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password changed successfully",
	})
}

// VerifyEmail godoc
type VerifyEmailRequest struct {
	Code string `json:"code"`
}

func (h *AuthHandler) VerifyEmail(c *fiber.Ctx) error {
	var req VerifyEmailRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Sanitize code
	req.Code = strings.TrimSpace(req.Code)

	if req.Code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Verification code is required",
		})
	}

	// Validate OTP format (must be 6 digits or with prefix)
	if !utils.IsValidOTPFormat(req.Code, true) {
		log.Printf("⚠️ OTP_INVALID_FORMAT: code=%s ip=%s", req.Code, c.IP())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid OTP format. Must be 6 digits.",
		})
	}

	// Log verification attempt
	log.Printf("📝 OTP_VERIFY: action=verify_email code=%s ip=%s", req.Code[:3]+"***", c.IP())

	if err := h.authService.VerifyEmail(req.Code); err != nil {
		log.Printf("❌ OTP_VERIFY_FAILED: code=%s ip=%s error=%s", req.Code[:3]+"***", c.IP(), err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Printf("✅ OTP_VERIFY_SUCCESS: ip=%s", c.IP())
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Email verified successfully. You can now login.",
	})
}

// ResendVerificationOTP godoc
type ResendOTPRequest struct {
	Email string `json:"email"`
}

func (h *AuthHandler) ResendVerificationOTP(c *fiber.Ctx) error {
	var req ResendOTPRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email is required",
		})
	}

	code, err := h.authService.ResendVerificationOTP(req.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := fiber.Map{
		"message": "Verification code sent to email",
	}

	// Only include code in development mode (when SMTP not configured)
	if code != "" {
		response["code"] = code
		response["dev_mode"] = true
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
