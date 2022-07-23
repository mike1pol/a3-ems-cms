all: clean deps build
build: rms
clean:
	rm -rf main
deps:
	go get github.com/gorilla/handlers
	go get github.com/gorilla/mux
	go get github.com/go-pg/migrations
	go get github.com/robfig/cron
	go get github.com/go-pg/pg
	go get github.com/gobuffalo/envy
	go get github.com/oxtoacart/bpool
	go get golang.org/x/net/context
	go get github.com/PuerkitoBio/goquery
rms: *.go
	go build -o main .
lint:
	golint
