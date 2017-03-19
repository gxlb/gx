package gp

//#GOGP_FILE_BEGIN

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)
type GOGPValueType int

func NewGOGPGlobalNamePrefixReger(name string) (*GOGPGlobalNamePrefixReger, error) {
	return nil, nil
}

type GOGPGlobalNamePrefixReger struct {
	regerId  uint8
	name     string
	subIdGen uint32
	regList  []*_GOGPGlobalNamePrefixRecord
}

type _GOGPGlobalNamePrefixRecord struct {
	//name string
	val GOGPValueType
	id  uint32
}

type GOGPGlobalNamePrefixRegerManager struct {
	list []*GOGPGlobalNamePrefixReger
}

//#GOGP_FILE_END
