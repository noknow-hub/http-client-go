//////////////////////////////////////////////////////////////////////
// client.go
//////////////////////////////////////////////////////////////////////
package v1

import (
    "io/ioutil"
    "net/http"
    "strings"
    "time"
)


type Client struct {
    Config *Config
}


//////////////////////////////////////////////////////////////////////
// New client.
//////////////////////////////////////////////////////////////////////
func NewClient(url string) *Client {
    return &Client{
        Config: &Config{
            Url: ,url
        },
    }
}
