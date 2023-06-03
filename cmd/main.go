package main

import (
	"flag"

	"code-in-gym/task_bot/config"
	"code-in-gym/task_bot/notify/email"
	"code-in-gym/task_bot/task"
)

func main() {
	var (
		configPath string
		server     string
		user       string
		pass       string
		to         string
		from       string
	)
	flag.StringVar(&configPath, "config", "daily.toml", "daily config file path")
	flag.StringVar(&server, "server", "", "email server")
	flag.StringVar(&user, "user", "", "email user")
	flag.StringVar(&pass, "pass", "", "email password")
	flag.StringVar(&to, "to", "", "email to")
	flag.StringVar(&from, "from", "", "email from")
	flag.Parse()

	// randomly choose some tasks
	config := config.NewConfig[any]()
	config.Load(configPath)

	// notify
	// Only support email now
	emailConfig := email.Config{}
	emailConfig.NewConfig(
		server, user, pass, to, from,
	)
	config.Notify.Email = append(config.Notify.Email, emailConfig)
	for _, e := range config.Notify.Email {
		for k, v := range config.Task {
			t := task.NewDailyTask(k, v)
			index := t.Random()
			e.NotifyFunc(k, t.Detail(index))
		}
	}
}
