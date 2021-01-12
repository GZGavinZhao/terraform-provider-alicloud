package alicloud

import (
	"fmt"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
)

type Brain_industrialService struct {
	client *connectivity.AliyunClient
}

func (s *Brain_industrialService) DescribeBrainIndustrialPidOrganization(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	conn, err := s.client.NewAistudioClient()
	if err != nil {
		return nil, WrapError(err)
	}
	action := "ListPidOrganizations"
	request := map[string]interface{}{
		"RegionId": s.client.RegionId,
	}
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-09-20"), StringPointer("AK"), nil, request, &runtime)
	if err != nil {
		err = WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
		return
	}
	addDebug(action, response, request)
	if fmt.Sprintf(`%v`, response["Code"]) != "200" {
		err = Error("ListPidOrganizations failed for " + response["Message"].(string))
		return object, err
	}
	v, err := jsonpath.Get("$.OrganizationList", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.OrganizationList", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(Error(GetNotFoundMessage("BrainIndustrial", id)), NotFoundWithResponse, response)
	}
	for _, v := range v.([]interface{}) {
		if v.(map[string]interface{})["OrganizationId"].(string) == id {
			return v.(map[string]interface{}), nil
		}
	}
	return object, nil
}