# GoWeber
## Preface
  > I'm coding this project!
  now this GoWeber can be used!
## Description
  A free GoLang framework for build a website, a simple frameword, quickly build your own website.
  
## How to use?
### Step one: Download the framework.
  > go get -u -v github.com/ZhuShaoQiang/GoWeber
  
  then you can find the GoWeber in
  >$GOPATH/github.com/ZhuShaoQiang/GoWeber
  
  but there is a tip: put GoWeber in your $GOPATH/src/ directory, so you can run this project like this:
  > go run $GOPATH/src/GoWeber/main/main.go
### Step tow: Do some changes!
  by default, this web will run at 0.0.0.0:8080, if you want to change it, you can find route.go in package route,
  then change the function: StartHTTPServer().
  but if you wanna run it at 0.0.0.0:80, you must have a permission of root, like
  > sudo go run ...
