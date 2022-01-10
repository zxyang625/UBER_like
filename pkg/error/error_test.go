package error

import (
	"fmt"
	"testing"
)

func TestErrorf(t *testing.T) {
	err1 := New(DiscoverClientFail, "cannot find ti")
	err2 := Errorf(DiscoverInstanceNotFound, "func1 failed, err:%v\n", err1)
	fmt.Println(err2)
	fmt.Printf("func2 failed, err: %v\n", New(DiscoverRegisterFail, "aaaaaaaaaaaaa"))
	serviceName := "1241324"
	fmt.Println(Errorf(DiscoverInstanceNotFound, "service %s has no instance", serviceName))
}
