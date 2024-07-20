package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func main() {
	uuidWithHypen := uuid.New()
	fmt.Println(uuidWithHypen)

	uuid := strings.Replace(uuidWithHypen.String(), "-", "", -1)
	fmt.Println(uuid)

}
