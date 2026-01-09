package services

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"github.com/resend/resend-go/v2"
	"github.com/workradar/server/internal/config"
)

// EmailService handles all email sending operations
type EmailService struct {
	resendClient *resend.Client
	fromEmail    string
	fromName     string
	apiKey       string
}

// NewEmailService creates a new email service instance
func NewEmailService() *EmailService {
	apiKey := config.AppConfig.ResendAPIKey
	return &EmailService{
		resendClient: resend.NewClient(apiKey),
		fromEmail:    config.AppConfig.SMTPFromEmail,
		fromName:     config.AppConfig.SMTPFromName,
		apiKey:       apiKey,
	}
}

// IsConfigured checks if Resend API is properly configured
func (s *EmailService) IsConfigured() bool {
	return s.apiKey != ""
}

// SendPasswordResetCode sends password reset verification code via email (6 digits)
func (s *EmailService) SendPasswordResetCode(toEmail, code string) error {
	if !s.IsConfigured() {
		log.Println("⚠️ Resend API not configured, skipping email send")
		return nil // Return nil untuk development mode
	}

	subject := "Workradar - Kode Reset Password"

	// HTML email template for PASSWORD RESET
	htmlTemplate := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="margin: 0; padding: 0; font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #f5f5f5;">
    <table width="100%" cellpadding="0" cellspacing="0" style="background-color: #f5f5f5; padding: 40px 0;">
        <tr>
            <td align="center">
                <table width="600" cellpadding="0" cellspacing="0" style="background-color: #ffffff; border-radius: 16px; box-shadow: 0 4px 20px rgba(0,0,0,0.1);">
                    <!-- Header - RED for Password Reset -->
                    <tr>
                        <td style="background: linear-gradient(135deg, #EF4444 0%, #DC2626 100%); padding: 40px; border-radius: 16px 16px 0 0; text-align: center;">
                            <h1 style="color: #ffffff; margin: 0; font-size: 28px; font-weight: 700;">
                                🔐 Reset Password
                            </h1>
                            <p style="color: rgba(255,255,255,0.9); margin: 10px 0 0 0; font-size: 14px;">
                                Workradar - Task Management App
                            </p>
                        </td>
                    </tr>
                    
                    <!-- Content -->
                    <tr>
                        <td style="padding: 40px;">
                            <h2 style="color: #1f2937; margin: 0 0 20px 0; font-size: 22px;">
                                Lupa Password?
                            </h2>
                            <p style="color: #6b7280; line-height: 1.6; margin: 0 0 30px 0;">
                                Kami menerima permintaan untuk reset password akun Workradar Anda. 
                                Gunakan kode 6 digit di bawah ini:
                            </p>
                            
                            <!-- Verification Code Box - 6 DIGITS -->
                            <div style="background: linear-gradient(135deg, #FEE2E2 0%, #FECACA 100%); border-radius: 12px; padding: 30px; text-align: center; margin: 0 0 30px 0; border: 2px solid #EF4444;">
                                <p style="color: #DC2626; font-size: 14px; margin: 0 0 10px 0; text-transform: uppercase; letter-spacing: 1px;">
                                    Kode Reset Password (6 Digit)
                                </p>
                                <h1 style="color: #B91C1C; font-size: 48px; letter-spacing: 12px; margin: 0; font-weight: 700; font-family: monospace;">
                                    {{.Code}}
                                </h1>
                            </div>
                            
                            <!-- Warning -->
                            <div style="background-color: #FEF3C7; border-left: 4px solid #F59E0B; padding: 15px; border-radius: 0 8px 8px 0; margin: 0 0 30px 0;">
                                <p style="color: #92400E; margin: 0; font-size: 14px;">
                                    ⚠️ <strong>Penting:</strong> Kode ini akan kadaluarsa dalam <strong>2 menit</strong>. 
                                    Jangan bagikan kode ini kepada siapapun.
                                </p>
                            </div>
                            
                            <p style="color: #6b7280; line-height: 1.6; margin: 0;">
                                Jika Anda tidak meminta reset password, abaikan email ini. 
                                Akun Anda tetap aman.
                            </p>
                        </td>
                    </tr>
                    
                    <!-- Footer -->
                    <tr>
                        <td style="background-color: #f9fafb; padding: 30px; border-radius: 0 0 16px 16px; text-align: center;">
                            <p style="color: #9ca3af; font-size: 12px; margin: 0 0 10px 0;">
                                Email ini dikirim otomatis oleh Workradar. Mohon jangan membalas email ini.
                            </p>
                            <p style="color: #9ca3af; font-size: 12px; margin: 0;">
                                © 2026 Workradar. All rights reserved.
                            </p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>
`

	// Parse template
	tmpl, err := template.New("email").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse email template: %w", err)
	}

	// Execute template with data
	var body bytes.Buffer
	data := struct {
		Code string
	}{
		Code: code,
	}
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute email template: %w", err)
	}

	// Send email via Resend
	return s.sendViaResend(toEmail, subject, body.String())
}

// SendAccountVerificationCode sends account verification code via email (4 digits)
func (s *EmailService) SendAccountVerificationCode(toEmail, code string) error {
	if !s.IsConfigured() {
		log.Println("⚠️ Resend API not configured, skipping email send")
		return nil // Return nil untuk development mode
	}

	subject := "Workradar - Verifikasi Akun Anda"

	// HTML email template for ACCOUNT VERIFICATION
	htmlTemplate := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="margin: 0; padding: 0; font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #f5f5f5;">
    <table width="100%" cellpadding="0" cellspacing="0" style="background-color: #f5f5f5; padding: 40px 0;">
        <tr>
            <td align="center">
                <table width="600" cellpadding="0" cellspacing="0" style="background-color: #ffffff; border-radius: 16px; box-shadow: 0 4px 20px rgba(0,0,0,0.1);">
                    <!-- Header - PURPLE for Account Verification -->
                    <tr>
                        <td style="background: linear-gradient(135deg, #6366F1 0%, #8B5CF6 100%); padding: 40px; border-radius: 16px 16px 0 0; text-align: center;">
                            <h1 style="color: #ffffff; margin: 0; font-size: 28px; font-weight: 700;">
                                ✉️ Verifikasi Email
                            </h1>
                            <p style="color: rgba(255,255,255,0.9); margin: 10px 0 0 0; font-size: 14px;">
                                Workradar - Task Management App
                            </p>
                        </td>
                    </tr>
                    
                    <!-- Content -->
                    <tr>
                        <td style="padding: 40px;">
                            <h2 style="color: #1f2937; margin: 0 0 20px 0; font-size: 22px;">
                                Selamat Datang! 🎉
                            </h2>
                            <p style="color: #6b7280; line-height: 1.6; margin: 0 0 30px 0;">
                                Terima kasih telah mendaftar di Workradar! 
                                Masukkan kode 4 digit di bawah ini untuk memverifikasi email Anda:
                            </p>
                            
                            <!-- Verification Code Box - 4 DIGITS -->
                            <div style="background: linear-gradient(135deg, #EEF2FF 0%, #E0E7FF 100%); border-radius: 12px; padding: 30px; text-align: center; margin: 0 0 30px 0; border: 2px solid #6366F1;">
                                <p style="color: #6366F1; font-size: 14px; margin: 0 0 10px 0; text-transform: uppercase; letter-spacing: 1px;">
                                    Kode Verifikasi (4 Digit)
                                </p>
                                <h1 style="color: #4F46E5; font-size: 56px; letter-spacing: 16px; margin: 0; font-weight: 700; font-family: monospace;">
                                    {{.Code}}
                                </h1>
                            </div>
                            
                            <!-- Warning -->
                            <div style="background-color: #FEF3C7; border-left: 4px solid #F59E0B; padding: 15px; border-radius: 0 8px 8px 0; margin: 0 0 30px 0;">
                                <p style="color: #92400E; margin: 0; font-size: 14px;">
                                    ⚠️ <strong>Penting:</strong> Kode ini akan kadaluarsa dalam <strong>2 menit</strong>. 
                                    Jangan bagikan kode ini kepada siapapun.
                                </p>
                            </div>
                            
                            <p style="color: #6b7280; line-height: 1.6; margin: 0;">
                                Jika Anda tidak membuat akun di Workradar, abaikan email ini.
                            </p>
                        </td>
                    </tr>
                    
                    <!-- Footer -->
                    <tr>
                        <td style="background-color: #f9fafb; padding: 30px; border-radius: 0 0 16px 16px; text-align: center;">
                            <p style="color: #9ca3af; font-size: 12px; margin: 0 0 10px 0;">
                                Email ini dikirim otomatis oleh Workradar. Mohon jangan membalas email ini.
                            </p>
                            <p style="color: #9ca3af; font-size: 12px; margin: 0;">
                                © 2026 Workradar. All rights reserved.
                            </p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>
`

	// Parse template
	tmpl, err := template.New("verification").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse email template: %w", err)
	}

	// Execute template with data
	var body bytes.Buffer
	data := struct {
		Code string
	}{
		Code: code,
	}
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute email template: %w", err)
	}

	// Send email via Resend
	return s.sendViaResend(toEmail, subject, body.String())
}

// SendVerificationCode - DEPRECATED: Use SendPasswordResetCode or SendAccountVerificationCode instead
func (s *EmailService) SendVerificationCode(toEmail, code string) error {
	// For backward compatibility, determine type by code length
	if len(code) == 4 {
		return s.SendAccountVerificationCode(toEmail, code)
	}
	return s.SendPasswordResetCode(toEmail, code)
}

// SendWelcomeEmail sends welcome email to new users
func (s *EmailService) SendWelcomeEmail(toEmail, userName string) error {
	if !s.IsConfigured() {
		log.Println("⚠️ Resend API not configured, skipping welcome email")
		return nil
	}

	subject := "Selamat Datang di Workradar! 🎉"

	htmlTemplate := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
</head>
<body style="margin: 0; padding: 0; font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #f5f5f5;">
    <table width="100%" cellpadding="0" cellspacing="0" style="background-color: #f5f5f5; padding: 40px 0;">
        <tr>
            <td align="center">
                <table width="600" cellpadding="0" cellspacing="0" style="background-color: #ffffff; border-radius: 16px; box-shadow: 0 4px 20px rgba(0,0,0,0.1);">
                    <tr>
                        <td style="background: linear-gradient(135deg, #6366F1 0%, #8B5CF6 100%); padding: 40px; border-radius: 16px 16px 0 0; text-align: center;">
                            <h1 style="color: #ffffff; margin: 0; font-size: 28px;">🎉 Selamat Datang!</h1>
                        </td>
                    </tr>
                    <tr>
                        <td style="padding: 40px;">
                            <h2 style="color: #1f2937; margin: 0 0 20px 0;">Halo, {{.UserName}}!</h2>
                            <p style="color: #6b7280; line-height: 1.6;">
                                Terima kasih telah bergabung dengan <strong>Workradar</strong>. 
                                Kami senang Anda memilih kami untuk membantu meningkatkan produktivitas Anda!
                            </p>
                            <p style="color: #6b7280; line-height: 1.6;">
                                Mulai sekarang, Anda dapat:
                            </p>
                            <ul style="color: #6b7280; line-height: 1.8;">
                                <li>📋 Membuat dan mengelola tugas harian</li>
                                <li>📊 Memantau beban kerja dengan grafik</li>
                                <li>🔔 Menerima pengingat otomatis</li>
                                <li>📅 Mengatur kalender personal</li>
                            </ul>
                            <p style="color: #6b7280; line-height: 1.6; margin-top: 30px;">
                                Selamat berproduktivitas! 💪
                            </p>
                        </td>
                    </tr>
                    <tr>
                        <td style="background-color: #f9fafb; padding: 30px; border-radius: 0 0 16px 16px; text-align: center;">
                            <p style="color: #9ca3af; font-size: 12px; margin: 0;">
                                © 2026 Workradar. All rights reserved.
                            </p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>
`

	tmpl, err := template.New("welcome").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse welcome template: %w", err)
	}

	var body bytes.Buffer
	data := struct {
		UserName string
	}{
		UserName: userName,
	}
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute welcome template: %w", err)
	}

	return s.sendViaResend(toEmail, subject, body.String())
}

// SendVIPUpgradeEmail sends confirmation email after VIP upgrade
func (s *EmailService) SendVIPUpgradeEmail(toEmail, userName, plan string) error {
	if !s.IsConfigured() {
		log.Println("⚠️ Resend API not configured, skipping VIP upgrade email")
		return nil
	}

	subject := "Selamat! Anda sekarang VIP Member 👑"

	htmlTemplate := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
</head>
<body style="margin: 0; padding: 0; font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #f5f5f5;">
    <table width="100%" cellpadding="0" cellspacing="0" style="background-color: #f5f5f5; padding: 40px 0;">
        <tr>
            <td align="center">
                <table width="600" cellpadding="0" cellspacing="0" style="background-color: #ffffff; border-radius: 16px; box-shadow: 0 4px 20px rgba(0,0,0,0.1);">
                    <tr>
                        <td style="background: linear-gradient(135deg, #F59E0B 0%, #D97706 100%); padding: 40px; border-radius: 16px 16px 0 0; text-align: center;">
                            <h1 style="color: #ffffff; margin: 0; font-size: 28px;">👑 VIP Member</h1>
                        </td>
                    </tr>
                    <tr>
                        <td style="padding: 40px;">
                            <h2 style="color: #1f2937; margin: 0 0 20px 0;">Halo, {{.UserName}}!</h2>
                            <p style="color: #6b7280; line-height: 1.6;">
                                Selamat! Akun Anda telah berhasil di-upgrade ke <strong>VIP Member</strong> 
                                dengan paket <strong>{{.Plan}}</strong>.
                            </p>
                            <p style="color: #6b7280; line-height: 1.6;">
                                Sekarang Anda dapat menikmati fitur eksklusif:
                            </p>
                            <ul style="color: #6b7280; line-height: 1.8;">
                                <li>📊 Grafik Mingguan & Bulanan</li>
                                <li>🔄 Pengaturan Repeat Task dengan End Date</li>
                                <li>⏰ Pilihan Reminder Fleksibel (5/10/15/30 menit)</li>
                                <li>☀️ Integrasi Cuaca</li>
                                <li>💪 Rekomendasi Kesehatan</li>
                            </ul>
                            <p style="color: #6b7280; line-height: 1.6; margin-top: 30px;">
                                Terima kasih atas dukungan Anda! 🙏
                            </p>
                        </td>
                    </tr>
                    <tr>
                        <td style="background-color: #f9fafb; padding: 30px; border-radius: 0 0 16px 16px; text-align: center;">
                            <p style="color: #9ca3af; font-size: 12px; margin: 0;">
                                © 2026 Workradar. All rights reserved.
                            </p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>
`

	tmpl, err := template.New("vip").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse VIP template: %w", err)
	}

	var body bytes.Buffer
	data := struct {
		UserName string
		Plan     string
	}{
		UserName: userName,
		Plan:     plan,
	}
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute VIP template: %w", err)
	}

	return s.sendViaResend(toEmail, subject, body.String())
}

// sendViaResend sends an HTML email using Resend API
func (s *EmailService) sendViaResend(to, subject, htmlBody string) error {
	req := &resend.SendEmailRequest{
		From:    fmt.Sprintf("%s <%s>", s.fromName, s.fromEmail),
		To:      []string{to},
		Subject: subject,
		Html:    htmlBody,
	}

	sent, err := s.resendClient.Emails.Send(req)
	if err != nil {
		log.Printf("❌ Failed to send email to %s via Resend: %v", to, err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("✅ Email sent successfully to %s (ID: %s)", to, sent.Id)
	return nil
}
