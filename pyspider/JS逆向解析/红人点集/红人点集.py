# -*- coding: utf-8 -*-
# @Time    : 2023/3/20 14:03
# @File    : 红人点集.py
# @desc    :
import hashlib
import time
import requests
import json


class SpiderHH1024HandlerByJs:
    def __init__(self):
        self._login_url = "https://user.hrdjyun.com/wechat/phonePwdLogin"
        self._username = "*******"
        self._password = "*******"
        self._headers = {
            "Accept": "application/json, text/plain, */*",
            "Accept-Language": "zh-CN,zh;q=0.9",
            "Connection": "keep-alive",
            "Content-Type": "application/json",
            "Origin": "http://www.hh1024.com",
            "Sec-Fetch-Dest": "empty",
            "Sec-Fetch-Mode": "cors",
            "Sec-Fetch-Site": "cross-site",
            "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
            "sec-ch-ua": "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Google Chrome\";v=\"110\"",
            "sec-ch-ua-mobile": "?0",
            "sec-ch-ua-platform": "\"macOS\""
        }

    def encryption_by_md5(self, text):
        return hashlib.md5(text.encode('utf-8')).hexdigest()

    def encryption_by_sha256(self, text):
        return hashlib.sha256(text.encode('utf-8')).hexdigest()

    def get_token(self):
        """
        分析网站登陆流程以及username和password加密方法进行模拟登陆获取token
        :return: token
        """
        token = ""
        pwd = self.encryption_by_md5(self._password)
        username = self._username
        t = int(time.time() * 1000)
        tenant = "1"
        const_str = "JzyqgcoojMiQNuQoTlbR5EBT8TsqzJ"
        sig = self.encryption_by_md5(username + pwd + str(t) + tenant + const_str)
        data = {
            "phoneNum": username,
            "pwd": pwd,
            "t": t,
            "tenant": 1,
            "sig": sig
        }
        data = json.dumps(data, separators=(',', ':'))
        response = requests.post(self._login_url, headers=self._headers, data=data).json()
        if response['status'] == 0:
            token = response['data']['token']
        return token

    def get_data(self, no="", search_data=None):
        """
        更加获取到的token进行登陆获取数据
        :param no:
        :param search_data:
        :return:
        """
        if search_data is None:
            search_data = {}
        token = self.get_token()
        param = {"no": no, "data": search_data}
        C = "kbn%&)@<?FGkfs8sdf4Vg1*+;`kf5ndl$"
        param = json.dumps(param, separators=(',', ':'))
        sign = "param=" + param + "&timestamp=" + str(int(time.time() * 1000)) + "&tenant=1&salt=" + C
        url = "https://ucp.hrdjyun.com:60359/api/dy"
        data = {
            "param": param.replace('"', '\"'),
            "sign": self.encryption_by_sha256(sign),
            "tenant": "1",
            "timestamp": int(time.time() * 1000),
            "token": token,
        }

        data = json.dumps(data, separators=(',', ':'))
        response = requests.post(url, headers=self._headers, data=data).json()
        if response.get('status') != 0:
            print("获取状态异常", response)
            return
        response_data = response.get('data')
        print(response_data)
        for single_data in response_data:
            print(single_data)


if __name__ == '__main__':
    l = SpiderHH1024HandlerByJs()
    l.get_data()
