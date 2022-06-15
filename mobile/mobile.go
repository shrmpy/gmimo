package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"

	"github.com/shrmpy/gmimo/egj"
)

func init() {
	egjgame, err := egj.NewGame()
	if err != nil {
		panic(err)
	}
	mobile.SetGame(egjgame)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
