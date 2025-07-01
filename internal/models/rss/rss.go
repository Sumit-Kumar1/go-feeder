package rss

import (
	"encoding/xml"
	"time"
)

// RSS Really Simple Syndication
type RSS struct {
	xml.Name
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title          string     `xml:"title"`
	Description    string     `xml:"description"`
	Link           string     `xml:"link"`
	Image          *Image     `xml:"image,omitempty"`
	LastBuildDate  *time.Time `xml:"lastBuildDate,omitempty"`
	Language       *string    `xml:"language,omitempty"`
	Copyright      *string    `xml:"copyright,omitempty"`
	ManagingEditor *string    `xml:"managingEditor,omitempty"`
	WebMaster      *string    `xml:"webMaster,omitempty"`
	PubDate        *time.Time `xml:"pubDate,omitempty"` // RFC 822
	Docs           *string    `xml:"docs,omitempty"`
	TTL            *int       `xml:"ttl,omitempty"`
	Category       *string    `xml:"category,omitempty"`
	Rating         any        `xml:"rating,omitempty"`
	TextInput      TextInput  `xml:"textInput,omitempty"`
	// A hint for aggregators telling them which hours/days they can skip
	SkipHours any    `xml:"skipHours,omitempty"`
	SkipDays  any    `xml:"skipDays,omitempty"`
	Items     []Item `xml:"item,omitempty"`
}

type Item struct {
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	GUID        *string  `xml:"guid,omitempty"`
	Author      *string  `xml:"author,omitempty"`
	Category    []string `xml:"category,omitempty"`
	Comments    *string  `xml:"comments,omitempty"`
	// Describes a media obj that is attached to item
	Enclosure Enclosure  `xml:"enclosure,omitempty"`
	PubDate   *time.Time `xml:"pubDate,omitempty"`
	Source    *string    `xml:"source,omitempty"`
}

type Enclosure struct {
	URL      *string `xml:"url,omitempty"`
	Length   int     `xml:"length,omitempty"`
	MIMEType string  `xml:"mimetype,omitempty"`
}

// The purpose of the <textInput> element is something of a mystery.
// You can use it to specify a search engine box. Or to allow a reader to provide feedback. Most aggregators ignore it.
type TextInput struct {
	Title       *string `xml:"title,omitempty"`
	Description *string `xml:"description,omitempty"`
	Name        *string `xml:"name,omitempty"`
	Link        *string `xml:"link,omitempty"`
}

type Image struct {
	URL    *string `xml:"url"`
	Title  *string `xml:"title"`
	Link   *string `xml:"link"`
	Width  *int    `xml:"width"`  // max value 144, default 88
	Height *int    `xml:"height"` // max : 400, default : 31
}
