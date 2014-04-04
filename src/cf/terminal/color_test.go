package terminal_test

import (
	. "cf/terminal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"runtime"
)

var _ = Describe("Testing with ginkgo", func() {
	Describe("CF_COLOR", func() {
		Context("On OSes that don't support colours", func() {
			BeforeEach(func() { OsSupportsColours = Nope })

			Context("When the CF_COLOR env variable is specified", func() {
				BeforeEach(func() { os.Setenv("CF_COLOR", "true") })
				itDoesntColourize()
			})

			Context("When the CF_COLOR env variable is not specified", func() {
				BeforeEach(func() { os.Setenv("CF_COLOR", "") })
				itDoesntColourize()
			})
		})

		Context("On OSes that support colours", func() {
			BeforeEach(func() { OsSupportsColours = Yep })

			Context("When the CF_COLOR env variable is not specified", func() {
				BeforeEach(func() { os.Setenv("CF_COLOR", "") })

				Context("And the terminal supports colours", func() {
					BeforeEach(func() { TerminalSupportsColours = Yep })
					itColourizes()
				})

				Context("And the terminal doesn't support colours", func() {
					BeforeEach(func() { TerminalSupportsColours = Nope })
					itDoesntColourize()
				})
			})

			Context("When the CF_COLOR env variable is set to 'true'", func() {
				BeforeEach(func() { os.Setenv("CF_COLOR", "true") })

				Context("And the terminal supports colours", func() {
					BeforeEach(func() { TerminalSupportsColours = Yep })
					itColourizes()
				})

				Context("Even if the terminal doesn't support colours", func() {
					BeforeEach(func() { TerminalSupportsColours = Nope })
					itColourizes()
				})
			})

			Context("When the CF_COLOR env variable is set to 'false', even if the terminal supports colours", func() {
				BeforeEach(func() {
					os.Setenv("CF_COLOR", "false")
					TerminalSupportsColours = Yep
				})

				itDoesntColourize()
			})
		})
	})

	Describe("OsSupportsColours", func() {
		It("Returns false on windows, and true otherwise", func() {
			if runtime.GOOS == "windows" {
				Expect(OsSupportsColours()).To(BeFalse())
			} else {
				Expect(OsSupportsColours()).To(BeTrue())
			}
		})
	})

	var (
		originalOsSupportsColours       func() bool
		originalTerminalSupportsColours func() bool
	)

	BeforeEach(func() {
		originalOsSupportsColours = OsSupportsColours
		originalTerminalSupportsColours = TerminalSupportsColours
	})

	AfterEach(func() {
		OsSupportsColours = originalOsSupportsColours
		TerminalSupportsColours = originalTerminalSupportsColours
	})
})

func itColourizes() {
	It("colourizes", func() {
		text := "Hello World"
		colorizedText := Colorize(text, 31, true)
		Expect(colorizedText).To(Equal("\033[1;31mHello World\033[0m"))
	})
}

func itDoesntColourize() {
	It("doesn't colourize", func() {
		text := "Hello World"
		colorizedText := Colorize(text, 31, true)
		Expect(colorizedText).To(Equal("Hello World"))
	})
}

func Yep() bool  { return true }
func Nope() bool { return false }
