package main

import (
	"encoding/xml"
	"strings"
	"time"
)

/*
Author: Kaya-Sem

RSS 2.0 SPEC:
https://www.rssboard.org/rss-specification

*/

func NewRSSFeed(channel Channel) *RSS {
	return &RSS{Version: "2.0", Channel: channel}

}

// Creates a new channel with the required title, link and description values. All other values should be supplied with the builder methods
func NewChannel(title, link, description string) *Channel {
	return &Channel{ChannelTitle: title, ChannelLink: link, ChannelDescription: description}
}

func (channel *Channel) PubDate(t time.Time) *Channel {
	channel.ChannelPubDate = time.Time(t).Format(time.RFC1123Z)
	return channel

}

func (channel *Channel) Webmaster(webmaster string) *Channel {
	channel.ChannelWebMaster = webmaster
	return channel

}

func (channel *Channel) SetItems(items []Item) *Channel {
	channel.ChannelItems = items
	return channel

}

func (channel *Channel) AddItem(item Item) *Channel {
	channel.ChannelItems = append(channel.ChannelItems, item)
	return channel

}

func NewItem(title, description string) *Item {
	return &Item{ItemTitle: title, ItemDescription: description}

}

func (item *Item) Link(link string) *Item {
	item.ItemLink = link
	return item

}

func (item *Item) Guid(guid string) *Item {
	item.ItemGuid = guid
	return item

}

func (item *Item) PubDate(t time.Time) *Item {
	item.ItemPubDate = time.Time(t).Format(time.RFC1123Z)
	return item

}

func (item *Item) Source(source string) *Item {
	item.ItemSource = source
	return item

}

func (item *Item) Author(author string) *Item {
	item.ItemAuthor = author
	return item

}

const INDENT = 4

/*
Mashalles the channel and it's items, provides a raw byte slice that can be written out
*/
func (rss *RSS) Marshall() ([]byte, error) {

	xmlData, err := xml.MarshalIndent(rss, "", strings.Repeat(" ", INDENT))
	if err != nil {
		return nil, err
	}

	rssFeed := []byte(xml.Header + string(xmlData))
	return rssFeed, nil

}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

/*
Missing spec implementations:
cloud
ttl
*/

type Channel struct {
	XMLName               xml.Name `xml:"channel"`
	ChannelLastBuildDate  string   `xml:"lastBuildDate,omitempty"`
	ChannelPubDate        string   `xml:"pubDate,omitempty"`
	ChannelTitle          string   `xml:"title"`
	ChannelLink           string   `xml:"link"`
	ChannelDescription    string   `xml:"description"`
	ChannelWebMaster      string   `xml:"webMaster,omitempty"`
	ChannelItems          []Item   `xml:"item"`
	ChannelCopyright      string   `xml:"copyright,omitempty"`
	ChannelManagingEditor string   `xml:"managingEditor,omitempty"`
	ChannelDocs           string   `xml:"docs,omitempty"`

	//ChannelLanguage       string   `xml:"language"`
}

/*
Missing spec implementations:
	- comments
	- catergory
	- enclosure
*/

// Struct representing an item in an RSS feed's channel
type Item struct {
	/* Only one of these is required, but I choose both */
	ItemTitle       string `xml:"title"`
	ItemDescription string `xml:"description"`

	/* Optional */
	ItemLink    string `xml:"link,omitempty"`
	ItemAuthor  string `xml:"author,omitempty"`
	ItemGuid    string `xml:"guid,omitempty"`
	ItemPubDate string `xml:"pubDate,omitempty"`
	ItemSource  string `xml:"source,omitempty"`
}
