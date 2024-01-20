package game

import (
	"context"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/qiwikgames/babos2d/pkg/engine/gameapp/gamecontext"
	"github.com/qiwikgames/babos2d/pkg/engine/gameapp/scene"
	"log"
)

type Game struct {
	currentScene   scene.Scene
	curSceneParams *scene.Params

	context *gamecontext.Context
	cancel  context.CancelFunc
}

func NewGame(
	sc scene.Scene,
	ctx *gamecontext.Context,
) *Game {

	return &Game{
		currentScene:   sc,
		context:        ctx,
		curSceneParams: &scene.Params{},
	}
}

func (g *Game) Update() error {
	if ok := g.curSceneParams.Get(g.context); ok {
		g.currentScene.OnDestroy(g.context)
		g.currentScene = g.curSceneParams.Scene
		g.currentScene.OnLoad(g.context)
		g.curSceneParams.Unset(g.context)
	}

	return g.currentScene.Update(g.context)
}

func (g *Game) Draw(screen *ebiten.Image) {
	if err := g.currentScene.Draw(screen); err != nil {
		log.Printf("current scene draw: %v", err)
		return
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.context.LogicalWidth, g.context.LogicalHeight
}
