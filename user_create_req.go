package usermodel

import "github.com/prakash-p-3121/errorlib"

type UserCreateReq struct {
	FirstName      *string `json:"first-name"`
	LastName       *string `json:"last-name"`
	EmailID        *string `json:"email-id"`
	CountryCode    *string `json:"country-code"`
	PhoneNumberStr *string `json:"phone-number"`
	Password       *string `json:"password"`
}

func (req *UserCreateReq) Validate() error {
	if req.FirstName == nil {
		return errorlib.NewBadReqError("first-name-null")
	}
	if req.LastName == nil {
		return errorlib.NewBadReqError("last-name-null")
	}
	if req.EmailID == nil {
		return errorlib.NewBadReqError("email-id-null")
	}
	if req.CountryCode == nil {
		return errorlib.NewBadReqError("country-code-null")
	}
	if req.PhoneNumberStr == nil {
		return errorlib.NewBadReqError("phone-number-null")
	}
	if req.Password == nil {
		return errorlib.NewBadReqError("password-null")
	}
	/* Add new validations like email id regex check, country code check etc.. */
	return nil
}
