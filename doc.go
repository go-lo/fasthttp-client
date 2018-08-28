/*
package fasthttpclient provides an implementation of a golo.Client using fasthttp to make calls, rather than the default net/http client.
This may provide speedups for simple request sequences.

It can be instantiated as per:

  import client "github.com/go-lo/fasthttp-client"

And then used as:

  golo.Client = client.New()

*/
package fasthttpclient
