package main
import "testing"

func TestCreateAndDisplay(t *testing.T) {
	size := 3
	g := createGrid(size)
	displayGrid(g)
}
func TestPlay(t *testing.T) {
	size := 5
	g := createGrid(size)
	g = play(g,P1,1,1)
	if g[1][1].owner != P1 {
		t.Errorf("expected p1 to = 1")
	}
	displayGrid(g)
	owned := getRowByOwner(g,P1,1)
	if len(owned) != 1 {
		t.Errorf("expected owned to have len of 1 but go %d",len(owned))
	}
	owned = getColByOwner(g,P1,1)
	if len(owned) != 1 {
		t.Errorf("expected owned to have len of 1 but go %d",len(owned))
	}
	owned = getDiagByOwner(g,P1,0)
	if len(owned) != 1 {
		t.Errorf("expected owned to have len of 1 but go %d",len(owned))
	}
	g = createGrid(size)
	g = play(g,P1,int(size/2),int(size/2))
	displayGrid(g)
	owned = getDiagByOwner(g,P1,0)
	if len(owned) != 1 {
		t.Errorf("expected owned diag to have len of 1 but go %d",len(owned))
	}
	g = createGrid(size)
	// Test winning by a row
	for x := 0; x < len(g); x++ {
		g = play(g,P1,1,x)
	}
	won,winner := boardIsWon(g)
	if won != true || winner != P1 {
		t.Errorf("expected true,%d got %v,%d",P1,won,winner)
	}
	displayGrid(g)
	// Test winning by a column
	g = createGrid(size)
	for y := 0; y < len(g); y++ {
		g = play(g,P2,y,size-1)
	}
	displayGrid(g)
	won,winner = boardIsWon(g)
	if won != true || winner != P2 {
		t.Errorf("expected true,%d got %v,%d",P2,won,winner)
	}
	// Test winning by a left diag
	g = createGrid(size)
	x := 0
	for y := 0; y < len(g); y++ {
		g = play(g,P2,y,x)
		x++
	}
	displayGrid(g)
	won,winner = boardIsWon(g)
	if won != true || winner != P2 {
		t.Errorf("expected true,%d got %v,%d",P2,won,winner)
	}
	// Test winning by a right diag
	g = createGrid(size)
	x = size - 1
	for y := 0; y < len(g); y++ {
		g = play(g,P2,y,x)
		x--
	}
	displayGrid(g)
	won,winner = boardIsWon(g)
	if won != true || winner != P2 {
		t.Errorf("expected true,%d got %v,%d",P2,won,winner)
	}
}
func TestAiPlay(t *testing.T) {
	size := 3
	// Test AI on row
	g := createGrid(size)
	g = play(g,P1,0,0)
	g = play(g,P1,0,1)
	g = aiPlay(g)
	displayGrid(g)
	p1_owned := getRowByOwner(g,P1,0)
	p2_owned := getRowByOwner(g,P2,0)
	if len(p1_owned) != 2 || len(p2_owned) != 1 {
		t.Errorf("ai failed to play blocking move")
		return
	}
	g = createGrid(size)
	g = play(g,P1,0,0)
	g = play(g,P1,0,2)
	g = aiPlay(g)
	displayGrid(g)
	p1_owned = getRowByOwner(g,P1,0)
	p2_owned = getRowByOwner(g,P2,0)
	if len(p1_owned) != 2 || len(p2_owned) != 1 {
		t.Errorf("ai failed to play blocking move")
		return
	}
	// Test AI on Column
	g = createGrid(size)
	g = play(g,P1,0,0)
	g = play(g,P1,1,0)
	g = aiPlay(g)
	displayGrid(g)
	p1_owned = getColByOwner(g,P1,0)
	p2_owned = getColByOwner(g,P2,0)
	if len(p1_owned) != 2 || len(p2_owned) != 1 {
		t.Errorf("ai failed to play blocking move")
		return
	}
	// Test AI on left diag
	g = createGrid(size)
	g = play(g,P1,0,0)
	g = play(g,P1,1,1)
	g = aiPlay(g)
	displayGrid(g)
	p1_owned = getDiagByOwner(g,P1,0)
	p2_owned = getDiagByOwner(g,P2,0)
	if len(p1_owned) != 2 || len(p2_owned) != 1 {
		t.Errorf("ai failed to play blocking move")
		return
	}
	// Test AI on left diag
	g = createGrid(size)
	g = play(g,P1,0,2)
	g = play(g,P1,1,1)
	g = aiPlay(g)
	displayGrid(g)
	p1_owned = getDiagByOwner(g,P1,1)
	p2_owned = getDiagByOwner(g,P2,1)
	if len(p1_owned) != 2 || len(p2_owned) != 1 {
		t.Errorf("ai failed to play blocking move")
		return
	}
}
