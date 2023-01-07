package visitor

import "strings"

type GraphicVisitor interface {
	VisitLine(*Line)
	VisitText(*Text)
}

type SpookySkeletonVisitor struct {
	noop
}

func (t *SpookySkeletonVisitor) VisitText(text *Text) {
	text.Text = strings.TrimRight(text.Text, "\n")
	text.Text += " - VISITED BY THE SPOOKY SKELETON\n"
}

type noop struct{}

func (n noop) VisitLine(*Line) {}
func (n noop) VisitText(*Text) {}
