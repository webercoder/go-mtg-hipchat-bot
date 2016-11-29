package lib_test

import (
	. "github.com/webercoder/go-mtg-service/lib"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MTGRetriever", func() {
	Context("NewMTGRetriever", func() {
		It("should create a retriever with a correct URL", func() {
			fakeURL := "http://www.example.com"
			mtgr := NewMTGRetriever(fakeURL)
			Expect(mtgr).To(BeAssignableToTypeOf(&MTGRetriever{}))
			Expect(mtgr.URL).To(Equal(fakeURL))
		})
	})

	Context("GetCardsByName", func() {
		var (
			mtgr *MTGRetriever
		)
		BeforeEach(func() {
			mtgr = NewMTGRetriever("https://api.deckbrew.com/mtg/cards")
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
