package main

import (
	"fmt"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/viper"
	"net"
	"os"
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

var MyAccessKey AccessKey

func main() {
	GetConfig()
	ip2409IPv6s := get2409IPv6s()
	if len(ip2409IPv6s) == 0 {
		fmt.Println("没有获取到2409开头的IPv6地址")
		return
	}
	fmt.Printf("IPv6地址2409开头: %v", ip2409IPv6s)
	err := UpdateIp("blog", "AAAA", ip2409IPv6s, "blog.pangxuejun.cn")
	if err != nil {
		return
	}
}

func get2409IPv6s() (ips []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrs {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
			} else if ipNet.IP.To16() != nil {
				ipV6 := ipNet.IP.To16().String()
				if strings.HasPrefix(ipV6, "2409") {
					ips = append(ips, ipNet.IP.To16().String())
				}
			}
		}
	}
	return
}

func UpdateIp(RR string, Type string, ips []string, subDomain string) (_err error) {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html

	myAccessKey := MyAccessKey.AccessKeyId
	myAccessKeySecret := MyAccessKey.AccessKeySecret

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
	record := options.Body.DomainRecords.Record[0]
	recordId := record.RecordId
	value := record.Value
	fmt.Printf("获取域名recordId:%v，value: %v", *recordId, *value)
	for _, v := range ips {
		if v == *value {
			fmt.Printf("域名解析检测到相同 v:%v，value: %v", v, *value)
			return
		}
	}
	fmt.Printf("域名解析更新，value: %v", ips[0])
	updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		RecordId: tea.String(*recordId),
		RR:       tea.String(RR),
		Type:     tea.String(Type),
		Value:    tea.String(ips[0]),
	}
	updateRuntime := &util.RuntimeOptions{}
	_, _err = client.UpdateDomainRecordWithOptions(updateDomainRecordRequest, updateRuntime)
	if _err != nil {
		return _err
	}
	return _err
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
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

func GetConfig() (accessKeyId string, accessKeySecret string) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config, %s", err))
	}
	err = viper.Unmarshal(&MyAccessKey)
	if err != nil {
		panic(fmt.Errorf("unable to decode into appConf, %v", err))
	}
	return
}
