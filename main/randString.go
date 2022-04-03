// get a random string

package main

import (
	"fmt"

	"github.com/polarbroadband/goto/util"
)

func randString() {
	fmt.Println(util.RandString(30))
}
