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
package diff

import (
	"github.com/glacion/spotify-utils/pkg/cache"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Source string `mapstructure:"source"`
	Target string `mapstructure:"target"`
	Apply  bool   `mapstructure:"apply"`
}

// Command represents the diff command
var Command = &cobra.Command{
	Use:   "diff",
	Short: "shows and optionally applies the diff between two playlists",
	Run: func(_ *cobra.Command, _ []string) {
		c := Config{}
		viper.Unmarshal(&c)
		log.Printf("%+v", c)
		cache.SetToken("kekw")
	},
}

func init() {
	Command.Flags().StringP("source", "s", "", "source playlist")
	Command.Flags().StringP("target", "t", "", "target playlist")
	Command.Flags().BoolP("apply", "a", false, "apply changes")
	Command.MarkFlagRequired("source")
	Command.MarkFlagRequired("target")
}
