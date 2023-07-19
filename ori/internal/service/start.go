package service

import (
	"ori/core/oriEngine"
	"ori/internal/service/task"
)

func Run(engine *oriEngine.OriEngine) {
	defer engine.Wg.Done()
	engine.Wg.Add(1)
	go new(task.Example).Run(engine)
}
