package main

import "fmt"
import "net"
import "net/url"

// URLs provide a uniform way to locate resources.
func main() { // We’ll parse this example URL, which includes a scheme, authentication info, host, port, path, query params, and query fragment.

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s) // Parse the URL and ensure there are no errors.
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme) // Accessing the scheme is straightforward.
	fmt.Println(u.User)   // User contains all authentication info; call Username and Password on this for individual values.
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	fmt.Println(u.Host) // The Host contains both the hostname and the port, if present. Use SplitHostPort to extract them.

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(u.Path) // Here we extract the path and the fragment after the #.
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery) // To get query params in a string of k=v format, use RawQuery. You can also parse query params into a map. The parsed query param maps are from strings to slices of strings, so index into [0] if you only want the first value.

	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
