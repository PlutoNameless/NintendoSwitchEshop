package NintendoSwitchEshop

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) GetGamesJapan() ([]GameJapan, error) {
	currentGamesRes, err := c.sendNewRequest(JP_GET_GAMES_CURRENT)
	if err != nil {
		return nil, err
	}

	comingGamesRes, err := c.sendNewRequest(JP_GET_GAMES_COMING)
	if err != nil {
		return nil, err
	}

	var currentGames GameJapanXML
	err = xml.Unmarshal(currentGamesRes, &currentGames)
	if err != nil {
		return nil, err
	}

	var comingGames GameJapanXML
	err = xml.Unmarshal(comingGamesRes, &comingGames)
	if err != nil {
		return nil, err
	}

	return append(currentGames.Games, comingGames.Games...), nil
}

func (c *Client) GetGamesEurope(limit int) ([]GameEurope, error) {
	if limit < 1 {
		limit = 1
	}

	if limit > 9999 {
		limit = 9999
	}

	locale := "en"
	u, err := url.Parse(fmt.Sprintf("https://search.nintendo-europe.com/%s/select", locale))
	if err != nil {
		return nil, err
	}

	query := u.Query()

	query.Set("fq", "type:GAME AND system_type:nintendoswitch* AND product_code_txt:*")
	query.Set("q", "*")
	query.Set("rows", strconv.Itoa(limit))
	query.Set("sort", "sorting_title asc")
	query.Set("start", strconv.Itoa(0))
	query.Set("wt", "json")

	u.RawQuery = query.Encode()

	europeGamesRes, err := c.sendNewRequest(u.String())
	if err != nil {
		return nil, err
	}

	var europeGames GamesEuropeJSON
	err = json.Unmarshal(europeGamesRes, &europeGames)
	if err != nil {
		return nil, err
	}

	return europeGames.Response.Games, nil
}

func (c *Client) GetPrices(country string, gameIds []string) (pr PriceResponse, err error) {
	if len(gameIds) == 0 {
		return pr, errors.New("gameIds is nil")
	}

	// TODO: 一次最多50个
	if len(gameIds) > PRICE_LIST_LIMIT {
		panic("todo")
	}

	u, err := url.Parse("https://api.ec.nintendo.com/v1/price")
	if err != nil {
		return pr, err
	}

	query := u.Query()

	query.Set("country", country)
	for _, id := range gameIds {
		query.Add("ids", id)
	}
	query.Set("limit", strconv.Itoa(50))
	query.Set("lang", "en")

	u.RawQuery = query.Encode()

	priceDataRes, err := c.sendNewRequest(u.String())
	if err != nil {
		return pr, err
	}

	err = json.Unmarshal(priceDataRes, &pr)
	if err != nil {
		return pr, err
	}

	if pr.Error.Code != 0 {
		return pr, errors.New(pr.Error.Message)
	}

	return pr, nil
}
