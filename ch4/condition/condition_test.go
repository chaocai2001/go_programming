package condition

import (
	"fmt"
	"testing"
)

func TestMultiCaseItems(t *testing.T) {
	for j := 0; j < 5; j++ {
		switch j {
		case 0, 2:
			fmt.Println(j, "even")
		case 1, 3:
			fmt.Println(j, "odd")
		default:
			fmt.Println(j, "unknown")
		}

	}
}

func TestSwitchWithoutCondition(t *testing.T) {
	for j := 0; j < 5; j++ {
		switch {
		case j%2 == 0:
			fmt.Println(j, "even")
		case j%2 == 1:
			fmt.Println(j, "odd")
		default:
			fmt.Println(j, "unknown")
		}

	}
}
