#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2023/3/12 16:49
# @File    : aiohttp_example.py
# 异步网络请求模块
import asyncio
import aiohttp


async def get_page_info(url):
    async with aiohttp.ClientSession() as session:
        async with session.get(url) as response:
            # 返回字符串形式数据
            print(await response.text())
            # 返回二进制形式数据
            # response.read()
            # 返回json形式数据
            # response.json()


urls = ['https://www.baidu.com/','https://www.bing.com','http://8.141.175.100/index']
tasks = []
for url in urls:
    t = asyncio.ensure_future(get_page_info(url))
    tasks.append(t)

loop = asyncio.get_event_loop()
loop.run_until_complete(asyncio.wait(tasks))
loop.close()

