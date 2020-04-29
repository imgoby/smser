package api

import "fmt"

type Service struct {

}

func (s *Service)test() {
	fmt.Println("ok")
}
