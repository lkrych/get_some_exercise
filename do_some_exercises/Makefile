go:
	go test 
javascript:
	node -r esm exercises_test.js | ../node_modules/.bin/tap-dot 
python:
	python exercises_test.py 
c:
	gcc exercises_test.c -lcunit -o test ./test 
elixir:
	elixir -r exercises.ex exercises_test.exs
ocaml_build:
	ocamlbuild -pkgs OUnit exercises_test.byte
ocaml_run:
	./exercises_test.byte
