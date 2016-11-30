package lib_test

import (
	"bytes"

	. "github.com/webercoder/go-mtg-hipchat-bot/lib"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TemplateManager", func() {
	Context("Execute", func() {
		It("should create card html", func() {
			var output bytes.Buffer
			tm := &TemplateManager{}
			card := struct {
				Name     string
				Cost     string
				TypeLine string
				Text     string
			}{
				"Fake Card",
				"3BRG",
				"Legendary Global Artifact Creature - Elf Goblin",
				"This is the card text.",
			}
			err := tm.Execute("card.html", card, &output)
			Expect(err).ToNot(HaveOccurred())
			Expect(output.String()).To(ContainSubstring(card.Name))
			Expect(output.String()).To(ContainSubstring(card.Cost))
			Expect(output.String()).To(ContainSubstring(card.TypeLine))
			Expect(output.String()).To(ContainSubstring(card.Text))
		})
	})
})
