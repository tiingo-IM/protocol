package constant

const (
	SignalingNotificationBegin               = 1600
	SignalingNotification                    = 1601
	RoomParticipantsConnectedNotification    = 1602
	RoomParticipantsDisconnectedNotification = 1603
	StreamChangedNotification                = 1604
	CustomSignalNotification                 = 1605
	// "X started a group call" — sdkws.CallStartedTips. Unlike the ring
	// (1601) this one is stored, with a seq, so that a member who opens
	// the group mid-call finds the call in the transcript and can join
	// it. The call's end is still the ordinary constant.Call record.
	CallStartedNotification = 1606

	SignalingNotificationEnd = 1699
)

const (
	MeetingNotEnd = 0
	MeetingEnd    = 1
)
