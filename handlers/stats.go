package handlers

import (
	log "github.com/Sirupsen/logrus"
)

func (e *Env) StatsCreate(text string) {
	_, err := e.Repo.IncWmap(text, 1)
	if err != nil {
		log.WithError(err).Fatal("[Handlers] Failed to create stats")
	}

	log.WithFields(log.Fields{
		"text":        text,
		"store_count": e.Repo.Count(),
	}).Info("[Handlers] Successfully create stats")
}
