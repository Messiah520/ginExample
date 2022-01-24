package auth

import (
	"fmt"
	"github.com/casbin/casbin"
)

func Check(e *casbin.Enforcer, sub, obj, act string) {
	ok := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}
