package metalcloud

//go:generate go run helper/gen_exports.go

import (
	"fmt"
)

//Infrastructure - the main infrastructure object
type Infrastructure struct {
	InfrastructureID               int                     `json:"infrastructure_id,omitempty" yaml:"id,omitempty"`
	InfrastructureLabel            string                  `json:"infrastructure_label" yaml:"label"`
	DatacenterName                 string                  `json:"datacenter_name" yaml:"datacenter"`
	InfrastructureSubdomain        string                  `json:"infrastructure_subdomain,omitempty" yaml:"subdomain,omitempty"`
	UserIDowner                    int                     `json:"user_id_owner,omitempty" yaml:"ownerID,omitempty"`
	UserEmailOwner                 string                  `json:"user_email_owner,omitempty" yaml:"ownerEmail,omitempty"`
	InfrastructureTouchUnixtime    string                  `json:"infrastructure_touch_unixtime,omitempty" yaml:"touchUnixTime,omitempty"`
	InfrastructureServiceStatus    string                  `json:"infrastructure_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	InfrastructureCreatedTimestamp string                  `json:"infrastructure_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	InfrastructureUpdatedTimestamp string                  `json:"infrastructure_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	InfrastructureChangeID         int                     `json:"infrastructure_change_id,omitempty" yaml:"changeID,omitempty"`
	InfrastructureDeployID         int                     `json:"infrastructure_deploy_id,omitempty" yaml:"deployID,omitempty"`
	InfrastructureDesignIsLocked   bool                    `json:"infrastructure_design_is_locked,omitempty" yaml:"designIsLocked,omitempty"`
	InfrastructureOperation        InfrastructureOperation `json:"infrastructure_operation,omitempty" yaml:"operation,omitempty"`
}

//InfrastructureOperation - object with alternations to be applied
type InfrastructureOperation struct {
	InfrastructureID               int    `json:"infrastructure_id,omitempty" yaml:"id,omitempty"`
	InfrastructureLabel            string `json:"infrastructure_label" yaml:"label"`
	DatacenterName                 string `json:"datacenter_name" yaml:"datacenter"`
	InfrastructureDeployStatus     string `json:"infrastructure_deploy_status,omitempty" yaml:"deployStatus,omitempty"`
	InfrastructureDeployType       string `json:"infrastructure_deploy_type,omitempty" yaml:"deployType,omitempty"`
	InfrastructureSubdomain        string `json:"infrastructure_subdomain,omitempty" yaml:"subdomain,omitempty"`
	UserIDOwner                    int    `json:"user_id_owner,omitempty" yaml:"ownerID,omitempty"`
	InfrastructureUpdatedTimestamp string `json:"infrastructure_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	InfrastructureChangeID         int    `json:"infrastructure_change_id,omitempty" yaml:"changeID,omitempty"`
	InfrastructureDeployID         int    `json:"infrastructure_deploy_id,omitempty" yaml:"deployID,omitempty"`
}

//ShutdownOptions controls how the deploy engine handles running instances
type ShutdownOptions struct {
	HardShutdownAfterTimeout   bool `json:"hard_shutdown_after_timeout"`
	AttemptSoftShutdown        bool `json:"attempt_soft_shutdown"`
	SoftShutdownTimeoutSeconds int  `json:"soft_shutdown_timeout_seconds"`
}

//InfrastructureCreate creates an infrastructure
func (c *Client) InfrastructureCreate(infrastructure Infrastructure) (*Infrastructure, error) {
	var createdObject Infrastructure

	err := c.rpcClient.CallFor(
		&createdObject,
		"infrastructure_create",
		c.user,
		infrastructure)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//infrastructureEdit alters an infrastructure
func (c *Client) infrastructureEdit(infrastructureID id, infrastructureOperation InfrastructureOperation) (*Infrastructure, error) {
	var createdObject Infrastructure

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"infrastructure_edit",
		infrastructureID,
		infrastructureOperation)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//infrastructureDelete deletes an infrastructure and all associated elements. Requires deploy
func (c *Client) infrastructureDelete(infrastructureID id) error {

	if err := checkID(infrastructureID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("infrastructure_delete", infrastructureID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//infrastructureOperationCancel reverts (undos) alterations done before deploy
func (c *Client) infrastructureOperationCancel(infrastructureID id) error {

	if err := checkID(infrastructureID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"infrastructure_operation_cancel",
		infrastructureID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//infrastructureDeploy initiates a deploy operation that will apply all registered changes for the respective infrastructure
func (c *Client) infrastructureDeploy(infrastructureID id, shutdownOptions ShutdownOptions, allowDataLoss bool, skipAnsible bool) error {

	if err := checkID(infrastructureID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"infrastructure_deploy",
		infrastructureID,
		shutdownOptions,
		nil,
		allowDataLoss,
		skipAnsible,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//Infrastructures returns a list of infrastructures
func (c *Client) Infrastructures() (*map[string]Infrastructure, error) {

	res, err := c.rpcClient.Call(
		"infrastructures",
		c.user)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]Infrastructure{}
		return &m, nil
	}

	var createdObject map[string]Infrastructure

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//infrastructureGet returns a specific infrastructure by id
func (c *Client) infrastructureGet(infrastructureID id) (*Infrastructure, error) {
	var infrastructure Infrastructure

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(&infrastructure, "infrastructure_get", infrastructureID)

	if err != nil {
		return nil, err
	}

	return &infrastructure, nil
}

//infrastructureUserLimits returns user metadata
func (c *Client) infrastructureUserLimits(infrastructureID id) (*map[string]interface{}, error) {
	var userLimits map[string]interface{}

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(&userLimits, "infrastructure_user_limits", infrastructureID)

	if err != nil {
		return nil, err
	}

	return &userLimits, nil
}

//CreateOrUpdate implements interface Applier
func (i Infrastructure) CreateOrUpdate(c interface{}) error {
	client := c.(*Client)

	var result *Infrastructure
	var err error

	if i.InfrastructureID != 0 {
		result, err = client.infrastructureGet(i.InfrastructureID)
	} else if i.InfrastructureLabel != "" {
		result, err = client.infrastructureGet(i.InfrastructureLabel)
	} else {
		return fmt.Errorf("id is required")
	}

	if err != nil {
		_, err = client.InfrastructureCreate(i)

		if err != nil {
			return err
		}
	} else {
		i.InfrastructureOperation.InfrastructureChangeID = result.InfrastructureOperation.InfrastructureChangeID
		_, err = client.InfrastructureEdit(i.InfrastructureID, i.InfrastructureOperation)

		if err != nil {
			return err
		}
	}

	return nil
}

//Delete implements interface Applier
func (i Infrastructure) Delete(c interface{}) error {
	client := c.(*Client)

	err := client.infrastructureDelete(i.InfrastructureID)

	if err != nil {
		return err
	}

	return nil
}
