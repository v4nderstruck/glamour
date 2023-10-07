package ansi

import (
	"bytes"
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
  // bug in github.com/muesli/reflow cutting of escape
	cmd := exec.Command("viu", path)
  var out bytes.Buffer
  cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
    w.Write([]byte("RenderErr: " + path))
		return err
	}
  //fmt.Println(len(out.String()))
  w.Write([]byte(out.String()))
  w.Write([]byte("\n"))
	return nil
}

func (e *ImageElement) Render(w io.Writer, ctx RenderContext) error {
	fullUrl := e.BaseURL + e.URL
	if len(fullUrl) > 0 {
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
