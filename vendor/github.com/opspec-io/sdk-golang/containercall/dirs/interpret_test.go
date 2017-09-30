package dirs

import (
	"errors"
	"fmt"
	"github.com/golang-utils/dircopier"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/expression"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/pkg"
	"io/ioutil"
	"path/filepath"
)

var _ = Context("Dirs", func() {
	tempDir, err := ioutil.TempFile("", "")
	if nil != err {
		panic(err)
	}
	Context("Interpret", func() {
		It("should call expression.EvalToDir w/ expected args", func() {
			/* arrange */

			containerDirPath := "/dummyDir1Path.txt"

			providedSCGContainerCallDirs := map[string]string{
				// implicitly bound
				containerDirPath: "",
			}
			providedPkgHandle := new(pkg.FakeHandle)
			providedScope := map[string]*model.Value{}
			providedScratchDir := "dummyScratchDir"

			fakeExpression := new(expression.Fake)
			// error to trigger immediate return
			fakeExpression.EvalToDirReturns(nil, errors.New("dummyError"))

			objectUnderTest := _Dirs{
				expression: fakeExpression,
			}

			/* act */
			objectUnderTest.Interpret(
				providedPkgHandle,
				providedScope,
				providedSCGContainerCallDirs,
				providedScratchDir,
			)

			/* assert */
			actualScope,
				actualExpression,
				actualPkgHandle := fakeExpression.EvalToDirArgsForCall(0)

			Expect(actualScope).To(Equal(providedScope))
			Expect(actualExpression).To(Equal(fmt.Sprintf("$(%v)", containerDirPath)))
			Expect(actualPkgHandle).To(Equal(providedPkgHandle))
		})
		Context("expression.EvalToDir errs", func() {
			It("should return expected error", func() {
				/* arrange */
				containerDirPath := "/dummyDir1Path.txt"
				providedSCGContainerCallDirs := map[string]string{
					// implicitly bound
					containerDirPath: "",
				}

				getContentErr := fmt.Errorf("dummyError")

				fakeExpression := new(expression.Fake)
				fakeExpression.EvalToDirReturns(nil, getContentErr)

				expectedErr := fmt.Errorf(
					"unable to bind %v to %v; error was %v",
					containerDirPath,
					fmt.Sprintf("$(%v)", containerDirPath),
					getContentErr,
				)

				objectUnderTest := _Dirs{
					expression: fakeExpression,
				}

				/* act */
				_, actualErr := objectUnderTest.Interpret(
					new(pkg.FakeHandle),
					map[string]*model.Value{},
					providedSCGContainerCallDirs,
					"dummyScratchDirPath",
				)

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		Context("expression.EvalToDir doesn't err", func() {
			Context("value.Dir not prefixed by rootFSPath", func() {
				It("should return expected results", func() {
					/* arrange */
					containerDirPath := "/dummyDir1Path.txt"

					fakeExpression := new(expression.Fake)
					filePath := tempDir.Name()
					fakeExpression.EvalToDirReturns(&model.Value{Dir: &filePath}, nil)

					expectedDCGContainerCallDirs := map[string]string{
						containerDirPath: filePath,
					}

					objectUnderTest := _Dirs{
						expression: fakeExpression,
						rootFSPath: "dummyRootFSPath",
					}

					/* act */
					actualDCGContainerCallDirs, actualErr := objectUnderTest.Interpret(
						new(pkg.FakeHandle),
						map[string]*model.Value{},
						map[string]string{
							// implicitly bound
							containerDirPath: "",
						},
						"dummyScratchDir",
					)

					/* assert */
					Expect(actualDCGContainerCallDirs).To(Equal(expectedDCGContainerCallDirs))
					Expect(actualErr).To(BeNil())

				})
			})
			Context("value.Dir prefixed by rootFSPath", func() {
				It("should call dircopier.OS w/ expected args", func() {
					/* arrange */
					providedScratchDir := "dummyScratchDir"
					containerDirPath := "/dummyDir1Path.txt"

					fakeExpression := new(expression.Fake)
					filePath := tempDir.Name()
					fakeExpression.EvalToDirReturns(&model.Value{Dir: &filePath}, nil)

					expectedPath := filepath.Join(providedScratchDir, containerDirPath)

					fakeDirCopier := new(dircopier.Fake)

					// err to trigger immediate return
					fakeDirCopier.OSReturns(errors.New("dummyError"))

					objectUnderTest := _Dirs{
						expression: fakeExpression,
						dirCopier:  fakeDirCopier,
					}

					/* act */
					objectUnderTest.Interpret(
						new(pkg.FakeHandle),
						map[string]*model.Value{},
						map[string]string{
							// implicitly bound
							containerDirPath: "",
						},
						providedScratchDir,
					)

					/* assert */
					actualSrcPath,
						actualDstPath := fakeDirCopier.OSArgsForCall(0)

					Expect(actualSrcPath).To(Equal(filePath))
					Expect(actualDstPath).To(Equal(expectedPath))

				})
				Context("dircopier.OS errs", func() {
					It("should return expected error", func() {
						/* arrange */
						containerDirPath := "/dummyDir1Path.txt"

						fakeExpression := new(expression.Fake)
						filePath := tempDir.Name()
						fakeExpression.EvalToDirReturns(&model.Value{Dir: &filePath}, nil)

						fakeDirCopier := new(dircopier.Fake)

						copyError := fmt.Errorf("dummyCopyError")

						// err to trigger immediate return
						fakeDirCopier.OSReturns(copyError)

						expectedErr := fmt.Errorf(
							"unable to bind %v to %v; error was %v",
							containerDirPath,
							fmt.Sprintf("$(%v)", containerDirPath),
							copyError,
						)

						objectUnderTest := _Dirs{
							expression: fakeExpression,
							dirCopier:  fakeDirCopier,
						}

						/* act */
						_, actualErr := objectUnderTest.Interpret(
							new(pkg.FakeHandle),
							map[string]*model.Value{},
							map[string]string{
								// implicitly bound
								containerDirPath: "",
							},
							"dummyScratchDirPath",
						)

						/* assert */
						Expect(actualErr).To(Equal(expectedErr))
					})
				})
			})
		})
	})
})
