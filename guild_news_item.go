package wow

import (
	"errors"
	"time"
	"fmt"
)

type GuildNewsItem struct {
	Type        string
	Character   string
	Timestamp   uint64
	ItemId      int
	item        *Item
	Achievement *Achievement
}


func (g *GuildNewsItem) Time() time.Time{
	return time.Unix(int64(g.Timestamp)/1000, int64(g.Timestamp)%1000)
}

func (g *GuildNewsItem) Ago() time.Duration {
	return time.Now().Sub(g.Time())
}

func (g *GuildNewsItem) HowLongAgo() string {
	d := g.Ago()
	if d.Hours() > 2 {
		return fmt.Sprintf("%d hours ago", int(d.Hours()))
	}
	if d.Minutes() > 2 {
		return fmt.Sprintf("%d minutes ago", int(d.Minutes()))
	}
	if d.Seconds() > 1 {
		return fmt.Sprintf("%d seconds ago", int(d.Seconds()))
	}
	return fmt.Sprintf("%d nanoseconds ago", int(d.Nanoseconds()))
}

func (g *GuildNewsItem) Item() (*Item, error) {
	var err error
	if g.item != nil {
		return g.item, nil
	}
	if g.ItemId != 0 {
		client := CurrentApiClient()
		if client != nil {
			g.item, err = client.GetItem(g.ItemId)
			return g.item, err
		} else {
			return nil, errors.New("No API client created. Create one via NewApiClient")
		}
	}
	return nil, errors.New("No ItemId set")
}

