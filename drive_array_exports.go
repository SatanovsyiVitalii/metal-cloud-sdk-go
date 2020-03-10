// Code generated by gen_exports.go DO NOT EDIT

package metalcloud

//DriveArrayCreate creates a drive array. Requires deploy.
func (c *Client) DriveArrayCreate(infrastructureID int, driveArray DriveArray) (*DriveArray, error) {
	return c.driveArrayCreate(infrastructureID,driveArray)
}

//DriveArrayCreateByLabel creates a drive array. Requires deploy.
func (c *Client) DriveArrayCreateByLabel(infrastructureLabel string, driveArray DriveArray) (*DriveArray, error) {
	return c.driveArrayCreate(infrastructureLabel,driveArray)
}

//DriveArrayEdit alters a deployed drive array. Requires deploy.
func (c *Client) DriveArrayEdit(driveArrayID int, driveArrayOperation DriveArrayOperation) (*DriveArray, error) {
	return c.driveArrayEdit(driveArrayID,driveArrayOperation)
}

//DriveArrayEditByLabel alters a deployed drive array. Requires deploy.
func (c *Client) DriveArrayEditByLabel(driveArrayLabel string, driveArrayOperation DriveArrayOperation) (*DriveArray, error) {
	return c.driveArrayEdit(driveArrayLabel,driveArrayOperation)
}

//DriveArrayDelete deletes a Drive Array with specified id
func (c *Client) DriveArrayDelete(driveArrayID int) error {
	return c.driveArrayDelete(driveArrayID)
}

//DriveArrayDeleteByLabel deletes a Drive Array with specified id
func (c *Client) DriveArrayDeleteByLabel(driveArrayLabel string) error {
	return c.driveArrayDelete(driveArrayLabel)
}

//DriveArrayDrives returns the drives of a drive array
func (c *Client) DriveArrayDrives(driveArray int) (*map[string]Drive, error) {
	return c.driveArrayDrives(driveArray)
}

//DriveArrayDrivesByLabel returns the drives of a drive array
func (c *Client) DriveArrayDrivesByLabel(driveArrLabel string) (*map[string]Drive, error) {
	return c.driveArrayDrives(driveArrLabel)
}
