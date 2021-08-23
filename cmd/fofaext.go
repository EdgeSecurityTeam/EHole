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
	"ehole/module/finger/source"
	"ehole/module/fofaext"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// fofaextCmd represents the fofaext command
var fofaextCmd = &cobra.Command{
	Use:   "fofaext",
	Short: "ehole的fofa提取模块",
	Long:  `从fofa api提取资产并保存成xlsx，支持大批量ip提取,支持fofa所有语法。`,
	Run: func(cmd *cobra.Command, args []string) {
		file := strings.Split(ext_output, ".")

		if len(file) == 2 {
			if file[1] == "xlsx" {
				if ext_fofaip != "" {
					results := source.Fafaips_out(ext_fofaip)
					fofaext.Fofaext(results, ext_output)
					os.Exit(1)
				}
				if ext_fofasearche != "" {
					fmt.Println(ext_fofasearche)
					results := source.Fofaall_out(ext_fofasearche)
					fofaext.Fofaext(results, ext_output)
					os.Exit(1)
				}
			} else {
				log.Println("文件名错误！！！")
			}
		} else {
			log.Println("文件名错误！！！")
		}
	},
}

var (
	ext_fofaip      string
	ext_fofasearche string
	ext_output      string
)

func init() {
	rootCmd.AddCommand(fofaextCmd)
	fofaextCmd.Flags().StringVarP(&ext_fofaip, "ipfile", "l", "", "从文本获取IP，在fofa搜索，支持大量ip，默认保存所有结果。")
	fofaextCmd.Flags().StringVarP(&ext_fofasearche, "fofa", "s", "", "从fofa提取资产，支持fofa所有语法，默认保存所有结果。")
	fofaextCmd.Flags().StringVarP(&ext_output, "output", "o", "results.xlsx", "指定输出文件名和位置，当前仅支持xlsx后缀的文件。")
}
