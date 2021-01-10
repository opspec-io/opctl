package node

import (
	"context"

	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
	coreFakes "github.com/opctl/opctl/sdks/go/node/core/fakes"
)

var _ = Context("_node", func() {
	Context("TryResolve", func() {
		It("should call apiClient.ListDescendants w/ expected args", func() {
			/* arrange */
			providedDataRef := "dummyDataRef"

			fakeCore := new(coreFakes.FakeCore)

			providedPullCreds := &model.Creds{
				Username: "dummyUsername",
				Password: "dummyPassword",
			}

			objectUnderTest := _node{
				core:      fakeCore,
				pullCreds: providedPullCreds,
			}

			/* act */
			objectUnderTest.TryResolve(
				context.Background(),
				providedDataRef,
			)

			/* assert */
			actualContext,
				actualReq := fakeCore.ListDescendantsArgsForCall(0)

			Expect(actualContext).To(Equal(context.TODO()))
			Expect(actualReq).To(Equal(model.ListDescendantsReq{
				PkgRef:    providedDataRef,
				PullCreds: providedPullCreds,
			}))
		})
		Context("apiClient.ListDirEntryd errs", func() {
			It("should return expected result", func() {
				/* arrange */
				fakeCore := new(coreFakes.FakeCore)

				listDirEntrysErr := errors.New("dummyError")
				fakeCore.ListDescendantsReturns(nil, listDirEntrysErr)

				objectUnderTest := _node{
					core: fakeCore,
				}

				/* act */
				_, actualErr := objectUnderTest.TryResolve(
					context.Background(),
					"dummyDataRef",
				)

				/* assert */
				Expect(actualErr).To(Equal(listDirEntrysErr))
			})
		})
		Context("apiClient.ListDescendants doesn't err", func() {
			It("should return expected result", func() {
				/* arrange */
				providedDataRef := "dummyDataRef"

				fakeCore := new(coreFakes.FakeCore)

				providedPullCreds := &model.Creds{
					Username: "dummyUsername",
					Password: "dummyPassword",
				}

				objectUnderTest := _node{
					core:      fakeCore,
					pullCreds: providedPullCreds,
				}

				/* act */
				actualHandle, actualErr := objectUnderTest.TryResolve(
					context.Background(),
					providedDataRef,
				)

				/* assert */
				Expect(actualHandle).To(Equal(newHandle(fakeCore, providedDataRef, providedPullCreds)))
				Expect(actualErr).To(BeNil())
			})
		})
	})
})