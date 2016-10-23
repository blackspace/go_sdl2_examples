package maze

import (
	"errors"
	"github.com/veandco/go-sdl2/sdl"
	"time"
	"math/rand"
)

type Maze struct {
	Cells [][]Cell

}

func NewMaze(n int) (m *Maze) {
	m=&Maze{}

	m.Cells=make([]([]Cell),n,n)
	for i:=0;i<n;i++ {
		m.Cells[i]=make([]Cell,n,n)
		for j:=0;j<n;j++ {
			m.Cells[i][j]=Cell{255}
		}
	}
	return
}

func (m *Maze)Len() int {
	return len(m.Cells)
}

func (m *Maze)Get(x,y int) *Cell{
	if x+1>m.Len() || y+1>m.Len() {
		return nil
	}
	if x<0 || y<0 {
		return nil
	}
	return &m.Cells[x][y]
}


func (m *Maze)GetAnyClosedCell() (x,y int,ok bool){
	for y:=0;y<len(m.Cells);y++ {
		for x:=0;x<len(m.Cells);x++ {
			cell:=m.Get(x,y)
			if cell.IsClosed() {
				return x,y,true
			}
		}
	}

	return
}

func (m *Maze)LeftBottom() (x,y int) {
	return 0,m.Len()-1
}

func (m *Maze)RightTop() (x,y int) {
	return m.Len()-1,0
}

func (m *Maze)IsOpen(x,y,direction int) bool {
	cell:=m.Get(x,y)
	switch direction {
	case UP:
		next_cell:=m.Get(x,y-1)

		if next_cell!=nil {
			if cell.HasTop() || next_cell.HasBottom() {
				return false
			} else {
				return true
			}
		} else {
			panic(errors.New("Cant UP"))
		}
	case DOWN:
		next_cell:=m.Get(x,y+1)

		if next_cell!=nil {
			if cell.HasBottom() || next_cell.HasTop() {
				return false
			} else {
				return true
			}
		} else {
			panic(errors.New("Cant DOWN"))
		}
	case LEFT:
		next_cell:=m.Get(x-1,y)

		if next_cell!=nil {
			if cell.HasLeft() || next_cell.HasRight() {
				return false
			} else {
				return true
			}

		} else {
			panic(errors.New("Cant LEFT"))
		}
	case RIGHT:
		next_cell:=m.Get(x+1,y)

		if next_cell!=nil {
			if cell.HasRight() || next_cell.HasLeft() {
				return false
			} else {
				return true
			}
		} else {
			panic(errors.New("Cant RIGHT"))
		}
	}

	panic(errors.New("The direction must be one of UP DOWN LEFT RIGHT"))

}

func (m *Maze)GetOpenPointSet(cx,cy int,ps *PointSet){
	ps.Add(cx,cy)

	if m.Get(cx,cy-1)!=nil && !ps.HasPoint(cx,cy-1) && m.IsOpen(cx,cy,UP) {
		m.GetOpenPointSet(cx,cy-1,ps)

	}

	if m.Get(cx,cy+1)!=nil && !ps.HasPoint(cx,cy+1) && m.IsOpen(cx,cy,DOWN) {
		m.GetOpenPointSet(cx,cy+1,ps)

	}

	if m.Get(cx-1,cy)!=nil && !ps.HasPoint(cx-1,cy) && m.IsOpen(cx,cy,LEFT) {
		m.GetOpenPointSet(cx-1,cy,ps)

	}

	if m.Get(cx+1,cy)!=nil && !ps.HasPoint(cx+1,cy) && m.IsOpen(cx,cy,RIGHT) {
		m.GetOpenPointSet(cx+1,cy,ps)
	}

}

func (m *Maze)HasClosed() (result int) {
	for y:=0;y<len(m.Cells);y++ {
		for x:=0;x<len(m.Cells);x++ {
			if m.Cells[x][y].IsClosed() {
				result++
			}
		}
	}
	return
}

func (m *Maze)WillMakeEmptyArea(x,y,a int) bool {
	switch a {
	case UP:
		c0:=m.Get(x-1,y-1);	c1:=m.Get(x,y-1);	c4:=m.Get(x+1,y-1)
		c2:=m.Get(x-1,y);	c3:=m.Get(x,y);		c5:=m.Get(x+1,y)

		var may_left_empty,may_right_empty bool

		if c0!=nil && c1!=nil && c2!=nil && c3!=nil {
			left :=NewMaze(2)
			left.Cells[0][0]=*c0; left.Cells[1][0]=*c1
			left.Cells[0][1]=*c2; left.Cells[1][1]=*c3

			if left.IsOpen(0,0,DOWN) && left.IsOpen(0,0,RIGHT) && left.IsOpen(0,1,RIGHT) {
				may_left_empty=true
			}
		}

		if c4!=nil && c5!=nil && c1!=nil && c3!=nil {
			right :=NewMaze(2)
			right.Cells[0][0]=*c1; right.Cells[1][0]=*c4
			right.Cells[0][1]=*c3; right.Cells[1][1]=*c5

			if right.IsOpen(0,0,RIGHT) && right.IsOpen(0,1,RIGHT) && right.IsOpen(1,0,DOWN) {
				may_right_empty=true
			}

		}

		if may_left_empty || may_right_empty {
			return true
		} else {
			return  false
		}
	case DOWN:
		c0:=m.Get(x-1,y);	c1:=m.Get(x,y);		c4:=m.Get(x+1,y)
		c2:=m.Get(x-1,y+1);	c3:=m.Get(x,y+1);	c5:=m.Get(x+1,y+1)

		var may_left_empty,may_right_empty bool

		if c0!=nil && c1!=nil && c2!=nil && c3!=nil {
			left :=NewMaze(2)
			left.Cells[0][0]=*c0; left.Cells[1][0]=*c1
			left.Cells[0][1]=*c2; left.Cells[1][1]=*c3

			if left.IsOpen(0,0,DOWN) && left.IsOpen(0,0,RIGHT) && left.IsOpen(0,1,RIGHT) {
				may_left_empty=true
			}
		}

		if c4!=nil && c5!=nil && c1!=nil && c3!=nil {
			right :=NewMaze(2)
			right.Cells[0][0]=*c1; right.Cells[1][0]=*c4
			right.Cells[0][1]=*c3; right.Cells[1][1]=*c5

			if right.IsOpen(0,0,RIGHT) && right.IsOpen(0,1,RIGHT) && right.IsOpen(1,0,DOWN) {
				may_right_empty=true
			}
		}

		if may_left_empty || may_right_empty {
			return true
		} else {
			return false
		}
	case LEFT:
		c0:=m.Get(x-1,y-1);	c1:=m.Get(x,y-1);
		c2:=m.Get(x-1,y);	c3:=m.Get(x,y);
		c4:=m.Get(x-1,y+1);	c5:=m.Get(x,y+1)

		var may_up_empty,may_down_empty bool

		if c0!=nil && c1!=nil && c2!=nil && c3!=nil {
			up :=NewMaze(2)
			up.Cells[0][0]=*c0; up.Cells[1][0]=*c1
			up.Cells[0][1]=*c2; up.Cells[1][1]=*c3


			if up.IsOpen(0,0,RIGHT) && up.IsOpen(0,0,DOWN) && up.IsOpen(1,0,DOWN) {
				may_up_empty=true
			}
		}

		if c4!=nil && c5!=nil && c2!=nil && c3!=nil {
			down :=NewMaze(2)
			down.Cells[0][0]=*c2; down.Cells[1][0]=*c3
			down.Cells[0][1]=*c4; down.Cells[1][1]=*c5


			if down.IsOpen(0,0,DOWN) && down.IsOpen(0,1,RIGHT) && down.IsOpen(1,0,DOWN) {
				may_down_empty=true
			}
		}

		if may_up_empty || may_down_empty {
			return true
		} else {
			return false
		}

	case RIGHT:
		c0:=m.Get(x,y-1);	c1:=m.Get(x+1,y-1);
		c2:=m.Get(x,y);		c3:=m.Get(x+1,y);
		c4:=m.Get(x,y+1);	c5:=m.Get(x+1,y+1)

		var may_up_empty,may_down_empty bool

		if c0!=nil && c1!=nil && c2!=nil && c3!=nil {
			up :=NewMaze(2)
			up.Cells[0][0]=*c0; up.Cells[1][0]=*c1
			up.Cells[0][1]=*c2; up.Cells[1][1]=*c3

			if up.IsOpen(0,0,RIGHT) && up.IsOpen(0,0,DOWN) && up.IsOpen(1,0,DOWN) {
				may_up_empty=true
			}
		}

		if c4!=nil && c5!=nil && c2!=nil && c3!=nil {
			down :=NewMaze(2)
			down.Cells[0][0]=*c2; down.Cells[1][0]=*c3
			down.Cells[0][1]=*c4; down.Cells[1][1]=*c5

			if down.IsOpen(0,0,DOWN) && down.IsOpen(0,1,RIGHT) && down.IsOpen(1,0,DOWN) {
				may_down_empty=true
			}
		}

		if may_up_empty || may_down_empty {
			return true
		} else {
			return false
		}
	}

	panic(errors.New("The action must be one of UP DOWN LEFT RIGHT"))
}

func (m *Maze)Draw(r *sdl.Renderer,w int) {
	n:=m.Len()

	r.SetDrawColor(255,255,255,255)

	for y:=0;y<n;y++ {
		for x:=0;x<n;x++ {
			cell:=m.Get(x,y)

			location_x:=x*w
			location_y:=y*w

			if cell.HasTop() {
				r.DrawLine(location_x,location_y,location_x+w,location_y)
			}

			if cell.HasBottom() {
				r.DrawLine(location_x,location_y+w,location_x+w,location_y+w)
			}

			if cell.HasLeft() {
				r.DrawLine(location_x,location_y,location_x,location_y+w)
			}

			if cell.HasRight() {
				r.DrawLine(location_x+w,location_y,location_x+w,location_y+w)
			}
		}

	}
}

func (m *Maze)FindPath(x0,y0,x1,y1 int,path * PointStack)  {
	if x0==x1 && y0==y1 {
		path.Push(x0,y0)
		return
	}

	path.Push(x0,y0)

	if m.Get(x0,y0-1)!=nil&&m.IsOpen(x0,y0,UP)&&!path.HasPoint(x0,y0-1) {
		m.FindPath(x0,y0-1,x1,y1,path)

		if lx,ly,ok:=path.Last(); ok && lx==x1 && ly==y1 {
			return
		}

	}
	if m.Get(x0,y0+1)!=nil&&m.IsOpen(x0,y0,DOWN)&&!path.HasPoint(x0,y0+1) {
		m.FindPath(x0,y0+1,x1,y1,path)

		if lx,ly,ok:=path.Last(); ok &&lx==x1 && ly==y1 {
			return
		}

	}
	if m.Get(x0+1,y0)!=nil&&m.IsOpen(x0,y0,RIGHT)&&!path.HasPoint(x0+1,y0) {
		m.FindPath(x0+1,y0,x1,y1,path)

		if lx,ly,ok:=path.Last();ok && lx==x1 && ly==y1 {
			return
		}

	}
	if m.Get(x0-1,y0)!=nil&&m.IsOpen(x0,y0,LEFT)&&!path.HasPoint(x0-1,y0){
		m.FindPath(x0-1,y0,x1,y1,path)

		if lx,ly,ok:=path.Last(); ok &&lx==x1 && ly==y1 {
			return
		}

	}

	path.Pop()

	return
}

func BuildMaze(w int) * Maze {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	mm :=NewMaze(w)


	RESTART:
	x,y,_:=mm.GetAnyClosedCell()
	ww := NewWorm(mm)
	ww.GetInMaze(x,y)

	for {
		next_act_0 :=make([]int,0,4)
		next_act_1 :=make([]int,0,4)
		next_act_final :=make([]int,0,4)
		if ww.UpCell()!=nil && !mm.IsOpen(ww.CurrentX,ww.CurrentY,UP) {
			next_act_0 =append(next_act_0,UP)
		}

		if ww.DownCell()!=nil && !mm.IsOpen(ww.CurrentX,ww.CurrentY,DOWN) {
			next_act_0 =append(next_act_0,DOWN)
		}

		if ww.LeftCell()!=nil && !mm.IsOpen(ww.CurrentX,ww.CurrentY,LEFT)  {
			next_act_0 =append(next_act_0,LEFT)
		}

		if ww.RightCell()!=nil && !mm.IsOpen(ww.CurrentX,ww.CurrentY,RIGHT) {
			next_act_0 =append(next_act_0,RIGHT)
		}

		for _,a:=range next_act_0 {
			switch a {
			case UP:
				if !mm.WillMakeEmptyArea(ww.CurrentX,ww.CurrentY,UP) {
					next_act_1=append(next_act_1,UP)
				}
			case DOWN:
				if !mm.WillMakeEmptyArea(ww.CurrentX,ww.CurrentY,DOWN) {
					next_act_1=append(next_act_1,DOWN)
				}
			case LEFT:
				if !mm.WillMakeEmptyArea(ww.CurrentX,ww.CurrentY,LEFT) {
					next_act_1=append(next_act_1,LEFT)
				}
			case RIGHT:
				if !mm.WillMakeEmptyArea(ww.CurrentX,ww.CurrentY,RIGHT) {
					next_act_1=append(next_act_1,RIGHT)
				}
			}
		}


		for _,a:=range next_act_1 {
			switch a {
			case UP:
				ps:=NewPointSet()
				mm.GetOpenPointSet(ww.CurrentX,ww.CurrentY,ps)
				if !ps.HasPoint(ww.CurrentX,ww.CurrentY-1)  {
					next_act_final =append(next_act_final,UP)
				}
			case DOWN:
				ps:=NewPointSet()
				mm.GetOpenPointSet(ww.CurrentX,ww.CurrentY,ps)
				if !ps.HasPoint(ww.CurrentX,ww.CurrentY+1){
					next_act_final =append(next_act_final,DOWN)
				}
			case LEFT:
				ps:=NewPointSet()
				mm.GetOpenPointSet(ww.CurrentX,ww.CurrentY,ps)
				if !ps.HasPoint(ww.CurrentX-1,ww.CurrentY) {
					next_act_final =append(next_act_final,LEFT)
				}
			case RIGHT:
				ps:=NewPointSet()
				mm.GetOpenPointSet(ww.CurrentX,ww.CurrentY,ps)
				if !ps.HasPoint(ww.CurrentX+1,ww.CurrentY) {
					next_act_final =append(next_act_final,RIGHT)
				}
			}
		}

		if len(next_act_final)!=0 {
			switch next_act_final[r.Intn(len(next_act_final))]{
			case UP:
				ww.Up()
			case DOWN:
				ww.Down()
			case LEFT:
				ww.Left()
			case RIGHT:
				ww.Right()
			}
		} else {
			goto RESTART
		}

		if c:=mm.HasClosed();c==0 {
			break
		}
	}

	return mm
}



