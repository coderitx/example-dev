#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2023/3/20 20:45
# @File    : 烯牛数据.py


import requests
import json
import execjs

headers = {
    "authority": "www.xiniudata.com",
    "accept": "application/json",
    "accept-language": "zh-CN,zh;q=0.9",
    "content-type": "application/json",
    "origin": "https://www.xiniudata.com",
    "referer": "https://www.xiniudata.com/industry/newest",
    "sec-ch-ua": "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Google Chrome\";v=\"110\"",
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": "\"macOS\"",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
}

# 通过扣 javascript代码破解参数
js = execjs.compile(open("烯牛数据.js").read())
js_data = js.call('result')
# print(js_data)

url = "https://www.xiniudata.com/api2/service/x_service/person_industry_list/list_industries_by_sort"

data = {
    "payload": js_data['payload'],
    "sig": js_data['sig'],
    "v": 1
}
data = json.dumps(data, separators=(',', ':'))
response = requests.post(url, headers=headers, data=data).json()

if response['code'] != 0:
    print("请求错误")
    print(response)
else:
    res = js.call('result_decode',response['d'])
    for item in json.loads(res)['list']:
        print(item)

