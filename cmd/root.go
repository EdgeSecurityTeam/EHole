/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ehole",
	Short: "EHole是一款对资产中重点系统指纹识别的工具",
	Long: "\n     ______    __         ______                 \n" +
		"    / ____/___/ /___ ____/_  __/__  ____ _____ ___ \n" +
		"   / __/ / __  / __ `/ _ \\/ / / _ \\/ __ `/ __ `__ \\\n" +
		"  / /___/ /_/ / /_/ /  __/ / /  __/ /_/ / / / / / /\n" +
		" /_____/\\__,_/\\__, /\\___/_/  \\___/\\__,_/_/ /_/ /_/ \n" +
		"			 /____/ https://forum.ywhack.com  By:shihuang\n\n" +
		"    EHole是一款对资产中重点系统指纹识别的工具，在红队作战中，信息收集\n是必不可少的环节，如何才能从大量的资产中提取有用的系统(如OA、VPN、Web\nlogic...)。EHole旨在帮助红队人员在信息收集期间能够快速从C段、大量杂乱\n的资产中精准定位到易被攻击的系统，从而实施进一步攻击。",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ehole.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ehole" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ehole")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
