package metalcloud

import (
	"fmt"
	"strings"
)

//OSTemplate A template can be created based on a drive and it has the same characteristics and holds the same information as the parent drive.
type OSTemplate struct {
	VolumeTemplateID                      int                     `json:"volume_template_id,omitempty" yaml:"id,omitempty"`
	VolumeTemplateLabel                   string                  `json:"volume_template_label,omitempty" yaml:"label,omitempty"`
	VolumeTemplateDisplayName             string                  `json:"volume_template_display_name,omitempty" yaml:"name,omitempty"`
	VolumeTemplateSizeMBytes              int                     `json:"volume_template_size_mbytes,omitempty" yaml:"sizeMBytes,omitempty"`
	VolumeTemplateLocalDiskSupported      bool                    `json:"volume_template_local_disk_supported,omitempty" yaml:"localDisk,omitempty"`
	VolumeTemplateIsOSTemplate            bool                    `json:"volume_template_is_os_template,omitempty" yaml:"isOsTemplate,omitempty"`
	VolumeTemplateImageBuildRequired      bool                    `json:"volume_template_image_build_required,omitempty" yaml:"isImageBuildRequired,omitempty"`
	VolumeTemplateProvisionViaOOB         bool                    `json:"volume_template_provision_via_oob,omitempty" yaml:"provisionViaOOB,omitempty"`
	VolumeTemplateBootMethodsSupported    string                  `json:"volume_template_boot_methods_supported,omitempty" yaml:"bootMethods,omitempty"`
	VolumeTemplateOsBootstrapFunctionName string                  `json:"volume_template_os_bootstrap_function_name,omitempty"`
	VolumeTemplateBootType                string                  `json:"volume_template_boot_type,omitempty" yaml:"bootType,omitempty"`
	VolumeTemplateDescription             string                  `json:"volume_template_description,omitempty" yaml:"description,omitempty"`
	VolumeTemplateCreatedTimestamp        string                  `json:"volume_template_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	VolumeTemplateUpdatedTimestamp        string                  `json:"volume_template_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	UserID                                int                     `json:"user_id,omitempty" yaml:"userID,omitempty"`
	VolumeTemplateOperatingSystem         *OperatingSystem        `json:"volume_template_operating_system,omitempty" yaml:"os,omitempty"`
	VolumeTemplateRepoURL                 string                  `json:"volume_template_repo_url,omitempty" yaml:"repoURL,omitempty"`
	VolumeTemplateDeprecationStatus       string                  `json:"volume_template_deprecation_status,omitempty" yaml:"deprecationStatus,omitempty"`
	OSTemplateCredentials                 *OSTemplateCredentials  `json:"os_template_credentials,omitempty" yaml:"credentials,omitempty"`
	VolumeTemplateTags                    []string                `json:"volume_template_tags,omitempty" yaml:"tags,omitempty"`
	OSTemplatePreBootArchitecture         string                  `json:"os_template_pre_boot_architecture,omitempty" yaml:"preBootArchitecture,omitempty"`
	OSAssetBootloaderLocalInstall         int                     `json:"os_asset_id_bootloader_local_install" yaml:"OSAssetIDBootloaderLocalInstall"`
	OSAssetBootloaderOSBoot               int                     `json:"os_asset_id_bootloader_os_boot" yaml:"OSAssetIDBootloaderOSBoot"`
	VolumeTemplateVariablesJSON           string                  `json:"volume_template_variables_json,omitempty" yaml:"variablesJSON,omitempty"`
	VolumeTemplateNetworkOperatingSystem  *NetworkOperatingSystem `json:"volume_template_network_operating_system,omitempty" yaml:"networkOS,omitempty"`
	VolumeTemplateVersion                 string                  `json:"volume_template_version,omitempty"`
	VolumeTemplateOSReadyMethod           string                  `json:"volume_template_os_ready_method,omitempty"`
}

//OSTemplateCredentials holds information needed to connect to an OS installed by an OSTemplate.
type OSTemplateCredentials struct {
	OSTemplateInitialUser                     string `json:"os_template_initial_user,omitempty" yaml:"initialUser,omitempty"`
	OSTemplateInitialPasswordEncrypted        string `json:"os_template_initial_password_encrypted,omitempty" yaml:"initialPasswordEncrypted,omitempty"`
	OSTemplateInitialPassword                 string `json:"os_template_initial_password,omitempty" yaml:"initialPassword,omitempty"`
	OSTemplateInitialSSHPort                  int    `json:"os_template_initial_ssh_port,omitempty" yaml:"initialSSHPort,omitempty"`
	OSTemplateChangePasswordAfterDeploy       bool   `json:"os_template_change_password_after_deploy,omitempty" yaml:"changePasswordAfterDeploy,omitempty"`
	OSTemplateUseAutogeneratedInitialPassword bool   `json:"os_template_use_autogenerated_initial_password,omitempty" yaml:"useAutogeneratedInitialPassword,omitempty"`
}

//OSTemplateOSAssetData holds asset-template information
type OSTemplateOSAssetData struct {
	OSAsset                           *OSAsset `json:"os_asset,omitempty"`
	OSAssetFilePath                   string   `json:"os_asset_file_path,omitempty"`
	OSTemplateOSAssetUpdatedTimestamp string   `json:"volume_template_os_asset_updated_timestamp,omitempty"`
	OSTemplateOSAssetVariablesJSON    string   `json:"volume_template_os_asset_variables_json,omitempty"`
}

//OSTemplateCreate creates a osTemplate object
func (c *Client) OSTemplateCreate(osTemplate OSTemplate) (*OSTemplate, error) {
	var createdObject OSTemplate

	userID := c.GetUserID()

	err := c.rpcClient.CallFor(
		&createdObject,
		"os_template_create",
		userID,
		osTemplate)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//OSTemplateDelete permanently destroys a OSTemplate.
func (c *Client) OSTemplateDelete(osTemplateID int) error {

	if err := checkID(osTemplateID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("os_template_delete", osTemplateID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//OSTemplateUpdate updates a osTemplate
func (c *Client) OSTemplateUpdate(osTemplateID int, osTemplate OSTemplate) (*OSTemplate, error) {
	var createdObject OSTemplate

	if err := checkID(osTemplateID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"os_template_update",
		osTemplateID,
		osTemplate)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//OSTemplateGet returns a OSTemplate specified by nOSTemplateID. The OSTemplate's protected value is never returned.
func (c *Client) OSTemplateGet(osTemplateID int, decryptPasswd bool) (*OSTemplate, error) {

	var createdObject OSTemplate

	if err := checkID(osTemplateID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"os_template_get",
		osTemplateID)

	if err != nil {

		return nil, err
	}

	if decryptPasswd && createdObject.OSTemplateCredentials != nil {
		passwdComponents := strings.Split(createdObject.OSTemplateCredentials.OSTemplateInitialPassword, ":")
		if len(passwdComponents) == 2 {
			if strings.Contains(passwdComponents[0], "Not authorized") {
				return nil, fmt.Errorf("Permission missing. %s", passwdComponents[1])
			} else {
				var passwd string

				err = c.rpcClient.CallFor(
					&passwd,
					"password_decrypt",
					passwdComponents[1],
				)
				if err != nil {
					return nil, err
				}

				createdObject.OSTemplateCredentials.OSTemplateInitialPassword = passwd
			}
		}
	}

	return &createdObject, nil
}

//OSTemplates retrieves a list of all the OSTemplate objects which a specified User is allowed to see through ownership or delegation. The OSTemplate objects never return the actual protected OSTemplate value.
func (c *Client) OSTemplates() (*map[string]OSTemplate, error) {

	userID := c.GetUserID()

	resp, err := c.rpcClient.Call(
		"os_templates",
		userID,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]OSTemplate{}
		return &m, nil
	}
	var createdObject map[string]OSTemplate

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//OSTemplateOSAssets returns the OSAssets assigned to an OSTemplate.
func (c *Client) OSTemplateOSAssets(osTemplateID int) (*map[string]OSTemplateOSAssetData, error) {
	if err := checkID(osTemplateID); err != nil {
		return nil, err
	}

	resp, err := c.rpcClient.Call(
		"os_template_os_assets",
		osTemplateID,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]OSTemplateOSAssetData{}
		return &m, nil
	}

	var createdObject map[string]OSTemplateOSAssetData

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//OSTemplateAddOSAsset adds an asset to a template
func (c *Client) OSTemplateAddOSAsset(osTemplateID int, osAssetID int, path string, variablesJSON string) error {

	// var cond bool
	resp, err := c.rpcClient.Call(
		"os_template_add_os_asset",
		osTemplateID,
		osAssetID,
		path,
		variablesJSON)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//OSTemplateRemoveOSAsset removes an asset from a template
func (c *Client) OSTemplateRemoveOSAsset(osTemplateID int, osAssetID int) error {

	resp, err := c.rpcClient.Call(
		"os_template_remove_os_asset",
		osTemplateID,
		osAssetID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//OSTemplateUpdateOSAssetPath updates an asset mapping
func (c *Client) OSTemplateUpdateOSAssetPath(osTemplateID int, osAssetID int, path string) error {

	resp, err := c.rpcClient.Call(
		"os_template_update_os_asset_path",
		osTemplateID,
		osAssetID,
		path)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//OSTemplateUpdateOSAssetVariables updates an asset variable
func (c *Client) OSTemplateUpdateOSAssetVariables(osTemplateID int, osAssetID int, variablesJSON string) error {

	resp, err := c.rpcClient.Call(
		"os_template_update_os_asset_variables",
		osTemplateID,
		osAssetID,
		variablesJSON)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//OSTemplateMakePublic makes a template public
func (c *Client) OSTemplateMakePublic(osTemplateID int) error {
	resp, err := c.rpcClient.Call(
		"os_template_make_public",
		osTemplateID,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//OSTemplateMakePrivate makes a template private
func (c *Client) OSTemplateMakePrivate(osTemplateID int, userID int) error {
	resp, err := c.rpcClient.Call(
		"os_template_make_private",
		osTemplateID,
		userID,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//CreateOrUpdate implements interface Applier
func (t OSTemplate) CreateOrUpdate(client MetalCloudClient) error {
	var err error
	var result *OSTemplate
	err = t.Validate()

	if err != nil {
		return err
	}

	if t.VolumeTemplateID != 0 {
		result, err = client.OSTemplateGet(t.VolumeTemplateID, false)
	} else {
		templates, err := client.OSTemplates()
		if err != nil {
			return err
		}

		for _, temp := range *templates {
			if temp.VolumeTemplateLabel == t.VolumeTemplateLabel {
				result = &temp
			}
		}
	}

	if err != nil {
		_, err = client.OSTemplateCreate(t)

		if err != nil {
			return err
		}
	} else {
		_, err = client.OSTemplateUpdate(result.VolumeTemplateID, t)

		if err != nil {
			return err
		}
	}

	return nil
}

//Delete implements interface Applier
func (t OSTemplate) Delete(client MetalCloudClient) error {
	var result *OSTemplate
	var id int
	err := t.Validate()

	if err != nil {
		return err
	}

	if t.VolumeTemplateID != 0 {
		id = t.VolumeTemplateID
	} else {
		templates, err := client.OSTemplates()
		if err != nil {
			return err
		}

		for _, temp := range *templates {
			if temp.VolumeTemplateLabel == t.VolumeTemplateLabel {
				result = &temp
			}
		}

		id = result.VolumeTemplateID
	}
	err = client.OSTemplateDelete(id)

	if err != nil {
		return err
	}

	return nil
}

//Validate implements interface Applier
func (t OSTemplate) Validate() error {
	if t.VolumeTemplateID == 0 && t.VolumeTemplateLabel == "" {
		return fmt.Errorf("id is required")
	}

	if t.VolumeTemplateDisplayName == "" {
		return fmt.Errorf("name is required")
	}

	if t.VolumeTemplateBootType == "" {
		return fmt.Errorf("bootType is required")
	}

	if t.VolumeTemplateOperatingSystem.OperatingSystemType == "" {
		return fmt.Errorf("type is required")
	}
	if t.VolumeTemplateOperatingSystem.OperatingSystemVersion == "" {
		return fmt.Errorf("version is required")
	}
	if t.VolumeTemplateOperatingSystem.OperatingSystemArchitecture == "" {
		return fmt.Errorf("architecture is required")
	}

	return nil
}
