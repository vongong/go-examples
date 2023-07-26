package main

import (
	"fmt"
	"log"
	"net/url"
	"path"
)

func main() {
	fmt.Println("Ex: 1")
	u, err := url.Parse("http://bing.com/auth")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "google.com"

	a := "10"
	b := "20"
	c := ""

	if a != "" {
		u.Path = path.Join(u.Path, "a", a)
	}
	if b != "" {
		u.Path = path.Join(u.Path, "b", b)
	}
	if c != "" {
		u.Path = path.Join(u.Path, "c", c)
	}

	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u)

	fmt.Println("Ex: 2")
	base, err := url.Parse("http://example.com")
	if err != nil {
		log.Fatal(err)
	}

	u, err = url.Parse("search")
	if err != nil {
		log.Fatal(err)
	}

	q = u.Query()
	q.Set("folder", "OHIO")
	u.RawQuery = q.Encode()

	fmt.Println(base.ResolveReference(u))

}
