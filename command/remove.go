/* Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

package command

import (
	"flag"

	"fmt"

	api "github.com/elodina/stack-deploy/framework"
)

type RemoveStackCommand struct{}

func (rsc *RemoveStackCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println("Stack name required to remove")
		return 1
	}

	var (
		flags  = flag.NewFlagSet("remove", flag.ExitOnError)
		apiUrl = flags.String("api", "", "Stack-deploy server address.")
		force  = flags.Bool("force", false, "Flag to force delete the stack with all children.")
	)
	flags.Parse(args[1:])

	name := args[0]

	stackDeployApi, err := resolveApi(*apiUrl)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return 1
	}
	client := api.NewClient(stackDeployApi)

	fmt.Printf("Removing stack %s\n", name)
	err = client.RemoveStack(&api.RemoveStackRequest{
		Name:  name,
		Force: *force,
	})
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return 1
	}

	fmt.Println("Stack removed")
	return 0
}

func (rsc *RemoveStackCommand) Help() string {
	return ""
}

func (rsc *RemoveStackCommand) Synopsis() string {
	return "Remove existing stack"
}
