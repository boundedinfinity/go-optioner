package stropt

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOptionals(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Optionals Suite")
}

var _ = Describe("Optional", func() {

})

// func TestSpec(t *testing.T) {
// 	Convey("OptionalString", t, func() {
// 		Convey("IsDefined()", func() {
// 			defined := optional.NewOptionalString(optional.StrPtr("x"))
// 			So(defined.IsDefined(), ShouldBeTrue)

// 			undefined := optional.NewOptionalString(nil)
// 			So(undefined.IsDefined(), ShouldBeFalse)
// 		})

// 		Convey("IsEmpty()", func() {
// 			defined := optional.NewOptionalString(optional.StrPtr("x"))
// 			So(defined.IsEmpty(), ShouldBeFalse)

// 			undefined := optional.NewOptionalString(nil)
// 			So(undefined.IsEmpty(), ShouldBeTrue)
// 		})

// 		Convey("Get()", func() {
// 			input := optional.NewOptionalString(optional.StrPtr("x"))
// 			expected := "x"
// 			So(input.Get(), ShouldEqual, expected)
// 		})

// 		Convey("GetOrElse()", func() {
// 			input := optional.NewOptionalString(nil)
// 			expected := "y"
// 			So(input.GetOrElse("y"), ShouldEqual, expected)
// 		})
// 	})
// }
