package authmeclient

import (
    "time"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "bytes"
    "errors"
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
    endPoint string
    apiKey   string
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

func NewAuthmeClientWithApiKey(endPoint string, apiKey string) AuthmeClientInterface {
    return &AuthmeClient{
        endPoint: endPoint,
        apiKey: apiKey,
    }
}

func (ac *AuthmeClient) GetOrder(referenceId  string) (Order, error) {
    log.Printf("Authme end point: " + ac.endPoint)

    url := ac.endPoint + "/order/" + referenceId
    log.Printf("Request url: %s", url)
    response, err := http.Get(url)
    if err != nil {
        return Order{}, err
    }
    content, err := ioutil.ReadAll(response.Body)
    log.Printf("Get order response: %s", content)
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























