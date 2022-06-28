package some

import "fmt"

type Some struct {
	Name string
}

func (p *Some) Say() {
	fmt.Printf("say: tag 1 %s\r", p.Name)
}
