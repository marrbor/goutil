package closer

import "fmt"

type Client interface {
	Close() error
}

// Close call specific client's close method. It is intended to be used with defer.
func Close(cl Client) {
	if err := cl.Close(); err != nil {
		fmt.Printf("%+v", err)
	}
}
