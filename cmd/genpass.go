/*
Copyright © 2022 Atom Pi <coder.atompi@gmail.com>

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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"gitee.com/autom-studio/webdavd/internal/utils"
)

// genpassCmd represents the genpass command
var genpassCmd = &cobra.Command{
	Use:   "genpass",
	Short: "Generate a bcrypted password hash string.",
	Long:  `Generate a bcrypted password hash string.`,
	Run: func(cmd *cobra.Command, args []string) {
		rawPassword := viper.GetString("password")
		hashedPassword, err := utils.PasswordHasher(rawPassword)
		if err != nil {
			fmt.Fprintln(os.Stderr, "generate bcrypted password failed:", err)
			os.Exit(1)
		} else {
			fmt.Fprintln(os.Stdout, hashedPassword)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(genpassCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genpassCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genpassCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	genpassCmd.PersistentFlags().StringP("password", "p", "", "raw password string")
	viper.BindPFlag("password", genpassCmd.PersistentFlags().Lookup("password"))
}
