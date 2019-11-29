package main
import (
	"fmt"
	"github.com/nu7hatch/gouuid"
)

func Uuid() *uuid.UUID{
	u4, err := uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)

	}
	return u4
}
