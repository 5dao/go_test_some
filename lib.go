package some

import "fmt"

type Some struct {
	Age int
}

func (p *Some) Say() {
	fmt.Printf("say: tag +2 %d\r", p.Age)
}
