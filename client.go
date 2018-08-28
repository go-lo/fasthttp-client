package fasthttpclient

import (
	"bufio"
	"bytes"
	"net/http"

	"github.com/valyala/fasthttp"
)

// Client wraps a fasthttp.Client, allowing for the casting of
// net/http types to fasthttp types.
type Client struct {
	fasthttpClient *fasthttp.Client
}

// New returns a new, fully instantiated Client which can be
// dropped in to a go-lo schedule
func New() *Client {
	return &Client{
		fasthttpClient: new(fasthttp.Client),
	}
}

// Do will action an http request and return stuff
func (c *Client) Do(in *http.Request) (out *http.Response, err error) {
	req := recastRequest(in)
	defer fasthttp.ReleaseRequest(req)

	resp := new(fasthttp.Response)

	err = c.fasthttpClient.Do(req, resp)
	if err != nil {
		return
	}

	return recastResponse(in, resp)
}

// recastRequest takes an http.Request, generated from usercode in a
// schedule, and turns it into a fasthttp.Request that the fasthttp
// client can use
func recastRequest(req *http.Request) *fasthttp.Request {
	output := fasthttp.AcquireRequest()

	output.Header.SetMethod(req.Method)
	for k, v := range req.Header {
		output.Header.Set(k, v[0])
	}

	output.SetRequestURI(req.URL.String())
	output.SetBodyStream(req.Body, int(req.ContentLength))

	return output
}

// recaseResponse takes a fasthttp.Response and turns it into an
// http.Response that our schedules can handle
func recastResponse(req *http.Request, resp *fasthttp.Response) (*http.Response, error) {
	buf := new(bytes.Buffer)

	reader := bufio.NewReader(buf)
	writer := bufio.NewWriter(buf)

	resp.Write(writer)
	writer.Flush()

	return http.ReadResponse(reader, req)
}
