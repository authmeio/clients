package authmeclient

import (
  "time"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "bytes"
  "errors"
  "sort"
  "fmt"
  "strings"
  "crypto/sha256"
)

var client = &http.Client{}

type AuthmeClientInterface interface {
  GetOrder(referenceId string) (Order, error)
  InitOrder(AuthenticationRequest) (Order, error)
  InitOrderWithApiKey(orderPlace AuthenticationRequest, apiKey string) (Order, error)
}

type Order struct {
  Id          uint64 `json:"-"`
  CreatedAt   time.Time
  ReferenceId string
  UserId      uint64
  Comment     string
  Status      string
  Details     string
  Token       *string
  Data        *string
}

type AuthmeClient struct {
  endPoint  string
  apiKey    string
  apiSecret string
}

type AuthenticationRequest struct {
  ReferenceId        string
  UserIdentifier     string
  UserIdentifierType string
  PublicKeyJson      string
  Comment            string
  Message            string
  Ip                 string
  Hash               string
  Client             string
  Data               string
}

func NewAuthmeClient(endPoint string) AuthmeClientInterface {
  return &AuthmeClient{
    endPoint: endPoint,
  }
}

func NewAuthmeClientWithApiKey(endPoint string, apiKey string, apiSecret string) AuthmeClientInterface {
  return &AuthmeClient{
    endPoint: endPoint,
    apiKey: apiKey,
    apiSecret: apiSecret,
  }
}

func (x *AuthmeClient) GenerateHash(request map[string]interface{}) string {
  var keys []string
  //log.Infof("%v", x)
  for k := range request {
    if k == "Hash" {
      continue
    }
    keys = append(keys, k)
  }

  //givenHash, ok := request["Hash"]
  //if !ok {
  //    return false
  //}
  sort.Strings(keys)
  values := make([]string, 1)
  values[0] = x.apiKey

  for _, key := range keys {
    var v string
    if request[key] != nil {
      v = fmt.Sprintf("%v", request[key])
    } else {
      v = ""
    }
    if len(v) < 1 {
      //values = append(values, "")
      continue
    }
    //log.Infof("%s: %v", key, v)
    values = append(values, v)
  }
  values = append(values, x.apiSecret)

  hashString := strings.Join(values, "|")
  h := sha256.New()
  h.Write([]byte(hashString))
  bs := h.Sum(nil)
  calculatedHash := fmt.Sprintf("%x", bs)
  return calculatedHash
}

func (ac *AuthmeClient) GetOrder(referenceId  string) (Order, error) {

  url := ac.endPoint + "/order/" + referenceId
  response, err := http.Get(url)
  if err != nil {
    return Order{}, err
  }
  content, err := ioutil.ReadAll(response.Body)
  var order Order
  json.Unmarshal(content, &order)
  return order, err
}

func (ac *AuthmeClient) InitOrder(order AuthenticationRequest) (Order, error) {
  if len(ac.apiKey) < 2 {
    panic(errors.New("ApiKey not set"))
  }
  return ac.InitOrderWithApiKey(order, ac.apiKey)
}

func (ac *AuthmeClient) InitOrderWithApiKey(orderPlace AuthenticationRequest, apiKey string) (Order, error) {
  jsonRequest, _ := json.Marshal(orderPlace)
  req, err := http.NewRequest("POST", ac.endPoint + "/order", bytes.NewBuffer(jsonRequest))
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("X-API-KEY", apiKey)

  resp, err := client.Do(req)

  if err != nil {
    return Order{}, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  var o Order
  json.Unmarshal([]byte(body), &o)
  return o, nil
}























