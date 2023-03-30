package task

import (
	"fmt"
	"ori/core/oriEngine"
)

type Example struct {
}

func (e *Example) Run(engine *oriEngine.OriEngine) {
	defer engine.Wg.Done()
	fmt.Println("执行example任务")
}
