package constant

import "testing"

func TestMsgAutoDeleteOptionsAreExactlyTheSixteenChoices(t *testing.T) {
	const day = 24 * 60 * 60
	want := []int64{
		1 * day, 2 * day, 3 * day, 4 * day, 5 * day, 6 * day,
		7 * day, 14 * day, 21 * day,
		30 * day, 60 * day, 90 * day, 120 * day, 150 * day, 180 * day,
		365 * day,
	}
	if len(MsgAutoDeleteOptions) != len(want) {
		t.Fatalf("got %d options, want %d", len(MsgAutoDeleteOptions), len(want))
	}
	for i, seconds := range want {
		if MsgAutoDeleteOptions[i] != seconds {
			t.Errorf("option %d = %d, want %d", i, MsgAutoDeleteOptions[i], seconds)
		}
	}
}

func TestIsValidMsgAutoDelete(t *testing.T) {
	for _, seconds := range MsgAutoDeleteOptions {
		if !IsValidMsgAutoDelete(seconds) {
			t.Errorf("%d should be valid", seconds)
		}
	}
	if !IsValidMsgAutoDelete(MsgAutoDeleteOff) {
		t.Error("off should be valid")
	}
	for _, seconds := range []int64{-86400, 1, 60, 120, 3600, 86399, 86401, 8 * 86400, 7 * 86400 * 4, 31 * 86400, 366 * 86400} {
		if IsValidMsgAutoDelete(seconds) {
			t.Errorf("%d should not be valid", seconds)
		}
	}
}
