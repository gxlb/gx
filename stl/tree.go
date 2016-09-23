package stl

type TreeTravelOrder uint //walk tree order

const (
	TravelDeep      = 1 << iota //deep first
	TravelWide                  //wide first
	TravelRootFirst             //if travel root first
	TravelDefault   = TravelDeep | TravelRootFirst
)

type TreeNode struct {
	Content interface{}
	//	Path    string
	//	Dir     bool
	//	Size    FileSize

	parent   *TreeNode
	children []*TreeNode
}

type TravelTreeFunc func(node *TreeNode) error

func (this *TreeNode) ForEach(fun TravelTreeFunc) error {
	return nil
}

//new iterator with default order(deep first)
func (this *TreeNode) NewTraveler() *TreeTraveler {
	return this.NewIteratorO(TravelDefault)
}

//new iterator with order
func (this *TreeNode) NewTravelerO(order TreeTravelOrder) *TreeTraveler {
	brotherId := -1
	if order|TravelRootFirst == 0 { //travel root last
		brotherId = -2
	}
	return &TreeTraveler{Node: this, order: order, brotherId: brotherId}
}

//build a path tree
func Tree(path string) (n *TreeNode, e error) {
	return
}

//traveler
type TreeTraveler struct {
	Node      *TreeNode //this node
	order     TreeTravelOrder
	brotherId int //brother index in parent
}

//travel next node
func (this *TreeTraveler) Next() *TreeNode {
	if this.brotherId <= 0 {
		this.brotherId++
		return this.Node
	}
	return nil
}

func (this *TreeTraveler) Prev() *TreeNode {
	if this.brotherId <= 0 {
		this.brotherId++
		return this.Node
	}
	return nil
}
