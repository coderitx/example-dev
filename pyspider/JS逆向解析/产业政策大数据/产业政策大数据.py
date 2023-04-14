#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2023/3/22 21:24
# @File    : 产业政策大数据.py


import requests
import execjs

headers = {
    "Accept": "application/json, text/plain, */*",
    "Accept-Language": "zh-CN,zh;q=0.9",
    "Cache-Control": "no-cache",
    "Connection": "keep-alive",
    "Content-Type": "application/octet-stream",
    "Origin": "http://www.spolicy.com",
    "Pragma": "no-cache",
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
}
cookies = {
    "JSESSIONID": "A5E3826DECD171D3E4C3AF42C9C5CF7F"
}
url = "http://www.spolicy.com/info_api/policyType/showPolicyType"
params = {
    "policyType": "6",
    "province": "",
    "city": "",
    "downtown": "",
    "garden": "",
    "centralId": "",
    "sort": 0,
    "homePageFlag": 1,
    "pageNum": 1,
    "pageSize": 7
}

data = execjs.compile(open("产业政策大数据.js",encoding='utf-8').read()).call('Start',params)
# print(data)

response = requests.post(url, headers=headers, cookies=cookies, data=bytes(data['data']), verify=False)

print(response.text)
print(response)