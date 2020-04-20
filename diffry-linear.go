package diffry

import (
	"fmt"
	"math"
	"strconv"

	"github.com/logrusorgru/aurora"
)

type point struct {
	x, y int
}

func (p *point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type bound struct {
	beg  point
	end  point
	a, b []rune
}

type snake []point

func (s *snake) Empty() bool {
	if len(*s) != 0 {
		return false
	}
	return true
}

func (b *bound) width() int {
	return b.end.x - b.beg.x
}

func (b *bound) height() int {
	return b.end.y - b.beg.y
}

func (b *bound) Delta() int {
	return b.width() - b.height()
}

func (b *bound) Size() int {
	return b.width() + b.height()
}

func (b *bound) Empty() bool {
	if b.width() == 0 && b.height() == 0 {
		return true
	}
	return false
}

func forwardSnake(b bound, fks, bks *kStore, d int) snake {

	var midSnake snake

	for k := d; k >= -d; k -= 2 {

		var x, y, prevX, prevY, prevK int
		cameFromUp := k == -d || (k != d && fks.get(k-1) < fks.get(k+1))

		if cameFromUp {
			prevK = k + 1
		} else {
			prevK = k - 1
		}

		x = fks.get(prevK)
		prevX = x
		if !cameFromUp {
			x += 1
		}

		y = (x - b.beg.x) - k + b.beg.y
		if d == 0 || !cameFromUp {
			prevY = y
		} else {
			prevY = y - 1
		}

		// Advance x, y to end of the snake if exists
		for x < b.end.x && y < b.end.y && b.a[x] == b.b[y] {
			x += 1
			y += 1
		}

		fks.set(k, x)

		delta := b.Delta()
		c := k - delta
		if c >= -(d-1) && c <= (d-1) {
			if (delta&1) == 1 && y >= bks.get(c) {
				midSnake = append(midSnake, point{prevX, prevY})
				midSnake = append(midSnake, point{x, y})
				break
			}
		}
	}
	return midSnake
}

func backwardSnake(b bound, fks, bks *kStore, d int) snake {

	var midSnake snake

	for c := d; c >= -d; c -= 2 {

		var x, y, prevX, prevY, prevC int
		cameFromRight := c == -d || (c != d && bks.get(c-1) < bks.get(c+1))

		if cameFromRight {
			prevC = c + 1
		} else {
			prevC = c - 1
		}

		y = bks.get(prevC)
		prevY = y
		if !cameFromRight {
			y -= 1
		}

		delta := b.Delta()
		k := c + delta

		// Need to use k diagonal to get the correct y
		x = (y - b.beg.y) + k + b.beg.x
		if d == 0 || !cameFromRight {
			prevX = x
		} else {
			prevX = x + 1
		}

		// Advance x, y to start of the snake if exists
		for x > b.beg.x && y > b.beg.y && b.a[x-1] == b.b[y-1] {
			x -= 1
			y -= 1
		}

		bks.set(c, y)

		if k >= -d && k <= d {
			if (delta&1) == 0 && x <= fks.get(k) {
				midSnake = append(midSnake, point{x, y})
				midSnake = append(midSnake, point{prevX, prevY})
				break
			}
		}
	}
	return midSnake
}

func findMidSnake(b bound) snake {

	if b.Empty() {
		return make([]point, 0)
	}

	d := b.Size()/2 + 1
	fks := newKStore(d)
	fks.set(1, b.beg.x)
	bks := newKStore(d)
	bks.set(1, b.end.y)

	var midSnake snake
	for i := 0; i <= d; i++ {
		fs := forwardSnake(b, fks, bks, i)
		if !fs.Empty() {
			midSnake = fs
			break
		}

		bs := backwardSnake(b, fks, bks, i)
		if !bs.Empty() {
			midSnake = bs
			break
		}
	}
	return midSnake
}

func buildSnake(b bound) snake {
	var fullSnake snake

	midSnake := findMidSnake(b)

	if midSnake.Empty() {
		//fmt.Println()
		fullSnake = append(fullSnake, b.beg)
		return fullSnake
	}
	//fmt.Printf("(%d, %d), (%d, %d)\n", midSnake[0].x, midSnake[0].y, midSnake[1].x, midSnake[1].y)
	headBox := bound{b.beg, midSnake[0], b.a, b.b}
	tailBox := bound{midSnake[len(midSnake)-1], b.end, b.a, b.b}

	headSnake := buildSnake(headBox)
	tailSnake := buildSnake(tailBox)

	if !headSnake.Empty() {
		fullSnake = append(fullSnake, headSnake...)
	}

	if !tailSnake.Empty() {
		fullSnake = append(fullSnake, tailSnake...)
	}

	return fullSnake
}

func buildMoves(b bound, s snake) []move {
	moves := make([]move, 0)

	for i := 1; i < len(s); i++ {

		x := s[i-1].x
		y := s[i-1].y

		// Follows the diagonal if present
		for x < s[i].x && y < s[i].y && b.a[x] == b.b[y] {
			m := move{point{x, y}, point{x + 1, y + 1}}
			moves = append(moves, m)
			x += 1
			y += 1
		}

		xd := math.Abs(float64(x - s[i].x))
		yd := math.Abs(float64(y - s[i].y))

		if xd > yd {
			m := move{point{x, y}, point{x + 1, y}}
			x += 1
			moves = append(moves, m)
		}

		if xd < yd {
			m := move{point{x, y}, point{x, y + 1}}
			y += 1
			moves = append(moves, m)
		}

		// Follows the diagonal if present
		for x < s[i].x && y < s[i].y && b.a[x] == b.b[y] {
			m := move{point{x, y}, point{x + 1, y + 1}}
			moves = append(moves, m)
			x += 1
			y += 1
		}

	}

	return moves
}

func genEdits(b bound, moves []move) []*edit {
	edits := make([]*edit, len(moves))
	n := len(moves)

	newLnr := 0
	eIdx := 0
	for i := 0; i <= n; i++ {
		oldp := moves[i].oldpos
		newp := moves[i].newpos

		var e edit
		if oldp.y == newp.y {
			e = edit{
				old:    &line{oldp.x + 1, string(b.a[oldp.x])},
				new:    nil,
				action: "-",
				data:   string(b.a[oldp.x]),
			}
		} else if oldp.x == newp.x {
			newLnr += 1
			e = edit{
				old:    nil,
				new:    &line{newLnr, string(b.b[oldp.y])},
				action: "+",
				data:   string(b.b[oldp.y]),
			}
		} else {
			newLnr += 1
			e = edit{
				old:    &line{oldp.x, string(b.a[oldp.x])},
				new:    &line{newLnr, string(b.b[oldp.y])},
				action: "=",
				data:   string(b.a[oldp.x]),
			}
		}
		edits[eIdx] = &e
		eIdx += 1
	}
	return edits
}

func printEdits(edits []*edit) {
	for i := 0; i < len(edits); i++ {
		e := edits[i]

		soldlnr := " "
		snewlnr := " "

		if e.old != nil {
			soldlnr = strconv.Itoa(e.old.nr)
		}

		if e.new != nil {
			snewlnr = strconv.Itoa(e.new.nr)
		}

		o := fmt.Sprintf("  %s    %s    %s    %s", e.action, soldlnr, snewlnr, e.data)

		if e.action == "+" {
			o = aurora.Green(o).String()
		} else if e.action == "-" {
			o = aurora.Red(o).String()
		}

		fmt.Println(o)
	}
}

func Diff(sa, sb string) {
	a := []rune(sa)
	b := []rune(sb)

	box := bound{
		point{0, 0},
		point{len(a), len(b)},
		a,
		b,
	}
	snake := buildSnake(box)
	moves := buildMoves(box, snake)
	edits := genEdits(box, moves)

	printEdits(edits)
}
