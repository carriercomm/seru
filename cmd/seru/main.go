/*
** Copyright [2012-2013] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */
package main

import (
	"github.com/indykish/seru/cmd"
	"os"
)

const (
	version = "0.1.0"
	header  = "Supported-Seru"
)

func buildManager(name string) *cmd.Manager {
	m := cmd.BuildBaseManager(name, version, header)
	m.Register(&Listdomain{}) //route53 list zones
	m.Register(&NewSubdomain{}) //route53 create CNAME
	m.Register(&DeleteSubdomain{})  //route53 delete CNAME
	return m
}

func main() {
	name := cmd.ExtractProgramName(os.Args[0])
	manager := buildManager(name)
	manager.Run(os.Args[1:])
}
