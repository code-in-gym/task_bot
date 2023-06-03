package email

import (
	"code-in-gym/task_bot/notify/base"
	"net"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

type Config struct {
	base.DefaultNotify
	Server string
	User   string
	Pass   string
	To     string
	From   string
}

func (c *Config) NewConfig(server, user, pass, to, from string) error {
	c.NotifyFunc = c.SendNotify
	c.Kind = "email"
	c.Server = server
	c.Pass = pass
	c.User = user
	c.To = to
	c.From = from
	return nil
}

func (c *Config) SendNotify(subject, content string) error {
	host, p, err := net.SplitHostPort(c.Server)
	if err != nil {
		return err
	}

	port, err := strconv.Atoi(p)
	if err != nil {
		return err
	}

	email := "Task" + "[" + c.User + "]"
	if c.From != "" {
		email = c.From
	}

	split := func(r rune) bool {
		return r == ';' || r == ','
	}

	recipients := strings.FieldsFunc(c.To, split)

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", recipients...)
	m.SetHeader("Subject", "[TaskBot] " + subject)
	m.SetBody("text/html; charset=UTF-8", content)

	d := gomail.NewDialer(host, port, c.User, c.Pass)
	err = d.DialAndSend(m)

	return err
}
