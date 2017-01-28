package handlers

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"regexp"
	"strings"
)

func (e *Env) StatsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	topWords, count := e.Repo.SortedByWords()
	topLetters := e.Repo.SortedByLetters()
	wsize := len(topWords)
	lsize := len(topLetters)
	wlimit := 5
	llimit := 5

	if wsize < 5 {
		wlimit = wsize
	}

	if lsize < 5 {
		llimit = lsize
	}

	response := map[string]interface{}{
		"count":      count,
		"total":      e.Repo.Count(),
		"topWords":   topWords[:wlimit],
		"topLetters": topLetters[:llimit],
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (e *Env) StatsCreate(payload string) {
	r, _ := regexp.Compile("\\W")
	words := strings.Fields(r.ReplaceAllString(payload, " "))

	for _, w := range words {
		w = strings.ToLower(w)
		_, err := e.Repo.IncWmap(w, 1)
		if err != nil {
			log.WithError(err).Fatal("[Handlers] Failed to create word stats")
		}

		log.WithField("word", w).Info("[Handlers] Word stored")

		for _, letter := range w {
			l := strings.ToLower(string(letter))
			_, err := e.Repo.IncLmap(l, 1)
			if err != nil {
				log.WithError(err).Fatal("[Handlers] Failed to create letter stats")
			}

			log.WithField("letter", l).Info("[Handlers] Letter stored")
		}
	}

	log.WithFields(log.Fields{
		"text":        payload,
		"store_count": e.Repo.Count(),
	}).Info("[Handlers] Words successfully proccessed")
}
