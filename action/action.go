package action

type ActionCfg struct {
	Name     string
	Type     string
	Para     string
	NextOk   string
	NextFail string
}

type Action interface {
	Init(cfg *ActionCfg)
	Execute() bool
}

type ActionManager struct {
}

func (this *ActionManager) RegAction(cfg *ActionCfg) {}
