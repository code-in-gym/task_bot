package base

type DefaultNotify struct {
	Kind string
	NotifyFunc func(string, string) error
}
