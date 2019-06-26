package nsEshop

import "testing"

func TestClient_GetGamesJapan(t *testing.T) {
	c := NewClient(nil)
	res, err := c.GetGamesJapan()
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}

func TestClient_GetGamesEurope(t *testing.T) {
	c := NewClient(nil)
	res, err := c.GetGamesEurope(3)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}

func TestClient_GetPrices(t *testing.T) {
	c := NewClient(nil)
	res, err := c.GetPrices("US", []string{"70010000002344", "70010000000185", "70010000020291"})
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}
