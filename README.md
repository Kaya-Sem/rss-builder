# RSS Builder

A Go package for building RSS 2.0 feeds with a fluent builder pattern. This package provides a simple and intuitive way to create RSS feeds that comply with the [RSS 2.0 specification](https://www.rssboard.org/rss-specification).

## Features

- Fluent builder pattern for easy RSS feed creation
- RSS 2.0 specification compliant
- Support for all major RSS elements
- XML marshaling with proper formatting
- Type-safe API

## Installation

```bash
go get github.com/Kaya-Sem/rss-builder
```

## Usage

Here's a basic example of how to create an RSS feed:

```go
package main

import (
    "time"
    "os"
)

func main() {
    // Create a new channel with required fields
    channel := NewChannel(
        "My RSS Feed",
        "https://example.com",
        "This is my RSS feed description",
    )

    // Add optional channel information
    channel.PubDate(time.Now())
    channel.Webmaster("webmaster@example.com")

    // Create items
    item1 := NewItem(
        "First Article",
        "This is the first article",
    ).Link("https://example.com/article1").
        Guid("article1").
        PubDate(time.Now()).
        Author("John Doe")

    item2 := NewItem(
        "Second Article",
        "This is the second article",
    ).Link("https://example.com/article2").
        Guid("article2").
        PubDate(time.Now()).
        Author("Jane Smith")

    // Add items to channel
    channel.AddItem(item1)
    channel.AddItem(item2)

    // Create RSS feed
    rss := NewRSSFeed(*channel)

    // Marshal to XML
    xmlData, err := rss.Marshall()
    if err != nil {
        panic(err)
    }

    // Write to file
    os.WriteFile("feed.xml", xmlData, 0644)
}
```

## API Reference

### Channel Creation

```go
// Create a new channel with required fields
channel := NewChannel(title, link, description)

// Add optional channel information
channel.PubDate(time.Time)      // Set publication date
channel.Webmaster(string)       // Set webmaster email
channel.SetItems([]Item)        // Set all items at once
channel.AddItem(Item)           // Add a single item
```

### Item Creation

```go
// Create a new item with required fields
item := NewItem(title, description)

// Add optional item information
item.Link(string)              // Set item link
item.Guid(string)              // Set item GUID
item.PubDate(time.Time)        // Set publication date
item.Source(string)            // Set source
item.Author(string)            // Set author
```

### RSS Feed Creation

```go
// Create a new RSS feed from a channel
rss := NewRSSFeed(channel)

// Marshal to XML
xmlData, err := rss.Marshall()
```

## Supported RSS Elements

### Channel Elements
- title (required)
- link (required)
- description (required)
- pubDate
- webMaster
- copyright
- managingEditor
- docs
- lastBuildDate

### Item Elements
- title (required)
- description (required)
- link
- author
- guid
- pubDate
- source

## Notes

Some RSS 2.0 specification elements are not yet implemented:
- cloud
- ttl
- comments
- category
- enclosure

