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
	})

	Context("GetCardsByNames (Functional Tests)", func() {
		var (
			dbsvc *DeckbrewService
		)
		BeforeEach(func() {
			dbsvc = NewDeckbrewService()
		})
		It("should return the requested cards with no special characters in the name", func() {
			cards := dbsvc.GetCardsByNames([]string{"panharmonicon"}, 0)
			Expect(cards).To(HaveLen(1))
			Expect(cards["panharmonicon"][0].Name).To(ContainSubstring("Panharmonicon"))
			Expect(cards["panharmonicon"][0].Types).To(ContainElement("artifact"))
		})
		It("should return the requested cards with a comma in the name", func() {
			cards := dbsvc.GetCardsByNames([]string{"Selvala,"}, 0)
			Expect(cards["Selvala,"]).To(HaveLen(2))
			Expect(cards["Selvala,"][0].Name).To(ContainSubstring("Explorer"))
			Expect(cards["Selvala,"][1].Name).To(ContainSubstring("Heart"))
			Expect(cards["Selvala,"][0].Types).To(ContainElement("creature"))
			Expect(cards["Selvala,"][1].Types).To(ContainElement("creature"))
		})
		It("should return the number of cards in the limit", func() {
			cards := dbsvc.GetCardsByNames([]string{"lightning"}, 2)
			Expect(cards["lightning"]).To(HaveLen(2))
		})
		It("should only return one card when there is an exact match", func() {
			cards := dbsvc.GetCardsByNames([]string{"blizzard"}, 0)
			Expect(cards["blizzard"]).To(HaveLen(1))
		})
		It("should return multiple requested cards", func() {
			queries := []string{"selvala, explorer", "selvala,"}
			cards := dbsvc.GetCardsByNames(queries, 0)
			Expect(cards["selvala, explorer"][0].Name).To(ContainSubstring("Explorer"))
			Expect(cards["selvala, explorer"][0].Types).To(ContainElement("creature"))
			Expect(cards["selvala,"][0].Name).To(ContainSubstring("Explorer"))
			Expect(cards["selvala,"][1].Name).To(ContainSubstring("Heart"))
			Expect(cards["selvala,"][0].Types).To(ContainElement("creature"))
			Expect(cards["selvala,"][1].Types).To(ContainElement("creature"))
		})
		It("should replace tokens with icons", func() {
			cards := dbsvc.GetCardsByNames([]string{"Chromanticore"}, 0)
			Expect(cards["Chromanticore"][0].Cost).To(ContainSubstring("(w)"))
			Expect(cards["Chromanticore"][0].Cost).To(ContainSubstring("(u)"))
			Expect(cards["Chromanticore"][0].Cost).To(ContainSubstring("(b)"))
			Expect(cards["Chromanticore"][0].Cost).To(ContainSubstring("(r)"))
			Expect(cards["Chromanticore"][0].Cost).To(ContainSubstring("(g)"))

			Expect(cards["Chromanticore"][0].Text).To(ContainSubstring("{2}"))
			Expect(cards["Chromanticore"][0].Text).To(ContainSubstring("(w)"))
			Expect(cards["Chromanticore"][0].Text).To(ContainSubstring("(u)"))
			Expect(cards["Chromanticore"][0].Text).To(ContainSubstring("(b)"))
			Expect(cards["Chromanticore"][0].Text).To(ContainSubstring("(r)"))
			Expect(cards["Chromanticore"][0].Text).To(ContainSubstring("(g)"))

			cards = dbsvc.GetCardsByNames([]string{"Sol Ring"}, 0)
			Expect(cards["Sol Ring"][0].Text).To(ContainSubstring("(t)"))
		})
	})
})
