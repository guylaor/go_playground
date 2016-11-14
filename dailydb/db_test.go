package main

import (
  "testing"
  dailydb "github.com/guylaor/dailydb"
)

func TestGetRule(t *testing.T) {
  feed := `[{"AdServer": "usa.ads-daily.net", "GEO": "USA"}, {"AdServer": "fr.ads-daily.net", "GEO": "France"}, {"AdServer": "eu.ads-daily.net", "GEO": "Europe"}]`
  dailydb.LoadRulesFromJson(feed)

  var tests = []struct {
    input, expected string
  }{
    {"Europe", "eu.ads-daily.net"},
    {"USA", "usa.ads-daily.net"},
    {"DUMMY", "default.ads-daily.net"},
  }
  for _, test := range tests {
    res := dailydb.GetRule(test.input)
    if (res != test.expected) {
      t.Errorf("dailydb.GetRule(%s) != %s", test.input, test.expected)
    }
  }
}
