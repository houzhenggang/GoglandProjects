package main

import "time"

type Project struct {
	Name string    `json:"name"`
	StartDate time.Time   `json:"date"`
}

type Projects []Project

