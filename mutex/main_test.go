package main
import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "hello world"
	wg.Add(2)
	go UpdateMessage("x")
	go UpdateMessage("Goodbye, cruel word!")
	wg.Wait()

	if msg != "Goodbye, cruel word!" {
		t.Error("incorrect value in msg")
	}
}
