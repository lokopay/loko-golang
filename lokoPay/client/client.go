package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"loko-golang/lokoPay/utils"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	publishableKey string
	secretKey      string
	baseUrl        string
	httpClient     *http.Client
}

func NewClient(publishableKey, secretKey, baseUrl string) *Client {
	return &Client{
		publishableKey: publishableKey,
		secretKey:      secretKey,
		baseUrl:        baseUrl,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) SetBaseUrl(baseUrl string) {
	c.baseUrl = baseUrl
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}

func (c *Client) setHeader(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("loko-publishable-key", c.publishableKey)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := uuid.New().String()
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	message := req.URL.String() + string(bodyBytes) + nonce + timestamp
	signature := utils.GenerateHMAC(message, c.secretKey)
	req.Header.Add("loko-timestamp", timestamp)
	req.Header.Add("loko-nonce", nonce)
	req.Header.Add("loko-signature", signature)
}

func (c *Client) Request(method string, path string, params interface{}) (string, error) {
	var paramsStr string
	switch params.(type) {
	case string:
		paramsStr = params.(string)
	case nil:
		paramsStr = ""
	default:
		paramsBytes, _ := json.Marshal(params)
		paramsStr = string(paramsBytes)
	}
	req, err := http.NewRequest(method, c.baseUrl+path, strings.NewReader(paramsStr))
	if err != nil {
		fmt.Println("http.NewRequest error:", err)
		return "", err
	}
	c.setHeader(req)
	response, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	// 检查响应状态码
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s", string(body))
	}
	return string(body), nil
}

func (c *Client) Post(path string, params interface{}) (string, error) {
	return c.Request(http.MethodPost, path, params)
}

func (c *Client) Get(path string, params interface{}) (string, error) {
	return c.Request(http.MethodGet, path, params)
}

// Encrypt 加密结构体中的加密字段
func (c *Client) Encrypt(obj interface{}) {
	// 获取对象的值和类型
	v := reflect.ValueOf(obj)

	// 如果是指针，获取它指向的元素
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 如果不是结构体，直接返回
	if v.Kind() != reflect.Struct {
		return
	}

	t := v.Type()

	// 遍历结构体字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// 检查字段是否有 `encryptAbleField` 标签
		if encryptAbleField, ok := fieldType.Tag.Lookup("encryptAbleField"); ok && encryptAbleField == "true" {
			// 仅加密字符串类型字段
			if field.Kind() == reflect.String {
				encryptedValue := utils.AesEncrypt(field.String(), c.secretKey)
				field.SetString(encryptedValue)
			}
		}

		// 如果字段是结构体或指针，需要递归解密
		if field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct) {
			c.Encrypt(field.Interface())
		}

		// 如果字段是切片，需要检查其元素类型
		if field.Kind() == reflect.Slice {
			for j := 0; j < field.Len(); j++ {
				elem := field.Index(j)
				if elem.Kind() == reflect.Ptr && elem.Elem().Kind() == reflect.Struct {
					c.Encrypt(elem.Interface())
				} else if elem.Kind() == reflect.Struct {
					c.Encrypt(elem.Addr().Interface())
				}
			}
		}
	}
}

// Decrypt 解密结构体中的加密字段
func (c *Client) Decrypt(obj interface{}) {
	// 获取对象的值和类型
	v := reflect.ValueOf(obj)

	// 如果是指针，获取它指向的元素
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 如果不是结构体，直接返回
	if v.Kind() != reflect.Struct {
		return
	}

	t := v.Type()

	// 遍历结构体字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// 检查字段是否有 `encryptAbleField` 标签
		if encryptAbleField, ok := fieldType.Tag.Lookup("encryptAbleField"); ok && encryptAbleField == "true" {
			// 仅解密字符串类型字段
			if field.Kind() == reflect.String {
				decryptedValue := utils.AesDecrypt(field.String(), c.secretKey)
				field.SetString(decryptedValue)
			}
		}

		// 如果字段是结构体或指针，需要递归解密
		if field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct) {
			c.Decrypt(field.Interface())
		}

		// 如果字段是切片，需要检查其元素类型
		if field.Kind() == reflect.Slice {
			for j := 0; j < field.Len(); j++ {
				elem := field.Index(j)
				if elem.Kind() == reflect.Ptr && elem.Elem().Kind() == reflect.Struct {
					c.Decrypt(elem.Interface())
				} else if elem.Kind() == reflect.Struct {
					c.Decrypt(elem.Addr().Interface())
				}
			}
		}
	}
}

func (c *Client) VerifySignature(url, body, lokoSignature string) bool {
	var nonce, timestamp, signature string
	re := regexp.MustCompile(`n=([^;]+);t=([^;]+);s=([^;]+)`)
	matches := re.FindStringSubmatch(lokoSignature)
	if len(matches) == 4 {
		nonce = matches[1]
		timestamp = matches[2]
		signature = matches[3]
	} else {
		return false
	}
	message := url + body + nonce + timestamp
	calSignature := utils.GenerateHMAC(message, c.secretKey)
	return calSignature == signature
}
