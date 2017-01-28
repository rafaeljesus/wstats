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
