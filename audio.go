package engi

import (
	"log"

	"golang.org/x/mobile/exp/audio/al"
)

const (
	defaultHeightModifier float32 = 1
)

// AudioComponent is a Component which is used by the AudioSystem
type AudioComponent struct {
	File       string
	Repeat     bool
	Background bool
	player     *Player
}

func (AudioComponent) Type() string {
	return "AudioComponent"
}

// AudioSystem is a System that allows for sound effects and / or music
type AudioSystem struct {
	*System
	HeightModifier float32
}

func (AudioSystem) Type() string {
	return "AudioSystem"
}

func (as *AudioSystem) New() {
	as.System = NewSystem()

	if as.HeightModifier == 0 {
		as.HeightModifier = defaultHeightModifier
	}

	if err := al.OpenDevice(); err != nil {
		log.Println("Error initializing AudioSystem:", err)
		return
	}

	Mailbox.Listen("CameraMessage", func(msg Message) {
		_, ok := msg.(CameraMessage)
		if !ok {
			return
		}

		// Hopefully not that much of an issue, when we receive it before the CameraSystem does
		// TODO: but it is when the CameraMessage is not Incremental (i.e. the changes are big)
		al.SetListenerPosition(al.Vector{cam.X() / Width(), cam.Y() / Height(), cam.Z() * as.HeightModifier})
	})
}

func (as *AudioSystem) Update(entity *Entity, dt float32) {
	var ac *AudioComponent
	if !entity.Component(&ac) {
		return
	}

	if ac.player == nil {
		f := Files.Sound(ac.File)
		if f == nil {
			return
		}

		var err error
		ac.player, err = NewPlayer(f, 0, 0)
		if err != nil {
			log.Println("Error initializing AudioSystem:", err)
			return
		}
	}

	if ac.player.State() != Playing {
		if ac.player.State() == Stopped {
			if !ac.Repeat {
				al.RewindSources(ac.player.source)
				al.StopSources(ac.player.source)
				entity.RemoveComponent(ac)
				return
			}
		}

		// Prepares if the track hasn't been buffered before.
		if err := ac.player.prepare(ac.Background, 0, false); err != nil {
			log.Println("Error initializing AudioSystem:", err)
			return
		}

		al.PlaySources(ac.player.source)

		if !ac.Background {
			var space *SpaceComponent
			if !entity.Component(&space) {
				return
			}

			ac.player.source.SetPosition(al.Vector{
				(space.Position.X + space.Width/2) / Width(),
				(space.Position.Y + space.Height/2) / Height(),
				0})
		}
	}
}
