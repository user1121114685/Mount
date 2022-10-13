// Package all imports all the backends
package all

import (
	// Active file systems
	_ "github.com/rclone/rclone/backend/smb"

	_ "github.com/rclone/rclone/backend/webdav"
)
