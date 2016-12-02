package lib_test

import (
	"net/url"

	. "github.com/webercoder/go-mtg-hipchat-bot/lib"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeckbrewService", func() {
	Context("NewDeckbrewService", func() {
		It("should create a retriever with a correct URL", func() {
			defaultURL := "https://api.deckbrew.com/mtg/cards"
			dbsvc := NewDeckbrewService()
			Expect(dbsvc).To(BeAssignableToTypeOf(&DeckbrewService{}))
			Expect(dbsvc.URL).To(Equal(defaultURL))
		})
	})

	Context("GetCardsByQuery (Functional Tests)", func() {
		var (
			dbsvc *DeckbrewService
		)
		BeforeEach(func() {
			dbsvc = NewDeckbrewService()
		})
		It("should return the requested cards with no special characters in the name", func() {
			cards, err := dbsvc.GetCardsByQuery("?name=panharmonicon")
			Expect(err).ToNot(HaveOccurred())
			Expect(cards).To(HaveLen(1))
			Expect(cards[0].Name).To(ContainSubstring("Panharmonicon"))
			Expect(cards[0].Types).To(ContainElement("artifact"))
		})
		It("should return the requested cards with a comma in the name", func() {
			cards, err := dbsvc.GetCardsByQuery("?name=" + url.QueryEscape("Selvala,"))
			Expect(err).ToNot(HaveOccurred())
			Expect(cards).To(HaveLen(2))
			Expect(cards[0].Name).To(ContainSubstring("Explorer"))
			Expect(cards[1].Name).To(ContainSubstring("Heart"))
			Expect(cards[0].Types).To(ContainElement("creature"))
			Expect(cards[1].Types).To(ContainElement("creature"))
		})
	})

	Context("GetCardsByQueries (Functional Tests)", func() {
		var (
			dbsvc *DeckbrewService
		)
		BeforeEach(func() {
			dbsvc = NewDeckbrewService()
		})
		It("should return the requested cards", func() {
			queries := []string{"?name=" + url.QueryEscape("selvala, explorer"), "?name=" + url.QueryEscape("selvala,")}
			cards := dbsvc.GetCardsByQueries(queries)
			Expect(cards[0][0].Name).To(ContainSubstring("Explorer"))
			Expect(cards[0][0].Types).To(ContainElement("creature"))
			Expect(cards[1][0].Name).To(ContainSubstring("Explorer"))
			Expect(cards[1][1].Name).To(ContainSubstring("Heart"))
			Expect(cards[1][0].Types).To(ContainElement("creature"))
			Expect(cards[1][1].Types).To(ContainElement("creature"))
		})
	})
})
