// This file is automatically generated. Do not modify it manually.

package main

import "github.com/mattermost/mattermost-server/model"

var manifest = &model.Manifest{
	Description:      "This plugin serves as a starting point for writing a Mattermost plugin.",
	Id:               "com.mattermost.plugin-starter-template",
	MinServerVersion: "5.12.0",
	Name:             "Plugin Starter Template",
	Server: &model.ManifestServer{Executables: &model.ManifestExecutables{
		DarwinAmd64:  "server/dist/plugin-darwin-amd64",
		LinuxAmd64:   "server/dist/plugin-linux-amd64",
		WindowsAmd64: "server/dist/plugin-windows-amd64.exe",
	}},
	SettingsSchema: &model.PluginSettingsSchema{Settings: []*model.PluginSetting{}},
	Version:        "0.1.0",
	Webapp:         &model.ManifestWebapp{BundlePath: "webapp/dist/main.js"},
}
