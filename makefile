GOCMD      = go
GOBUILD    = $(GOCMD) build
GOCLEAN    = $(GOCMD) clean
GOTEST     = $(GOCMD) test
GOGET      = $(GOCMD) get
GOVET      = $(GOCMD) vet
GOGENERATE = $(GOCMD) generate
GOTOOL     = $(GOCMD) tool
GOLINT     = golint

BUILDENV =

TESTOPT = -v
LDFLAGS = -ldflags '-w -s'

DST  = tree
SRC  = tree.go

all: lint build test coverage cprofile mprofile

lint: $(SRC)
	$(GOVET) ./...
	$(GOLINT) ./...

mem.prof:
	$(GOTEST) $(TESTOPT) -memprofile mem.prof

cpu.prof:
	$(GOTEST) $(TESTOPT) -cpuprofile cpu.prof

cover.prof:
	$(GOTEST) $(TESTOPT) -coverprofile cover.prof

mprofile: mem.prof
	$(GOTOOL) pprof --text $+

cprofile: cpu.prof
	$(GOTOOL) pprof --text $+

coverage: cover.prof
	$(GOTOOL) cover -func=$<

test: $(SRC)
	$(GOTEST) ./...

$(DST): $(SRC)
	$(BUILDENV) $(GOBUILD) $(LDFLAGS) -o $@ ./cmd/...

build: $(DST)

clean:
	rm -f $(DST) *.prof *.test
