/*
Copyright © 2023 Can Güvendiren

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/glacion/spotify-utils/cmd/diff"
	"github.com/glacion/spotify-utils/pkg/cache"
	"github.com/glacion/spotify-utils/pkg/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	ApplicationName = "spotify-utils"
	Version         = "0.0.0"
	ConfigFilename  = "config.json"
	CacheFilename   = "cache.json"
)

var (
	configPath string
	cachePath  string
	verbose    bool
)

var Command = &cobra.Command{
	Use:     ApplicationName,
	Short:   "utilities for spotify",
	Version: Version,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		config.ConfigureLogger(verbose)

		err := config.Configure(cmd, configPath)
		log.Fatal().Err(err)

		err = cache.Init(cachePath)
		log.Fatal().Err(err)
	},
}

func init() {
	Command.AddCommand(diff.Command)

	configDir, err := os.UserConfigDir()
	log.Fatal().Err(err)
	defaultConfig := filepath.Join(configDir, ApplicationName, ConfigFilename)

	cacheDir, err := os.UserCacheDir()
	log.Fatal().Err(err)
	defaultCache := filepath.Join(cacheDir, ApplicationName, CacheFilename)

	Command.PersistentFlags().StringVar(
		&cachePath,
		"cache",
		defaultCache,
		"cache file",
	)
	Command.PersistentFlags().StringVar(
		&configPath,
		"config",
		defaultConfig,
		"config file",
	)
	Command.PersistentFlags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"verbose output",
	)
}
