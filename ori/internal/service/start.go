package service

import (
	"ori/internal/engine"
	"ori/internal/service/task"
)

func Run(engine *engine.OriEngine) {
	defer engine.Wg.Done()
	engine.Wg.Add(1)
	go new(task.Example).Run(engine)
}
