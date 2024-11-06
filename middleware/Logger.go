package middleware

import (
	"GinVueA/define"
	"GinVueA/models"
	"GinVueA/service"
	"GinVueA/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// LoggerToDb 保存日志到数据表
func LoggerToDb() gin.HandlerFunc {
	return func(c *gin.Context) {
		writer := responseWriter{
			c.Writer,
			bytes.NewBuffer([]byte{}),
		}

		c.Writer = writer
		// 开始时间
		startTime := time.Now()
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		var body []byte
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				fmt.Println(err)
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}

		// 获取客户端浏览器和操作系统类型
		os, browser := getOsAndBrowserInfo(c)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 请求地址
		requestURL := c.Request.URL.Path
		// 排除日志管理和图片请求
		if requestURL == "/log" || strings.Index(requestURL, "/uploadFile") > -1 {
			return
		}

		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := utils.GetClientIP(c)
		ipInfo := GetIp2Region(clientIP)
		err := models.DB.Create(&models.SysLog{
			Browser:     os + "_" + browser,         // 浏览器和操作系统
			RequestUri:  requestURL,                 // 请求地址
			ClassMethod: reqUri,                     // 请求路由地址
			HttpMethod:  reqMethod,                  // 请求方式
			UseTime:     latencyTime.Milliseconds(), //执行时间（毫秒）
			RemoteAddr:  clientIP,                   // ip地址
			StatusCode:  statusCode,                 // 状态码
			Response:    writer.b.String(),          // 返回结果
			Params:      string(body),               // 请求参数
			Country:     ipInfo.Country,             // 国家
			Region:      ipInfo.Region,              //区域
			Province:    ipInfo.Province,            // 省份
			City:        ipInfo.City,                // 城市
			Isp:         ipInfo.Isp,                 // 运营商
		}).Error

		if err != nil {
			fmt.Println("保存日志失败！", err)
		}

	}
}

// 自定义结构体，实现gin.ResponseWriter接口
type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

// 获取操作系统和浏览器
func getOsAndBrowserInfo(c *gin.Context) (os string, browser string) {
	var (
		os1      string // 操作系统
		browser1 string // 浏览器
	)

	// 获取用户代理信息
	userAgent := c.Request.Header.Get("User-Agent")
	user := strings.ToLower(userAgent)
	// 判断操作系统类型
	if strings.Contains(user, "windows") {
		os1 = "Windows"
	} else if strings.Contains(user, "mac") {
		os1 = "Mac"
	} else if strings.Contains(user, "x11") {
		os1 = "Unix"
	} else if strings.Contains(user, "android") {
		os1 = "Android"
	} else if strings.Contains(user, "iphone") {
		os1 = "Iphone"
	} else {
		os1 = "UnKnown, More-Info: " + userAgent
	}
	// 判断用户使用的浏览器
	// 通过正则表达式匹配用户代理字符串，以确定浏览器的类型
	if strings.Contains(userAgent, "MSIE") {
		browser1 = "Internet Explorer"
	} else if strings.Contains(userAgent, "Firefox") {
		browser1 = "Firefox"

	} else if strings.Contains(userAgent, "Chrome") {
		browser1 = "Chrome"
	} else {
		browser1 = "Unknown browser"
	}
	return os1, browser1
}

// 重写Write方法
func (w responseWriter) Write(b []byte) (int, error) {
	// 向一个byte.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	// 完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

// GetIp2Region 根据IP地址获取归属地
func GetIp2Region(ip string) *service.IpInfo {
	ipInfo := new(service.IpInfo)
	var dbPath = define.DbPath
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		fmt.Printf("错误信息如下：%s\n\n", err.Error())
		return ipInfo
	}
	defer searcher.Close()

	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("错误信息如下(%s)：%s\n", ip, err)
		return ipInfo
	}

	ipRegion := strings.Split(region, "|")
	newStrList := make([]string, 5)
	newStrList = ipRegion
	// 国家
	ipInfo.Country = newStrList[0]
	// 区域
	ipInfo.Region = newStrList[1]
	// 省份
	ipInfo.Province = newStrList[2]
	// 城市
	ipInfo.City = newStrList[3]
	// 运营商
	ipInfo.Isp = newStrList[4]
	fmt.Printf("{region: %s,took: %s }\n", newStrList, time.Since(tStart))
	return ipInfo

}
