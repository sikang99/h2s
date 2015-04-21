# Makefile for h2s
all: usage

PROGRAM=h2s
EDITOR=vim

edit e:
	$(EDITOR) main.go

readme md:
	$(EDITOR) README.md

build b:
	go build
	@ls -alF --color=auto

run r:
	./$(PROGRAM)

test t:
	./$(PROGRAM)

rebuild:
	rm -f ./$(PROGRAM)
	go build
	@ls -alF --color=auto

install i:
	go install

clean:
	rm -f ./$(PROGRAM)

make m:
	$(EDITOR) Makefile

# ---------------------------------------------------------------------------
git-hub gh:
	ssh -T git@github.com

git-update gu:
	git init
	git add README.md Makefile $(PROGRAM).go img/
	git commit -m "image added"
	git push -u https://sikang99@github.com/sikang99/$(PROGRAM) master

git-status gs:
	git status
	git log --oneline -5

git-origin go:
	git init
	git add README.md Makefile $(PROGRAM).go
	git commit -m "start to code"
	git remote add origin https://sikang99@github.com/sikang99/$(PROGRAM)
	git push -u origin master

gencert:
	openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 9999 -nodes

# ---------------------------------------------------------------------------
usage:
	@echo ""
	@echo "Makefile for '$(PROGRAM)', by Stoney Kang, 2015/04/12"
	@echo ""
	@echo "usage: make [edit|readme|build|run|test|rebuild|clean|git]"
	@echo "	edit    : edit source"
	@echo "	build   : compile source"
	@echo "	run     : execute $(PROGRAM)"
	@echo "	test    : test $(PROGRAM) options"
	@echo "	install : install $(PROGRAM) to $(GOPATH)/bin"
	@echo "	gu      : git upload $(PROGRAM) to github.com"
	@echo ""
