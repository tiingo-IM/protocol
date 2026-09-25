package third

import (
	"errors"

	"github.com/openimsdk/protocol/constant"
)

func (x *FcmUpdateTokenReq) Check() error {
	if x.PlatformID > constant.BotPlatformID || x.PlatformID < constant.IOSPlatformID {
		return errors.New("platformID is invalidate")
	}
	if x.FcmToken == "" {
		return errors.New("FcmToken is empty")
	}
	if x.Account == "" {
		return errors.New("account is empty")
	}
	return nil
}

// PushVendors are the push channels a device can register. fcm is
// Google's; the rest are the phone makers' own services.
var PushVendors = map[string]struct{}{
	"fcm": {}, "hms": {}, "honor": {}, "xiaomi": {}, "oppo": {}, "vivo": {},
}

func (x *UpdatePushTokenReq) Check() error {
	if x.PlatformID > constant.BotPlatformID || x.PlatformID < constant.IOSPlatformID {
		return errors.New("platformID is invalidate")
	}
	if _, ok := PushVendors[x.Vendor]; !ok {
		return errors.New("vendor is invalid")
	}
	if x.Token == "" {
		return errors.New("token is empty")
	}
	if x.Account == "" {
		return errors.New("account is empty")
	}
	return nil
}

func (x *SetAppBadgeReq) Check() error {
	if x.UserID == "" {
		return errors.New("UserID is empty")
	}
	return nil
}

func (x *InitiateMultipartUploadReq) Check() error {
	if x.UrlPrefix == "" {
		return errors.New("UrlPrefix is empty")
	}
	return nil
}

func (x *CompleteMultipartUploadReq) Check() error {
	if x.UrlPrefix == "" {
		return errors.New("UrlPrefix is empty")
	}
	return nil
}

func (x *CompleteFormDataReq) Check() error {
	if x.UrlPrefix == "" {
		return errors.New("UrlPrefix is empty")
	}
	return nil
}

func (x *DeleteOutdatedDataReq) Check() error {
	if x.Limit <= 0 {
		return errors.New("limit must be greater than 0")
	}
	if len(x.ObjectGroup) == 0 {
		return errors.New("ObjectGroup is empty")
	}
	return nil
}
