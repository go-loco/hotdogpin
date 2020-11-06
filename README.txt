INSTALLATION

Install go (go languge) for the server to run -> https://golang.org/doc/install
Install dependencies
  go get -u github.com/gin-gonic/gin
  go get -u github.com/stripe/stripe-go

Install Stripe CLI also

RUN 
go run *.go
To run Stripe CLI: stripe listen --forward-to http://localhost/webhook

TEST
On the browser type http://localhost/ 