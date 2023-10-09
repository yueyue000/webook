package tencent

import (
	"context"
	"errors"
	"fmt"
	"github.com/ecodeclub/ekit"
	"github.com/ecodeclub/ekit/slice"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

var ErrNoPhoneNumbers = errors.New("no phone numbers")

type Service struct {
	appID    *string
	SignName *string
	client   *sms.Client
}

func NewService(client *sms.Client, appid string, sigName string) *Service {
	return &Service{
		appID:    ekit.ToPtr[string](appid),
		SignName: ekit.ToPtr[string](sigName),
		client:   client,
	}
}

func (s *Service) Send(ctx context.Context, templateID string, args []string, numbers ...string) error {
	if len(numbers) == 0 {
		return ErrNoPhoneNumbers
	}

	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = s.appID
	req.SignName = s.SignName
	req.TemplateId = ekit.ToPtr[string](templateID)
	req.PhoneNumberSet = s.toStringPtrSlice(numbers)
	req.TemplateParamSet = s.toStringPtrSlice(args)
	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}
	for _, status := range resp.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) != "Ok" {
			return fmt.Errorf("发送短信失败,status:%+v", status)
		}
	}
	return nil
}

func (s *Service) toStringPtrSlice(src []string) []*string {
	return slice.Map[string, *string](src, func(index int, src string) *string {
		return &src
	})
}
