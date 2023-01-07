package visitor

import (
	"strings"
)

type GraphicVisitor interface {
	VisitLine(*Line)
	VisitText(*Text)
}

type SpookySkeletonVisitor struct {
	noopVisitor
}

func (t *SpookySkeletonVisitor) VisitText(text *Text) {
	text.Text = strings.TrimRight(text.Text, "\n")
	text.Text += " - VISITED BY THE SPOOKY SKELETON\n"
}

type SpookyLineVisitor struct {
	noopVisitor
}

func (l *SpookyLineVisitor) VisitLine(line *Line) {
	line.Text = strings.ReplaceAll(line.Text, "_", "~")
}

const skeletonhehe = `
░░░░░░░░░░░░▄▐░░░░░░
░░░░░░▄▄▄░░▄██▄░░░░░
░░░░░▐▀█▀▌░░░░▀█▄░░░
░░░░░▐█▄█▌░░░░░░▀█▄░
░░░░░░▀▄▀░░░▄▄▄▄▄▀▀░
░░░░▄▄▄██▀▀▀▀░░░░░░░
░░░█▀▄▄▄█░▀▀░░░░░░░░
░░░▌░▄▄▄▐▌▀▀▀░░░░░░░
▄░▐░░░▄▄░█░▀▀░░░░░░░ U HAVE BEEN SPOOKED BY THE
▀█▌░░░▄░▀█▀░▀░░░░░░░
░░░░░░░▄▄▐▌▄▄░░░░░░░
░░░░░░░▀███▀█░▄░░░░░
░░░░░░▐▌▀▄▀▄▀▐▄░░░░░ SPOOKY SKELETON
░░░░░▐▀░░░░░░▐▌░░░░░
░░░░░░█░░░░░░░░█░░░░
░░░░░▐▌░░░░░░░░░█░░░
░░░░░█░░░░░░░░░░▐▌░░ `

type SPOOKYSKELETON struct {
}

func (S SPOOKYSKELETON) VisitLine(line *Line) {
	line.Text = skeletonhehe
}

func (S SPOOKYSKELETON) VisitText(text *Text) {
	text.Text = skeletonhehe
}

/////////////////////////////////////////

type noopVisitor struct{}

func (n noopVisitor) VisitLine(*Line) {}
func (n noopVisitor) VisitText(*Text) {}
