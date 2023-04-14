#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2023/3/19 21:36
# @File    : 企明科技.py

import requests
import execjs


def get_encrypt_data():
    headers = {
        "Accept": "application/json, text/plain, */*",
        "Accept-Language": "zh-CN,zh;q=0.9",
        "Connection": "keep-alive",
        "Content-Type": "application/x-www-form-urlencoded",
        "Origin": "https://www.qimingpian.com",
        "Sec-Fetch-Dest": "empty",
        "Sec-Fetch-Mode": "cors",
        "Sec-Fetch-Site": "cross-site",
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
        "sec-ch-ua": "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Google Chrome\";v=\"110\"",
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-platform": "\"macOS\""
    }
    url = "https://vipapi.qimingpian.cn/DataList/investmentTotalVip"
    data = {
        "time_interval": "",
        "industry": "",
        "page": "1",
        "num": "20",
        "unionid": ""
    }
    response = requests.post(url, headers=headers, data=data).json()
    if response.get('status') == 0:
        return response.get('encrypt_data')
    return ""


def deciphering_by_js():
    encrypt_data = get_encrypt_data()
    js = execjs.compile(open("企明科技.js", "r", encoding="utf-8").read())
    data = js.call("deciphering", encrypt_data)
    return data


if __name__ == '__main__':
    data_list = deciphering_by_js().get('list')
    for item in data_list:
        print(item)
