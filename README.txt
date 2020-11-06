INSTALLATION

Install go (go language) for the server to run -> https://golang.org/doc/install

Install go dependencies for this project. From the command line:
  go get -u github.com/gin-gonic/gin
  go get -u github.com/stripe/stripe-go

Install Stripe CLI also

CLONE THE PROJECT

RUN 
go run *.go
To run Stripe CLI: stripe listen --forward-to http://localhost/webhook

TEST
On the browser type http://localhost/ 
It runs on port 80
