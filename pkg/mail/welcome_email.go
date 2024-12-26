package mail

import (
	"github.com/matcornic/hermes/v2"
	"github.com/sajadsalem/startkit/config"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	MailConfig *config.MailConfig
}

func NewMail(config *config.MailConfig) *Mail {
	return &Mail{
		MailConfig: config,
	}
}

func (m *Mail) Welcome(email, name, otp string) error {
	// Configure hermes by setting a theme and your product info
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Genie Supplier",
			Link: "http://localhost:8080",
			// Optional product logo
			Logo: "https://upload.wikimedia.org/wikipedia/commons/d/df/Go_gopher_app_engine_color.jpg",
		},
	}

	eh := hermes.Email{
		Body: hermes.Body{
			Name: name,
			Intros: []string{
				"Welcome to Genie Supplier! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Please copy your invite code:",
					InviteCode:   otp,
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	emailHTML, err := h.GenerateHTML(eh)
	if err != nil {
		return err
	}

	gm := gomail.NewMessage()
	gm.SetHeader("From", "armandofx7@gmail.com")
	gm.SetHeader("To", email)
	gm.SetHeader("Subject", "Welcome to Genie Supplier")
	gm.SetBody("text/html", emailHTML)

	d := gomail.NewDialer(m.MailConfig.Host, m.MailConfig.Port, m.MailConfig.Username, m.MailConfig.Password)

	if err := d.DialAndSend(gm); err != nil {
		return err
	}

	return nil
}
