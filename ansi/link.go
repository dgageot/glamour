package ansi

import (
	"bytes"
	"io"
	"net/url"
)

// A LinkElement is used to render hyperlinks.
type LinkElement struct {
	BaseURL  string
	URL      string
	Children []ElementRenderer
}

func (e *LinkElement) Render(w io.Writer, ctx RenderContext) error {
	u, err := url.Parse(e.URL)
	if err == nil && "#"+u.Fragment != e.URL {
		for _, child := range e.Children {
			// Description
			var b bytes.Buffer
			if err := child.Render(&b, NewRenderContext(Options{})); err != nil {
				return err
			}
			description := b.String()

			// Show url
			token := resolveRelativeURL(e.BaseURL, e.URL)
			link := "\x1b]8;;" + token + "\x07" + description + "\x1b]8;;\x07" + "\u001b[0m"

			el := &BaseElement{
				Token:  link,
				Prefix: "",
				Style:  ctx.options.Styles.Link,
			}

			if err := el.Render(w, ctx); err != nil {
				return err
			}
		}
		// if the URL only consists of an anchor, ignore it
	} else {
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
		el := &BaseElement{
			Token:  resolveRelativeURL(e.BaseURL, e.URL),
			Prefix: " ",
			Style:  ctx.options.Styles.Link,
		}
		if err := el.Render(w, ctx); err != nil {
			return err
		}
	}

	return nil
}
