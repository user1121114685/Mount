// Package all imports all the commands
package all

import (
	// Active commands
	_ "github.com/rclone/rclone/cmd"
	_ "github.com/rclone/rclone/cmd/cmount"
	_ "github.com/rclone/rclone/cmd/config"
	_ "github.com/rclone/rclone/cmd/mount"
	_ "github.com/rclone/rclone/cmd/mount2"
)
