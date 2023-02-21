// Copyright (C) 2017 NTT Innovation Institute, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/cloudwan/gohan/extension/goext"
	"github.com/ttsubo2000/goext_example/todo"
)

const (
	entryCustomAction = "entry_custom_action"
)

// Init is an entry point of this extension
func Init(env goext.IEnvironment) error {
	// Find the pre-loaded schema
	schema := env.Schemas().Find("entry")

	// Register runtime type for the schema (raw)
	schema.RegisterRawType(todo.Entry{})

	// Register schema handler
	schema.RegisterResourceEventHandler(goext.PreUpdateTx, func(ctx goext.Context, res goext.Resource, env goext.IEnvironment) *goext.Error {
		env.Logger().Infof("Called resource pre-update-in-transaction handler")

		// Cast resource to its runtime type
		entry := res.(*todo.Entry)

		// Modify resource
		entry.Name = "name changed in pre_update_in_transaction event"
		env.Logger().Infof("Modified Todo resource in pre-update-in-transaction handler: %v", entry)

		return nil
	}, goext.PriorityDefault)

	// Register a custom action handler
	schema.RegisterResourceEventHandler(entryCustomAction, func(ctx goext.Context, res goext.Resource, env goext.IEnvironment) *goext.Error {
		env.Logger().Infof("Called resource custom action handler")

		return nil
	}, goext.PriorityDefault)

	return nil
}
