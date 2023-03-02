package service

import (
	"ori/app/http"
	"ori/internal/core/oriEngine"
)

func Run(engine *oriEngine.OriEngine) {
	engine.Wg.Add(1)
	go http.Run(engine)
}
