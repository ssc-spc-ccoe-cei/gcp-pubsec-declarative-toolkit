// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package cmd

import (
	"fmt"

	"arete/internal/cmdsolution"
	
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// solutionCmd represents the create command
var solutionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Solutions",
	Run: func(cmd *cobra.Command, args []string) {
		whiteBold := color.New(color.FgWhite).Add(color.Bold).SprintFunc()

		sl := cmdsolution.SolutionsList{}

		err := sl.GetSolutions()

		if err != nil {
			return
		}
		 for _, sl := range sl.Solutions {
		 	fmt.Printf("%s \n  %s\n", whiteBold(sl.Solution), sl.Description)
		 }
	},
}

// init the command and add flags
func init() {
	solutionCmd.AddCommand(solutionListCmd)
}