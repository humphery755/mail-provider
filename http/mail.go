package http

import (
	"net/http"
	"strings"

	"github.com/humphery755/mail-provider/config"
	"github.com/humphery755/smtp"
	"github.com/toolkits/web/param"
)

func configProcRoutes() {

	http.HandleFunc("/sender/mail", func(w http.ResponseWriter, r *http.Request) {
		cfg := config.Config()
		token := param.String(r, "token", "")
		if cfg.Http.Token != token {
			http.Error(w, "no privilege", http.StatusForbidden)
			return
		}

		tos := param.MustString(r, "tos")
		subject := param.MustString(r, "subject")
		content := param.MustString(r, "content")
		tos = strings.Replace(tos, ",", ";", -1)

		s := smtp.New(cfg.Smtp.Addr, cfg.Smtp.Username, cfg.Smtp.Password)
		if(cfg.Smtp.Istls){
                err := s.SendMail4Tls(cfg.Smtp.From, tos, subject, content)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
        } else {
                err := s.SendMail(cfg.Smtp.From, tos, subject, content)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
        }
        http.Error(w, "success", http.StatusOK)
	})

}
