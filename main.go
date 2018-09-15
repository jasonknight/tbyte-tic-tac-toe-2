package main
import (
	"fmt"
)
const UNOWNED = -1
const P1 = 0
const P2 = 1
type sq struct {
	display string
	owner int // UNOWNED, P1 or P2
	x,y int
}
type row []sq
type grid []row

func createGrid(size int) grid {
	var g grid
	for y := 0; y < size; y++ {
		var r row
		for x := 0; x < size; x++ {
			var s sq
			s.owner = UNOWNED
			s.display = "_|"
			s.x = x
			s.y = y
			r = append(r,s)
		}
		g = append(g,r)
	}
	return g
}
func displayGrid(g grid) {
	size := len(g)
		fmt.Print("   ")
	for i := 0; i <= size; i++ {
		fmt.Print("_,")
	}
	fmt.Print("\n")
	for y := 0; y < size; y++ {
		fmt.Print("   |")
		for x := 0; x < size; x++ {
			s := g[y][x]
			fmt.Print(s.display)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")

}
func play(g grid, player, y, x int) grid {
	g[y][x].owner = player	
	if player == P1 {
		g[y][x].display = "0|"
	} else {
		g[y][x].display = "X|"
	}
	return g
}
func getRowByOwner(g grid, p, y int) row {
	var r row
	for x := 0; x < len(g[y]); x++ {
		s := g[y][x]
		if s.owner == p {
			r = append(r,s)
		}
	}
	return r
}
func getColByOwner(g grid, p, x int) row {
	var r row
	for y := 0; y < len(g); y++ {
		s := g[y][x]
		if s.owner == p {
			r = append(r,s)
		}
	}
	return r
}
func getDiagByOwner(g grid, p, l_or_r int) row {
	var r row
	if l_or_r == 0 { 
		// left diag
		x := 0
		for y := 0; y < len(g); y++ {
			s := g[y][x]
			if s.owner == p {
				r = append(r,s)
			}
			x++
		}
	} else {
		// right diag
		x := len(g) - 1
		for y := 0; y < len(g); y++ {
			s := g[y][x]
			if s.owner == p {
				r = append(r,s)
			}
			x--
		}
	}
	return r
}
func getUnowned(g grid) row {
	var r row
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g); x++ {
			s := g[y][x]
			if s.owner == UNOWNED {
				r = append(r,s)
			}
		}
	}
	return r
}
func boardIsWon(g grid) (bool,int) {
	var owned row
	owners := []int{P1,P2}
	for i := 0; i < len(owners); i++ {
		for y := 0; y < len(g); y++ {
			owned = getRowByOwner(g,owners[i],y)
			if len(owned) == len(g) {
				return true,owners[i]
			}
		}
		for x := 0; x < len(g); x++ {
			owned = getColByOwner(g,owners[i],x)
			if len(owned) == len(g) {
				return true,owners[i]
			}
		}
		for j := 0; j < 2; j++ {
			owned = getDiagByOwner(g,owners[i],j)
			if len(owned) == len(g) {
				return true,owners[i]
			}
		}
	}
	return false,UNOWNED
}
func aiPlay(g grid) grid {
	state := 0
	last_len := 0
	the_x := 0
	the_y := 0
	the_diag := 0
	var owned row
	for y := 0; y < len(g); y++ {
		owned = getRowByOwner(g,P1,y)
		if len(owned) >= int(len(g)/2) {
			state = 0
			last_len = len(owned)
			the_y = y
		}
	}
	for x := 0; x < len(g); x++ {
		owned = getColByOwner(g,P1,x)
		if len(owned) >= int(len(g)/2) {
			if len(owned) > last_len {
				state = 1
				last_len = len(owned)
				the_x = x
			}
		}
	}
	for j := 0; j < 2; j++ {
		owned = getDiagByOwner(g,P1,j)
		if len(owned) >= int(len(g)/2) {
			if len(owned) > last_len {
				state = 2
				last_len = len(owned)
				the_diag = j
			}
		}
	}
	switch state {
		case 0:
			// the row is longest
			for x := 0; x < len(g); x++ {
				if g[the_y][x].owner == UNOWNED {
					return play(g,P2,the_y,x)
				}
			}
		case 1:
			// the column is longest
			for y := 0; y < len(g); y++ {
				if g[y][the_x].owner == UNOWNED {
					return play(g,P2,y,the_x)
				}
			}
		case 2:
			// the diag is longest
			var x int
			if the_diag == 0 { x = 0 } else { x = len(g)-1}
			for y := 0; y < len(g); y++ {
				if g[y][x].owner == UNOWNED {
					return play(g,P2,y,x)
				}
				if the_diag == 0 { x++ } else { x-- }
			}
		default:
			// need to play a random square
	}
	return g
}
