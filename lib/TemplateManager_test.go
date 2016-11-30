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
			card := &DeckbrewServiceResponseItem{
				Name:       "Fake Card",
				Cost:       "3BRG",
				Supertypes: []string{"Legendary", "Global"},
				Types:      []string{"Artifact", "Creature"},
				Subtypes:   []string{"Elf", "Goblin"},
				Text:       "This is the card text.",
			}
			err := tm.Execute("card.html", card, &output)
			Expect(err).ToNot(HaveOccurred())
			Expect(output.String()).To(ContainSubstring(card.Name))
			Expect(output.String()).To(ContainSubstring(card.Cost))
			Expect(output.String()).To(ContainSubstring("Legendary Global Artifact Creature - Elf Goblin"))
			Expect(output.String()).To(ContainSubstring(card.Text))
		})
	})
})
