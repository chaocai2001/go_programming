package panic_recover

import (
	"os"
	"testing"
)

func TestPanicRecover(t *testing.T) {
	t.Log("Start")
	//panic(errors.New("something wrong"))
	os.Exit(-1)
	t.Log("End")
}
