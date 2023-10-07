package ansi

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

// An ImageElement is used to render images elements.
type ImageElement struct {
	Text    string
	BaseURL string
	URL     string
	Child   ElementRenderer
}

func RenderToTermFromFile(path string, w io.Writer) error {
	cmd := exec.Command("viu", path)
  var out bytes.Buffer
  cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return err
	}
  w.Write(out.Bytes())
	return nil
}

func (e *ImageElement) Render(w io.Writer, ctx RenderContext) error {
	fullUrl := resolveRelativeURL(e.BaseURL, e.URL)
	if len(fullUrl) > 0 {
    fmt.Println("fullUrl:", fullUrl)
		RenderToTermFromFile(fullUrl, w)
	}
	if len(e.Text) > 0 {
		el := &BaseElement{
			Token: e.Text,
			Style: ctx.options.Styles.ImageText,
		}
		err := el.Render(w, ctx)
		if err != nil {
			return err
		}
	}
	if len(e.URL) > 0 {
		el := &BaseElement{
			Token:  resolveRelativeURL(e.BaseURL, e.URL),
			Prefix: " ",
			Style:  ctx.options.Styles.Image,
		}
		err := el.Render(w, ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
