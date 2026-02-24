package test

import (
	"blog/controllers"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var HOST = "http://localhost:9999"

// go test ./test/ -run TestHealth -v
func TestHealth(t *testing.T) {
	Get("/health")
	Post("/health", nil)
}

// go test ./test/ -run TestPostRegister -v
func TestPostRegister(t *testing.T) {
	path := "/api/v1/auth/register"
	// 1. 创建结构体实例
	register := controllers.RegisterRequest{
		Username: "testuser1",
		Email:    "testuser1@example.com",
		Password: "password123",
	}

	Post(path, register)
}

// go test ./test/ -run TestPostLogin -v
func TestPostLogin(t *testing.T) {
	path := "/api/v1/auth/login"
	// 1. 创建结构体实例
	login := controllers.LoginRequest{
		Username: "testuser1",
		Password: "password123",
	}

	Post(path, login)
}

// go test ./test/ -run TestPostProfile -v
func TestPostProfile(t *testing.T) {
	path := "/api/v1/profile"
	Post1(path)
}

func Get(path string) {
	// 1. 定义API地址
	url := HOST + path

	// 2. 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("GET请求失败: %v", err)
	}
	// 重要: 函数执行完毕后关闭响应体，防止资源泄漏
	defer resp.Body.Close()

	// 3. 读取服务器返回的响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应体失败: %v", err)
	}

	// 4. 打印结果
	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应内容:", string(body))
	fmt.Println("API地址:", url)
}

func Post1(path string) {
	// 3. 定义API地址
	url := HOST + path

	// 4. 发送POST请求
	// 参数: URL, Content-Type, 请求体(需要包装为io.Reader)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Fatalf("POST请求失败: %v", err)
	}
	// 重要: 函数执行完毕后关闭响应体，防止资源泄漏
	defer resp.Body.Close()

	// 5. 读取服务器返回的响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应体失败: %v", err)
	}

	// 6. 打印结果
	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应内容:", string(body))
	fmt.Println("API地址:", url)
}

func Post(path string, js interface{}) {
	// 2. 将结构体编码为JSON格式的字节切片 ([]byte)
	jsonData, err := json.Marshal(js)
	if err != nil {
		log.Fatalf("JSON编码失败: %v", err)
	}

	// 3. 定义API地址
	url := HOST + path

	// 4. 发送POST请求
	// 参数: URL, Content-Type, 请求体(需要包装为io.Reader)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("POST请求失败: %v", err)
	}
	// 重要: 函数执行完毕后关闭响应体，防止资源泄漏
	defer resp.Body.Close()

	// 5. 读取服务器返回的响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应体失败: %v", err)
	}

	// 6. 打印结果
	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应内容:", string(body))
	fmt.Println("API地址:", url)
}
