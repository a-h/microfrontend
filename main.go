package main

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/a-h/templ/turbo"
)

type articleContent struct{}

func (ac articleContent) Render(ctx context.Context, w io.Writer) error {
	_, err := io.WriteString(w, `This is the async loaded content`)
	return err
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 5)
	turbo.ReplaceWithContext(r.Context(), w, "article", articleContent{})
}

func main() {
	h := http.NewServeMux()
	h.Handle("/article/", http.HandlerFunc(articleHandler))
	h.Handle("/", templ.Handler(home()))
	http.ListenAndServe("localhost:8000", h)
}
