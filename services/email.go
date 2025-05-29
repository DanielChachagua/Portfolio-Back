package services

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/DanielChachagua/Portfolio-Back/models"
)

func SendEmail(emailContact *models.EmailContact) error {
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	from := os.Getenv("EMAIL_CONTACT")
	to := os.Getenv("EMAIL_CONTACT")
	pass := os.Getenv("EMAIL_PASSWORD")

	if host == "" || port == "" || from == "" || pass == "" || to == "" {
		return models.ErrorResponse(500, "Error al enviar el correo", nil)
	}

	auth := smtp.PlainAuth("", from, pass, host)

	htmlBody := BuildEmailHTML(emailContact)

	msg := "MIME-Version: 1.0\r\n"
	msg += "Content-Type: text/html; charset=\"UTF-8\"\r\n"
	msg += "Subject: PORTFOLIO\r\n"
	msg += "\r\n" + htmlBody

	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return models.ErrorResponse(500, "Error al enviar el correo", err)
	}

	return nil
}

func BuildEmailHTML(contact *models.EmailContact) string {
    return fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <head>
            <meta charset="UTF-8">
            <style>
                body { font-family: Arial, sans-serif; background: #f9f9f9; padding: 20px; }
                .container { background: #fff; border-radius: 8px; box-shadow: 0 2px 8px #eee; padding: 24px; max-width: 500px; margin: auto; }
                h2 { color: #2c3e50; }
                p { margin: 8px 0; }
                .label { font-weight: bold; color: #34495e; }
                .value { color: #555; }
                hr { border: none; border-top: 1px solid #eee; margin: 16px 0; }
            </style>
        </head>
        <body>
            <div class="container">
                <h2>Nuevo mensaje de contacto</h2>
                <hr>
                <p><span class="label">Nombre:</span> <span class="value">%s</span></p>
                <p><span class="label">Email:</span> <span class="value">%s</span></p>
                <p><span class="label">Teléfono:</span> <span class="value">%s</span></p>
                <p><span class="label">Asunto:</span> <span class="value">%s</span></p>
                <hr>
                <p><span class="label">Mensaje:</span></p>
                <p class="value">%s</p>
            </div>
        </body>
        </html>
    `, contact.Name, contact.Email, contact.Phone, contact.Issue, contact.Body)
}