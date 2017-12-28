package http

import (
	"net/http"
	"strings"

	"../config"
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
		if len(tos) == 0 {
			http.Error(w, "tos is null", http.StatusForbidden)
			return
		}
		subject := param.MustString(r, "subject")
		content := param.MustString(r, "content")
		tos = strings.Replace(tos, ",", ";", -1)

		s := smtp.New(cfg.Smtp.Addr, cfg.Smtp.Username, cfg.Smtp.Password)
		if(cfg.Smtp.Istls){
                err := s.SendMail4TLS(cfg.Smtp.From, tos, subject, content,"")
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
        } else {
                err := s.SendMail(cfg.Smtp.From, tos, subject, content,"")
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
        }
        http.Error(w, "success", http.StatusOK)
	})

}
