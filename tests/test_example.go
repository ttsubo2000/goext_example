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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ttsubo2000/goext_example/todo"
)

// Schemas returns a list of extra schemas that will be loaded with this test
func Schemas() []string {
	return []string{
		"../todo/entry.yaml",
		"../todo/link.yaml",
	}
}

// Binaries returns a path to the plugin's binary that this test depends on
func Binaries() []string {
	return []string{"../example.so"}
}

// Test is an entry point of this test
func Test(mockEnv goext.MockIEnvironment) {
	mockEnv.Logger().Notice("Running example test suite")

	Describe("Example tests", func() {
		var (
			schema  goext.ISchema
			entry   *todo.Entry
			err     error
			context goext.Context
			tx      goext.ITransaction
		)

		BeforeEach(func() {
			tx, err = mockEnv.Database().Begin()
			Expect(err).To(BeNil())

			context = goext.MakeContext().WithTransaction(tx)

			schema = mockEnv.Schemas().Find("entry")
			Expect(schema).To(Not(BeNil()))

			entry = &todo.Entry{
				ID:          mockEnv.Util().NewUUID(),
				Name:        "example name",
				Description: "example description",
				TenantID:    "admin",
			}
		})

		AfterEach(func() {
			mockEnv.Reset()
		})

		It("Smoke test CRUD", func() {
			Expect(schema.CreateRaw(entry, context)).To(Succeed())
			Expect(schema.UpdateRaw(entry, context)).To(Succeed())
			Expect(schema.DeleteRaw(entry.ID, context)).To(Succeed())
		})

		It("Should change name in pre_update event", func() {
			Expect(schema.CreateRaw(entry, context)).To(Succeed())
			entry.Name = "random name"
			Expect(schema.UpdateRaw(entry, context)).To(Succeed())
			Expect(entry.Name).To(Equal("name changed in pre_update_in_transaction event"))
		})
	})
}
