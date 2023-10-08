package glamour

import (
	"io/ioutil"
	"testing"
)

const markdown_img = "testdata/image_embed.md"

func TestImage(t *testing.T) {
	r, err := NewTermRenderer(
		WithStandardStyle("dark"),
    WithWordWrap(120),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Read markdown image: " + markdown_img)
	in, err := ioutil.ReadFile(markdown_img)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Render markdown image")
	b, err := r.Render(string(in))
	if err != nil {
		t.Fatal(err)
	}

  t.Logf("%x", b[len(b)-200:])
	t.Log(b)
}
