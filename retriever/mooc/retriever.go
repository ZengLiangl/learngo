package mooc

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever：{Contents = %s}", r.Contents)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return r.Contents
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
