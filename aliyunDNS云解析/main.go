package main

import (
	"flag"
	"fmt"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

/*
* author: longjunzhi
* date:
* description: 监听自身ip变化，进行域名dns解析。
* 监听ipV4、ipV6、dns解析
 */

type AccessKey struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
}

type AliyunCloudResolution struct {
	AccessKey AccessKey `json:"access_key"`
}

var MyAliyunCloudResolution AliyunCloudResolution

func main() {
	GetConfig()

	// 定义命令行参数
	var host string
	var ip string

	flag.StringVar(&host, "host", "", "The full domain name (e.g., dfs.sd.pangxuejun.cn).")
	flag.StringVar(&ip, "ip", "", "The IP addresses to update, separated by commas (e.g., 234.343.34,123.45.67.89).")

	// 解析命令行参数
	flag.Parse()

	// 检查必需的参数是否已提供
	if host == "" || ip == "" {
		fmt.Println("Usage: go run main.go -host <host> -type <type> -ip <ip>")
		os.Exit(1)
	}
	ipType := "A"
	// 如果IP字符串ipv6 则ipType等于"AAAA"
	if strings.Contains(ip, ":") {
		ipType = "AAAA"
	}
	// 提取一级域名前面的部分
	parts := strings.Split(host, ".")
	if len(parts) < 3 {
		fmt.Printf("Invalid domain: %s\n", host)
		return
	}
	subDomain := strings.Join(parts[:2], ".")
	fmt.Println("subDomain: %s, ipType: %s, ip: %s, host: %s", subDomain, ipType, []string{ip}, host)
	err := UpdateIp(subDomain, ipType, ip, host)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UpdateIp(RR string, Type string, ip string, subDomain string) (_err error) {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	myAccessKey := MyAliyunCloudResolution.AccessKey.AccessKeyId
	myAccessKeySecret := MyAliyunCloudResolution.AccessKey.AccessKeySecret

	client, _err := CreateClient(tea.String(myAccessKey), tea.String(myAccessKeySecret))
	if _err != nil {
		return _err
	}
	// 获取解析记录详情
	describeSubDomainRecordsRequest := &alidns20150109.DescribeSubDomainRecordsRequest{
		SubDomain: tea.String(subDomain),
	}
	runtime := &util.RuntimeOptions{}
	options, err := client.DescribeSubDomainRecordsWithOptions(describeSubDomainRecordsRequest, runtime)
	if err != nil {
		return err
	}
	if len(options.Body.DomainRecords.Record) == 0 {
		fmt.Println("未查询到域名解析记录")
		return
	}
	record := options.Body.DomainRecords.Record[0]
	recordId := record.RecordId
	value := record.Value
	fmt.Printf("获取域名recordId:%v，value: %v", *recordId, *value)
	if ip == *value {
		fmt.Printf("域名解析检测到相同 v:%v，value: %v", ip, *value)
		return
	}
	fmt.Printf("域名解析更新，value: %v", ip)
	updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		RecordId: tea.String(*recordId),
		RR:       tea.String(RR),
		Type:     tea.String(Type),
		Value:    tea.String(ip),
	}
	updateRuntime := &util.RuntimeOptions{}
	_, _err = client.UpdateDomainRecordWithOptions(updateDomainRecordRequest, updateRuntime)
	if _err != nil {
		return _err
	}
	return _err
}

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *alidns20150109.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("alidns.cn-beijing.aliyuncs.com")
	_result = &alidns20150109.Client{}
	_result, _err = alidns20150109.NewClient(config)
	return _result, _err
}

func GetConfig() {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(filepath.Dir(path))
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config, %s", err))
	}
	err = viper.Unmarshal(&MyAliyunCloudResolution)
	if err != nil {
		panic(fmt.Errorf("unable to decode into appConf, %v", err))
	}
	return
}
