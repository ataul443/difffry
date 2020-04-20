package diffry

type kStore struct {
	arr    []int
	offset int
}

//type trace struct {
//	editScriptLen int
//	traces        []*kStore
//	a          []byte
//	b          []byte
//}

//func (t *trace) push(ks *kStore){
//	k := newKStore(len(ks.arr)/2)
//	copy(k.arr, ks.arr)
//	t.traces = append(t.traces, k)
//}

func newKStore(size int) *kStore {
	return &kStore{make([]int, (size*2)+1), size}
}

func (k *kStore) get(idx int) int {
	if idx < 0 {
		return k.arr[k.offset-idx]
	}
	return k.arr[idx]
}

func (k *kStore) set(idx, val int) {
	if idx < 0 {
		k.arr[k.offset-idx] = val
		return
	}
	k.arr[idx] = val
}

//type point struct {
//	x, y int
//}

type move struct {
	oldpos point
	newpos point
}

//type snake struct {
//	moves []move
//	a, b []byte
//}

//func (s *snake) push(m move) {
//	s.moves = append(s.moves, m)
//}

//func newSnake(a, b []byte) *snake {
//	return &snake{
//		make([]move, 0),
//		a,
//		b,
//	}
//}

//func (s *snake) Print() {
//	for i := 0; i < len(s.moves) ; i++ {
//		oldp := s.moves[i].oldpos
//		newp := s.moves[i].newpos
//		fmt.Printf("(%d, %d) -> (%d, %d)\n", oldp.x, oldp.y, newp.x, newp.y)
//	}
//	fmt.Println()
//}

type line struct {
	nr   int
	data string
}

type edit struct {
	old    *line
	new    *line
	action string
	data   string
}

// type editSript struct {
// 	edits []*edit
// }

//func newEditScript(a, b string) *editSript {
//	s := genSnake(genTraces([]byte(a), []byte(b)))
//	//s.Print()
//	edits := genEdits(s)
//	es := &editSript{edits}
//	return es
//}

//func genTraces(a []byte, b []byte) *trace {
//	an := len(a)
//	bn := len(b)

//	maxEdits := an + bn

//	ks := newKStore(maxEdits)

//	tc := &trace{}
//	tc.a = a
//	tc.b = b

//	var gotToEnd bool
//	for d := 0; d <= maxEdits; d++ {
//		for k := -d; k <= d; k += 2 {

//			var farX, prevK int
//			shouldGoDown := k == -d || (k != d && ks.get(k-1) < ks.get(k+1))

//			if shouldGoDown {
//				prevK = k + 1
//			} else {
//				prevK = k - 1
//			}

//			farX = ks.get(prevK)
//			if !shouldGoDown {
//				farX += 1
//			}
//			farY := farX - k

//			// Advance farX, farY to end of the snake if exists
//			for farX < an && farY < bn && a[farX] == b[farY] {
//				farX += 1
//				farY += 1
//			}

//			if farX >= an && farY >= bn {
//				gotToEnd = true
//				break
//			}

//			ks.set(k, farX)
//		}
//		tc.push(ks)

//		if gotToEnd {
//			break
//		}
//	}
//	return tc
//}

//func genSnake(tc *trace) *snake {
//	x := len(tc.a)
//	y := len(tc.b)

//	k := x - y

//	s := newSnake(tc.a, tc.b)

//	for d := len(tc.traces)-1; d > 0; d-- {

//		t := tc.traces[d-1]

//		shouldGoUp := k == -d || k != d && t.get(k-1) < t.get(k+1)

//		var prevX, prevY, prevK int

//		if shouldGoUp {
//			prevK = k + 1
//		} else {
//			prevK = k - 1
//		}

//		prevX = t.get(prevK)
//		prevY = prevX - prevK

//		for x > prevX && y > prevY {
//			mv := move{point{x-1, y-1}, point{x, y}}
//			s.push(mv)
//			x -= 1
//			y -= 1
//		}

//		mv := move{point{prevX, prevY}, point{x, y}}
//		s.push(mv)
//		x = prevX
//		y = prevY
//		k = prevK

//		if prevX == 0 && prevY == 0 {
//			break
//		}
//	}

//	return s
//}

//func genEdits(s *snake) []*edit {
//	edits := make([]*edit, len(s.moves))
//	n := len(s.moves)

//	newLnr := 0
//	eIdx := 0
//	for i := n-1; i >= 0; i-- {
//		oldp := s.moves[i].oldpos
//		newp := s.moves[i].newpos

//		var e edit
//		if oldp.y == newp.y {
//			e = edit{
//				old: &line{oldp.x + 1, string(s.a[oldp.x])},
//				new: nil,
//				action: "-",
//				data: string(s.a[oldp.x]),
//			}
//		} else if oldp.x == newp.x {
//			newLnr += 1
//			e = edit{
//				old: nil,
//				new: &line{newLnr, string(s.b[oldp.y])},
//				action: "+",
//				data: string(s.b[oldp.y]),
//			}
//		} else {
//			newLnr += 1
//			e = edit{
//				old: &line{oldp.x, string(s.a[oldp.x])},
//				new: &line{newLnr, string(s.b[oldp.y])},
//				action: "=",
//				data: string(s.a[oldp.x]),
//			}
//		}
//		edits[eIdx] = &e
//		eIdx += 1
//	}
//	return edits
//}

//func printEditscript(es *editSript){
//	for i := 0; i < len(es.edits); i++ {
//		e := es.edits[i]

//		soldlnr := " "
//		snewlnr := " "

//		if e.old != nil {
//			soldlnr = strconv.Itoa(e.old.nr)
//		}

//		if e.new != nil {
//			snewlnr = strconv.Itoa(e.new.nr)
//		}

//		o := fmt.Sprintf("  %s    %s    %s    %s", e.action, soldlnr, snewlnr, e.data)

//		if e.action == "+" {
//			o = aurora.Green(o).String()
//		} else if e.action == "-" {
//			o = aurora.Red(o).String()
//		}

//		fmt.Println(o)
//	}
//}

///*
//ABCABBA
//CBABAC

//abbcghatt
//lokiubbcatt

//*/

//func Diff(a, b string) {
//	printEditscript(newEditScript(a, b))
//}
