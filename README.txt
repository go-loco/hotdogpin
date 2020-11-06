****
INSTALLATION

Install go (go language) for the server to run -> https://golang.org/doc/install

Set GOPATH variable (old school) -> (from command line on mac)
cd
mkdir go
export GOPATH=~/go

Get the project and dependencies (stripe and gin-gonic). From the command line:
go get -u github.com/go-loco/hotdogpin

Install Stripe CLI:
https://stripe.com/docs/payments/handling-payment-events

****
RUN 

From the command line:
cd $GOPATH/src/github.com/go-loco/hotdogpin
go run *.go

Run Stripe CLI to mock webhooks -> stripe listen --forward-to http://localhost/webhook

****
TEST
On the browser type http://localhost/ 
It runs on port 80
