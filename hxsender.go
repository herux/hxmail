package hxmail

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

type HxMailSender struct {
	server   string
	port     int
	username string
	password string
}

type hxLoginCred struct {
	username string
	password string
}

func (a *hxLoginCred) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}

func (a *hxLoginCred) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func newSender(server string, port int, username, password string) *HxMailSender {
	return &HxMailSender{
		server:   server,
		port:     port,
		username: username,
		password: password,
	}
}

func loginCred(username, password string) smtp.Auth {
	return &hxLoginCred{username, password}
}

func (ms *HxMailSender) tryConnect() (auth smtp.Auth, err error) {
	host := fmt.Sprintf("%s:%d", ms.server, ms.port)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	client, err := smtp.NewClient(conn, ms.server)
	if err != nil {
		return nil, err
	}

	tlsconfig := &tls.Config{
		ServerName: ms.server,
	}

	if err = client.StartTLS(tlsconfig); err != nil {
		return nil, err
	}

	auth = loginCred(ms.username, ms.password)
	if err = client.Auth(auth); err != nil {
		return nil, err
	}

	return auth, nil
}

func (ms *HxMailSender) sendMail(auth smtp.Auth, mail *HxMail) error {
	host := fmt.Sprintf("%s:%d", ms.server, ms.port)
	to := strings.Split(mail.to, ",")

	msg, err := mail.toBytes()
	if err != nil {
		return err
	}

	err = smtp.SendMail(host, auth, ms.username, to, msg)
	if err != nil {
		return err
	}

	return nil
}
