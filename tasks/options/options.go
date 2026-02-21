package options

type ActionType int

const (
	ActionHelp ActionType = iota
	ActionAdd
	ActionMarkDone
	ActionList
	ActionDelete
)

var actionName = map[ActionType]string{
	ActionHelp:     "help",
	ActionAdd:      "add",
	ActionMarkDone: "done",
	ActionList:     "list",
	ActionDelete:   "delete",
}

func (t ActionType) String() string {
	return actionName[t]
}

func actionTypeFromString(s string) ActionType {
	for value, index := range actionName {
		if index == s {
			return value
		}
	}
	return ActionHelp
}

func Parse(s []string) Options {
	action := ActionHelp
	option := Options{}
	var extra string
	if s != nil && len(s) != 0 {
		action = actionTypeFromString(s[0])
	}
	option.ActionType = action
	if len(s) >= 2 && (action == ActionAdd || action == ActionMarkDone || action == ActionDelete) {
		extra = s[1]
	}

	return Options{
		ActionType: action,
		extra:      extra,
	}
}

type Options struct {
	ActionType
	extra string
}
