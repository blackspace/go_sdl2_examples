package maze

import "errors"

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type Worm struct {
	Maze * Maze
	CurrentX,CurrentY int
}

func NewWorm() *Worm {
	return &Worm{}
}

func (w *Worm)UpCell() *Cell {
	return w.Maze.Get(w.CurrentX,w.CurrentY-1)
}

func (w *Worm)Up()  {
	cell:=w.Maze.Get(w.CurrentX,w.CurrentY)
	next_cell :=w.UpCell()

	if next_cell !=nil  {
		cell.EraseTop()
		next_cell.EraseBottom()
		w.CurrentY=w.CurrentY-1

	} else {
		panic(errors.New("If Up,It will go out the maze"))
	}
}

func (w *Worm)DownCell()  *Cell  {
	return w.Maze.Get(w.CurrentX,w.CurrentY+1)
}

func (w *Worm)Down()  {
	cell:=w.Maze.Get(w.CurrentX,w.CurrentY)
	next_cell:=w.DownCell()

	if next_cell !=nil  {
		cell.EraseBottom()
		next_cell.EraseTop()
		w.CurrentY=w.CurrentY+1

	} else {
		panic(errors.New("If Down,It will go out the maze"))
	}

}

func (w *Worm)LeftCell() *Cell {
	return w.Maze.Get(w.CurrentX-1,w.CurrentY)
}

func (w *Worm)Left()  {
	cell:=w.Maze.Get(w.CurrentX,w.CurrentY)
	next_cell:=w.LeftCell()

	if next_cell !=nil  {
		cell.EraseLeft()
		next_cell.EraseRight()
		w.CurrentX=w.CurrentX-1
	} else {
		panic(errors.New("If Left,It will go out the maze"))
	}
}

func (w *Worm)RightCell() *Cell {
	return w.Maze.Get(w.CurrentX+1,w.CurrentY)
}

func (w *Worm)Right()  {
	cell:=w.Maze.Get(w.CurrentX,w.CurrentY)
	next_cell:=w.RightCell()

	if next_cell !=nil {
		cell.EraseRight()
		next_cell.EraseLeft()
		w.CurrentX=w.CurrentX+1

	} else {
		panic(errors.New("If Right,It will go out the maze"))
	}
}

func (w *Worm)GetInMaze(m * Maze) {
	w.Maze=m
	x,y:=m.LeftBottom()
	cell:=m.Get(x,y)
	cell.EraseLeft()
	w.CurrentX=x
	w.CurrentY=y
}

