package constant

import "slices"

// Message auto-delete: a conversation (one-to-one or group) can ask that
// every message sent into it from then on is deleted for everyone a fixed
// time after it was sent. The setting is stamped onto each message as
// sdkws.MsgData.expireAt when it is sent, so changing or turning off the
// setting later never changes a message that was already sent.
//
// Only these durations are accepted, in seconds. A month is 30 days and a
// year 365, because the setting is a length of time, not a calendar date.
// MsgAutoDeleteOff turns the setting off.
const (
	MsgAutoDeleteOff int64 = 0

	msgAutoDeleteDay   int64 = 24 * 60 * 60
	msgAutoDeleteWeek  int64 = 7 * msgAutoDeleteDay
	msgAutoDeleteMonth int64 = 30 * msgAutoDeleteDay
	msgAutoDeleteYear  int64 = 365 * msgAutoDeleteDay
)

// MsgAutoDeleteOptions lists every duration a conversation may choose,
// shortest first. A var only because Go has no constant slices — never
// mutate it.
var MsgAutoDeleteOptions = []int64{
	1 * msgAutoDeleteDay,
	2 * msgAutoDeleteDay,
	3 * msgAutoDeleteDay,
	4 * msgAutoDeleteDay,
	5 * msgAutoDeleteDay,
	6 * msgAutoDeleteDay,
	1 * msgAutoDeleteWeek,
	2 * msgAutoDeleteWeek,
	3 * msgAutoDeleteWeek,
	1 * msgAutoDeleteMonth,
	2 * msgAutoDeleteMonth,
	3 * msgAutoDeleteMonth,
	4 * msgAutoDeleteMonth,
	5 * msgAutoDeleteMonth,
	6 * msgAutoDeleteMonth,
	1 * msgAutoDeleteYear,
}

// IsValidMsgAutoDelete reports whether seconds is off or one of
// MsgAutoDeleteOptions.
func IsValidMsgAutoDelete(seconds int64) bool {
	return seconds == MsgAutoDeleteOff || slices.Contains(MsgAutoDeleteOptions, seconds)
}
