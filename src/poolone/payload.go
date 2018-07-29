package poolone

import "fmt"

type Payload struct{
	Name string
}

//@name task process function
func (p *Payload) Process(){
	fmt.Printf("%v 执行任务,任务处理完成！", p.Name)
}
