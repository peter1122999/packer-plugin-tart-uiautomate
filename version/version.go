package version

import packersdkversion "github.com/hashicorp/packer-plugin-sdk/version"

var GitVersion = "0.0.1"

var PluginVersion = packersdkversion.NewRawVersion(GitVersion)
