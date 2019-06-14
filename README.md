## 🏃‍♀️ Get Some Exercise 🏃

You need go installed to run the exercise generator

`brew install go`

Run `go run exercise_generator.go` to get started.

This script will create a `./do_some_exercises` folder with the file `exercises.filesuffixyouchoose` and `exercises_test.filesuffixyouchoose`.

To test your code, navigate into the `./do_some_exercises` folder and use the makefile to run your tests

* Go: `make go`
* Python: `make python`
* C: `make c` ~ no exercises yet
* Elixir: `make elixir` ~ no exercises yet
* OCAML: `make ocaml`

Fill in the exercises in `./do_some_exercises/exercises` until all your tests pass.

### Dependencies

Each language you use in this repo has it's own dependencies. If you want to run the exercises you are going to need to install the languages and their dependencies. You do NOT need to install all of them. Just the ones you are going to use.

I wrote this on a Mac. To get all the required dependencies, use homebrew

### js
* `brew install node`
* `brew install yarn`
* `yarn install`

### python
* `brew install python`

### elixir
* `brew install elixir`

### ocaml
* `brew install ocaml`

### c
* you need to install c if you don't have it
* `brew install cunit` - C unit testing framework

### Want to add exercises to the languages that already exist?

Use the `generate_ex_sol_test.py` script in the utilities folder. 

Coming soon: Programatically generate your own language support :)