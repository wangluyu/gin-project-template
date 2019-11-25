package mail

import (
	"github.com/matcornic/hermes"
)

type Welcome struct {
}

func (w *Welcome) Email(name string, link string, info []hermes.Entry) hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: name,
			Intros: []string{
				"Welcome to XXX! We're very excited to have you on board.",
			},
			Dictionary: info,
			Actions: []hermes.Action{
				{
					Instructions: "To get started with XXX, please click here:",
					Button: hermes.Button{
						Text: "Confirm your account",
						Link: link,
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
}
