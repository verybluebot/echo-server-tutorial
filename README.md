# Golnag Server using Echo Tutorials Series - [Link](https://www.youtube.com/watch?v=_pww3NJuWnk&list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY)
Sources for [tutorial series on Youtube](https://www.youtube.com/watch?v=_pww3NJuWnk&list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY) for creating a web server in Golang using the [Echo](https://github.com/labstack/echo) Package.

## Description 
The [Series](https://www.youtube.com/watch?v=_pww3NJuWnk&list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY) will go through most of [Echo](https://github.com/labstack/echo) features like requests types, CRUD, playing with headers, cookies
Middlewares and more.
Starting from the very basic 'hello world' and finishing with a full production server ready for a real
life app.

## Structure
Each part have a video and a branch of the end state, each next part starts on top of the one before it.
master is the last published part.

## Installation
Clone the project:
```
  git clone https://github.com/verybluebot/echo-server-tutorial.git
 ```
 cd into it:
 ```
  cd echo-server-tutorial
 ```
 
 define `GOPATH`:
 note: this is asuming you are inside the root of the project.
 ```
  echo GOPATH=`pwd` 
 ```
 
 build binaries: 
 ``` 
 go install main
 
 ```
 
 run the server: 
 
 ```
 bin/main
 
 ```
 
 Note: if you want to use anything other the `master` aka the last publish part, switch to
 its branch and build binaries.
 
 list all branches:
 ```
  git branch -a
 ```
 checkout to the branch you want and build/run like before:
 ```
  git checkout part_1_hello_world
 
  go install main
 
  bin/main
 ```
  
