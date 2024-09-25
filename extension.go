package glamour

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type HyperlinkRenderer struct {
	html.Config
}

func NewHyperlinkRenderer() renderer.NodeRenderer {
	return &HyperlinkRenderer{
		Config: html.NewConfig(),
	}
}

func (r *HyperlinkRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	//reg.Register(ast.KindLink, r.renderHyperlink)
	//reg.Register(ast.KindAutoLink, r.renderHyperlink)
}

func (r *HyperlinkRenderer) renderHyperlink(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("-START-")
	} else {
		_, _ = w.WriteString("-END-")
	}
	return ast.WalkContinue, nil
}
