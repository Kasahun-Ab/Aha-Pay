package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

// SendEmail sends a beautiful and visually appealing email with OTP displayed.
func SendEmail(to, subject, otp string) error {
	// HTML template with OTP and styling
	emailTemplate := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <style>
        body {
            font-family: 'Arial', sans-serif;
            background-color: #f4f7fa;
            color: #333;
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 600px;
            margin: 50px auto;
            background-color: #ffffff;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        }
        .header {
            text-align: center;
            padding-bottom: 20px;
            border-bottom: 2px solid #4CAF50;
        }
        .header h1 {
            color: #4CAF50;
            font-size: 28px;
            margin: 0;
        }
        .content {
            text-align: center;
            margin: 30px 0;
        }
        .otp-code {
            font-size: 36px;
            font-weight: bold;
            color: #F0C341;
            padding: 20px;
            background-color: #f8f8f8;
            border-radius: 8px;
            display: inline-block;
        }
        .footer {
            font-size: 12px;
            text-align: center;
            color: #777;
            margin-top: 40px;
        }
        .footer p {
            margin: 0;
        }
        .footer a {
            color: #4CAF50;
            text-decoration: none;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Password Reset Request</h1>
        </div>
        <div class="content">
            <p>Hi there,</p>
            <p>We received a request to reset your password. To complete the process, please use the following OTP:</p>
            <p class="otp-code">%s</p>
            <p>This OTP will expire in 15 minutes. If you didn't request a password reset, please ignore this email.</p>
        </div>
        <div class="footer">
            <p>Thank you for using our service!</p>
            <p>&copy; 2025 AhaPay Inc. | <a href="https://yourwebsite.com">Visit our website</a></p>
        </div>
    </div>
</body>
</html>`, subject, otp)

	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", "kybeltal848@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", emailTemplate)

	// Configure SMTP dialer
	d := gomail.NewDialer("smtp.gmail.com", 587, "kybeltal848@gmail.com", "msem uvob lkgb jnpx")

	// Send the email
	return d.DialAndSend(m)
}
