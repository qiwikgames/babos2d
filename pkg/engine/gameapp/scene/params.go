package scene

import (
	"github.com/qiwikgames/babos2d/pkg/engine/gameapp/gamecontext"
)

const ChangeSceneKey = "change_scene"

type Params struct {
	Scene Scene
}

func (p *Params) Get(ctx *gamecontext.Context) (ok bool) {
	if v := ctx.Get(ChangeSceneKey); v != nil {
		if casted, ok := v.(Scene); ok {
			p.Scene = casted
			return ok
		}
	}
	return false
}

func (p *Params) Set(ctx *gamecontext.Context) {
	ctx.Put(ChangeSceneKey, p.Scene)
}

func (p *Params) Unset(ctx *gamecontext.Context) {
	ctx.Unset(ChangeSceneKey)
}
