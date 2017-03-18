package gp

//#GOGP_FILE_BEGIN
//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)
type Product interface{}
type GOGPParaType int
type GOGPIdType uint32

type Creater interface {
	Create(GOGPPara GOGPParaType) Product
	Name() string
	Id() GOGPIdType
}

type GOGPGlobalNamePrefixFactory struct {
	lock bool //if the factory is lock for regist
}

func (this *GOGPGlobalNamePrefixFactory) RegCreater(creater Creater) error {
	return nil
}

func (this *GOGPGlobalNamePrefixFactory) Lock() {
	this.lock = true
}

func (this *GOGPGlobalNamePrefixFactory) GetCreaterById(id GOGPIdType) (creater Creater, ok bool) {
	return
}

func (this *GOGPGlobalNamePrefixFactory) GetCreaterByName(name string) (creater Creater, ok bool) {
	return
}

func (this *GOGPGlobalNamePrefixFactory) CreateById(id GOGPIdType, GOGPPara GOGPParaType) (product Product, ok bool) {
	return
}

func (this *GOGPGlobalNamePrefixFactory) CreateByName(name string, GOGPPara GOGPParaType) (product Product, ok bool) {
	return
}

//#GOGP_FILE_END
