package main

import (
	"fmt"
	"learngo/retriever/mooc"
	real2 "learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func post(poster Poster) {
	poster.Post("https://www.imooc.com", map[string]string{
		"name":   "aliang",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "https://www.imooc.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "another faked imooc.com"})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mooc.Retriever{Contents: "this is a fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Println(download(r))
	inspect(r)

	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)

	moocRetriever, ok := r.(*mooc.Retriever)
	if ok {
		fmt.Println(moocRetriever.Contents)
	} else {
		fmt.Println("not a mooc retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v \n", r, r)
	fmt.Printf("> Type switch")
	switch v := r.(type) {
	case *mooc.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
	fmt.Println()
}
