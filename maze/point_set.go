package maze

type PointSet struct {
	Data []int
}

func NewPointSet() *PointSet {
	return &PointSet{Data:make([]int,0,1000)}
}

func (p *PointSet)Add(x,y int) {
	if !p.IsExists(x,y) {
		p.Data =append(p.Data,x,y)
	}
}


func (p *PointSet)IsExists(x,y int) bool {
	for i:=0;i<len(p.Data);i=i+2 {
		px:=p.Data[i]
		py:=p.Data[i+1]

		if px==x && py==y {
			return true
		}
	}

	return false
}