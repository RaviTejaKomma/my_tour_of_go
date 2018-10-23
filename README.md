# My tour of the Go tutorial

### Installing Go
- You can install Go via terminal or you can download the binary release from [here](https://golang.org/dl/)

- After Go has finished installing, we’ll need to create a workspace specifically for Go code.

### Setting up the workspace folder
- Create your Go folder via the terminal:
	- mkdir my_tour_of_go
- Change directory into your Go folder:
  - cd my_tour_of_go

- Go requires a specific folder structure in order to manage your local packages, binary files and your code.

- From within the Go folder enter the following commands:
  - mkdir bin pkg src
### Setting up the environment
- In order to run Go within your terminal, you’ll need to setup the GOPATH, GOBIN environment variables. This will allow Go to properly locate your workspace in order to run your code and manage your packages and binary files.
- Edit your ~/.bash_profile to add the following line:
  - export GOPATH=$HOME/go
  - export GOBIN=$GOPATH/bin
  - export PATH=$GOBIN:$PATH
- Save and exit your editor. Then, source your ~/.bash_profile.
  - source ~/.bash_profile
### Let’s create a Hello World Go program and make sure you can run it in your terminal.
- Enter the following commands:
  - cd src
  - mkdir hello_world
  - touch hello.go
- open your hello.go file in some editor like VSCode, Vim etc and the code.

```go
package main

import "fmt"

func main(){
	fmt.Println("Hello Go, I am here.")
}
```
- Go back to the hello_world folder in the terminal,  and type following command.
  - go run hello.go
- The proper output would be:
  - Hello Go, I am here.

If you received that output, you have properly setup your Go workspace on your machine

