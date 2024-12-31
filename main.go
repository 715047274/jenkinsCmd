/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/715047274/jenkinsCmd/cmd"
	"github.com/spf13/viper"
)

var cfgFile string

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("$HOME")
		viper.SetConfigFile(".yourApp")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Print("using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	cmd.Execute()
}
