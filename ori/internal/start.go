package internal

import (
	"ori/internal/core/oriEngine"
)

func Run(ctx *oriEngine.OriEngine) {
	defer ctx.Wg.Done()
}
