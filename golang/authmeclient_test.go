package authmeclient

import "testing"

func TestHashCalculation(t *testing.T) {

  ac := AuthmeClient{
    apiKey: "test",
    apiSecret: "secret",
  }

  request := make(map[string]interface{})
  request["test"] = "hello"

  calculatedHash := ac.GenerateHash(request)
  t.Logf("Calculated hash: %s", calculatedHash)
  if calculatedHash != "292115573432d504b797a836e3a1936ea5ce9ef61bd90c8096dbd86c449d3d75" {
    t.Fail()
    t.Error("Failed, expected hash 292115573432d504b797a836e3a1936ea5ce9ef61bd90c8096dbd86c449d3d75")
  }

}