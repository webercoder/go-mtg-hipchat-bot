package lib_test

import (
	. "github.com/webercoder/go-mtg-hipchat-bot/lib"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeckbrewService", func() {
	Context("NewDeckbrewService", func() {
		It("should create a retriever with a correct URL", func() {
			defaultURL := "https://api.deckbrew.com/mtg/cards"
			mtgr := NewDeckbrewService()
			Expect(mtgr).To(BeAssignableToTypeOf(&DeckbrewService{}))
			Expect(mtgr.URL).To(Equal(defaultURL))
		})
	})

	Context("GetCardsByName (Functional Tests)", func() {
		var (
			mtgr *DeckbrewService
		)
		BeforeEach(func() {
			mtgr = NewDeckbrewService()
		})
		It("should return the requested cards with no special characters in the name", func() {
			cards, err := mtgr.GetCardsByName("Panharmonicon")
			Expect(err).ToNot(HaveOccurred())
			Expect(cards).To(HaveLen(1))
			Expect(cards[0].Name).To(ContainSubstring("Panharmonicon"))
		})
		It("should return the requested cards with a comma in the name", func() {
			cards, err := mtgr.GetCardsByName("Selvala,")
			Expect(err).ToNot(HaveOccurred())
			Expect(cards).To(HaveLen(2))
			Expect(cards[0].Name).To(ContainSubstring("Explorer"))
			Expect(cards[1].Name).To(ContainSubstring("Heart"))
		})
	})
})
