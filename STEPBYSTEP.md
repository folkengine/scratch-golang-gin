# Step by Step Guide 

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