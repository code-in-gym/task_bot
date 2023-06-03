package notify

import "code-in-gym/task_bot/notify/email"

type Notify interface {
	SendNotify()
}

type Config struct {
	Email []email.Config `toml:"email"`
}
