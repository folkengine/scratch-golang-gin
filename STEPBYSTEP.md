# Step by Step Guide 

This is a step by step guide to me working through Kulshekhar Kabra's [Building Go Web Applications and Microservices Using Gin
](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) tutorial.

## Adding Gogland to .gitignore

I've been a big fan of Jetbrain's [Gogland IDE](https://www.jetbrains.com/go/). 
Of course, every time you use a Jetbrains IDE on a git repo you've got to add
tge .idea directory to your .gitignore file. 

## Glide init

I've been using [Glide](https://glide.sh/) for my go dependency management.

```
$> cd scratch-golang-gin
$> glide init
```
 
## Install Gin

```
$> glide get github.com/gin-gonic/gin
```

## Create the Templates

Note that I'm using the .template extension for templates instead of .html. 
Using html extension for something that isn't a working html page just feels wrong to
this ol' web dev.