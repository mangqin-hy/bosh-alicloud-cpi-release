/*
 * Copyright (C) 2017-2017 Alibaba Group Holding Limited
 */
package action

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cpi:reboot_vm", func() {
	It("can run reboot vm", func() {
		r :=  caller.RunTest([]byte(`


		`), )
		Expect(r.Error).NotTo(HaveOccurred())
	})
})