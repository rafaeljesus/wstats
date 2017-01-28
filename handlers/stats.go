package handlers

import (
	log "github.com/Sirupsen/logrus"
	"regexp"
	"strings"
)

func (e *Env) StatsCreate(payload string) {
	r, _ := regexp.Compile("\\W")
	words := strings.Fields(r.ReplaceAllString(payload, " "))

	for _, w := range words {
		w = strings.ToLower(w)
		_, err := e.Repo.IncWmap(w, 1)
		if err != nil {
			log.WithError(err).Fatal("[Handlers] Failed to create stats")
		}

		log.WithField("word", w).Info("[Handlers] Word stored")
	}

	log.WithFields(log.Fields{
		"text":        payload,
		"store_count": e.Repo.Count(),
	}).Info("[Handlers] Words successfully proccessed")
}
