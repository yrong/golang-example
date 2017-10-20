package main

import (
	"github.com/remogatto/prettytest"
	factory "github.com/yrong/demo/factory_method_with_reuse"
	"log"
)

type mySuite struct {
	prettytest.Suite
}

func main() {
	wa := factory.New(factory.Widget_A, factory.WidgetInfo{1001, "A"})
	wa.Add("thinga", "ma", "bop")
	//wa.Remove("bop")
	//wa.Remove("ma")
	//wa.Remove("xxx")

	wb := factory.New(factory.Widget_B, factory.WidgetInfo{1001, "A"})
	wb.Add("thinga", "ma", "bop")
	result := wa.IsEqual(wb)
	log.Println(result)

}

//func TestRunner(t *testing.T) {
//	prettytest.Run(t, new(mySuite))
//}
//
//func (s *mySuite) TestNew() {
//	wi := WidgetInfo{1001, "A"}
//	s.Equal(wi.name, "A")
//
//	w := Widget{WidgetInfo{1001, "A"}}
//	s.Equal(w.name, "A")
//
//	wa := New(Widget_A, WidgetInfo{1001, "A"})
//	s.Equal(wa.GetInfo().id, 1001)
//	s.Equal(wa.GetInfo().name, "A")
//}
//
//func (s *mySuite) TestAdd() {
//	wa := New(Widget_A, WidgetInfo{1001, "A"})
//	wa.Add("thinga", "ma", "bop", 1, 2)
//	s.Equal(wa.Size(), 5)
//}
//
//func (s *mySuite) TestRemove() {
//	wa := New(Widget_A, WidgetInfo{1001, "A"})
//	wa.Add("thinga", "ma", "bop")
//	wa.Remove("bop")
//	s.Equal(wa.Size(), 2)
//
//	wa.Remove("ma")
//	s.Equal(wa.Size(), 1)
//
//	wa.Remove("xxx")
//	s.Equal(wa.Size(), 1)
//}
//
//func (s *mySuite) TestIsEqual() {
//	wa := New(Widget_A, WidgetInfo{1001, "A"})
//	wa.Add("thinga", "ma", "bop")
//	wb := New(Widget_B, WidgetInfo{1001, "B"})
//	wb.Add("thinga", "ma", "bop")
//	s.Equal(wa.IsEqual(wb), false)
//
//	wc := New(Widget_A, WidgetInfo{1001, "A"})
//	wc.Add("thinga", "ma", "bop")
//	s.Equal(wa.IsEqual(wc), true)
//
//	wc.SetInfo(1001, "C")
//	s.Equal(wa.IsEqual(wc), false)
//
//	wc.SetInfo(1001, "A")
//	s.Equal(wa.IsEqual(wc), true)
//
//	wc.Remove("bop")
//	s.Equal(wa.IsEqual(wc), false)
//}
