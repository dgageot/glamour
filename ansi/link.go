package ansi

import (
	"bytes"
	"fmt"
	"io"
	"net/url"

	"github.com/savioxavier/termlink"
)

// A LinkElement is used to render hyperlinks.
type LinkElement struct {
	BaseURL  string
	URL      string
	Children []ElementRenderer
}

func (e *LinkElement) Render(w io.Writer, ctx RenderContext) error {
	u, err := url.Parse(e.URL)
	if err == nil && "#"+u.Fragment != e.URL { // if the URL only consists of an anchor, ignore it
		token := resolveRelativeURL(e.BaseURL, e.URL)

		for _, child := range e.Children {
			var b bytes.Buffer
			if err := child.Render(&b, ctx); err != nil {
				return err
			}

			if _, err := fmt.Fprint(w, termlink.Link(b.String(), token)); err != nil {
				return err
			}
		}
	} else {
		// Original code
		for _, child := range e.Children {
			if r, ok := child.(StyleOverriderElementRenderer); ok {
				st := ctx.options.Styles.LinkText
				if err := r.StyleOverrideRender(w, ctx, st); err != nil {
					return err
				}
			} else {
				var b bytes.Buffer
				if err := child.Render(&b, ctx); err != nil {
					return err
				}
				el := &BaseElement{
					Token: b.String(),
					Style: ctx.options.Styles.LinkText,
				}
				if err := el.Render(w, ctx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
