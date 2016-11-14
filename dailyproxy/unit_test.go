package dailyproxy_test

import (
  "testing"
  "github.com/guylaor/dailyproxy"
)

func TestFindGeoByIpAddress(t *testing.T) {
  res := dailyproxy.FindGeoByIpAddress("127.0.0.1")
  if (res != "Local") {
    t.Error("findGeoByIpAddress('127.0.0.1') is not Local")
  }
}

func TestMultipleIpLocations(t *testing.T) {
  var tests = []struct {
      ip string
      expected string
    }{
    {"127.0.0.1", "Local"},
    {"127.5.5.52", "Local"},
    {"28.1.1.5", "France"},
    {"28.255.21.2", "France"},
    {"25.2.2.1", "Europe"},
    {"200.0.0.1", "USA"},
  }
  for _, test := range tests {
    if res := dailyproxy.FindGeoByIpAddress(test.ip); res != test.expected {
      t.Errorf("dailyproxy.FindGeoByIpAddress(%s) != %s", test.ip, test.expected)
    }
  }
}
