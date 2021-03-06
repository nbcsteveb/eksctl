package cloudformation

import (
	"encoding/json"
)

// AWSRDSOptionGroup_OptionSetting AWS CloudFormation Resource (AWS::RDS::OptionGroup.OptionSetting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations-optionsettings.html
type AWSRDSOptionGroup_OptionSetting struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations-optionsettings.html#cfn-rds-optiongroup-optionconfigurations-optionsettings-name
	Name *Value `json:"Name,omitempty"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations-optionsettings.html#cfn-rds-optiongroup-optionconfigurations-optionsettings-value
	Value *Value `json:"Value,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSOptionGroup_OptionSetting) AWSCloudFormationType() string {
	return "AWS::RDS::OptionGroup.OptionSetting"
}

func (r *AWSRDSOptionGroup_OptionSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
