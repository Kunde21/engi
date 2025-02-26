#Engi 
A cross-platform game engine written in Go following an interpretation of the Entity Component System paradigm. Engi is currently compilable for Mac OSX, Linux, Windows and the Web. With the release of Go 1.4, sporting Android and the inception of IOS compatibility, mobile will soon be added as a release target.

Currently documentation is pretty scarce, this is because I have not *completely* finalized the API and am about to go through a "pretification" process in order to increase elegance and usability. For a basic up to date example of most features, look to the demos/hello.go and demos/pong/pong.go files. These files are currently your best friends for learning engi, well and me (feel free to shoot me a DM or issue whenever you want!).

## Examples
If you want to see some projects using engi, check out the [demos folder](http://https://github.com/paked/engi/tree/master/demos). [Newbrict/OakTale](https://github.com/Newbrict/OakTale) is probably the best place to see the newer features (tilemaps, z-layers, physics)

## Getting in touch / Contributing
Currently we are active on the [gopher slack](http://gophers.slack.com) (You can request an invite here: http://bit.ly/go-slack-signup) under the [#engi](https://gophers.slack.com/messages/engi/) channel.

The roadmap is available on [trello here](https://trello.com/b/VdqrmpEz/github-com-paked-engi). If you wish to have something added, or want to discuss a new feature you should begin talking about it in either an issue or on the previously mentioned slack channel.

## Docs

Before you read the basic doc, here are a few notes for me (and other contributors) about ideas to achieve elegance

* A potential ```engi.Files.Add(engi.NewResource("world", "world.txt"), engi.NewResource("face", "data/face.png"))``` --DONE
* Initialize batch in a cleaner manner in custom worlds --DONE
* Neater systems adding
* Clean entity construction from an external file
* Add or re look at ```New__Component__()``` functions for SpaceComponent, RenderComponent, and others
* Automatically detect which systems should be added to based off component depends on
* Revisit the camera API and its locating in ```World{}```
* Presets (prefabs?) for the easy re-initialization of an entity pattern

##Installation
```go get -u github.com/paked/engi```

Install Dependencies (Debian/Ubuntu):  
```
sudo apt-get install libopenal1 libopenal-dev
sudo apt-get install libglu1-mesa-dev freeglut3-dev mesa-common-dev
sudo apt-get install libalut0 libalut-dev
sudo apt-get install mesa-common-dev xorg-dev libgl1-mesa-dev
go get
go get github.com/stretchr/testify/assert
```
*TODO* Verify and detail dependencies by development platform

##Getting Started
```
package main
   
import (
	"github.com/paked/engi"
)

type Game struct {
	engi.World
}

func (game *Game) Setup() {
	engi.SetBg(0xffffff)
	game.AddSystem(&engi.RenderSystem{})
}

func main() {
	engi.Open("Title", 800, 800, false, &Game{})
}

```

First we start off by declaring that it is a runnable file, then import the engi library. Inside the ```main()``` function we finish off by opening the window, the four parameters that are passed in are ```Window Title```, ```Window Width```, ```Window Height```, ```Fullscreen Mode (as a bool)``` and finally an instance of ```Game```.

If you were to run this code, a white 800x800 window would appear on your screen.


*TODO* Write about entities

*TODO* Write about components

*TODO* Write about systems





