package hooks

import "github.com/sirupsen/logrus"

type LogrusHookCollection []LogrusLogHook

type LogrusLogHook struct {
	Hook             logrus.Hook
	IsProductionOnly bool
}

func NewHookCollection(
	reporter LogrusReporter,
) LogrusHookCollection {
	return []LogrusLogHook{
		{
			IsProductionOnly: true,
			Hook:             reporter,
		},
	}
}
