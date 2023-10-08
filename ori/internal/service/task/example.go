package task

import (
	"fmt"
	"ori/internal/engine"
)

type Example struct {
}

func (e *Example) Run(engine *engine.OriEngine) {
	defer engine.Wg.Done()
	//从下面开始写你得常驻任务代码
	fmt.Println("执行example任务")
}
