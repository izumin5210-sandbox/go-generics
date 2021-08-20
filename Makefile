.PHONY: run
run:
	go1.17 run -gcflags="-G=3" . > README.md
