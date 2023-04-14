#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2023/3/12 16:15
# @File    : async_example.py


import asyncio
import time

async def spider(url):
    print(f"start request {url} ...")
    # 阻塞操作必须手动挂起才会执行异步
    # time.sleep(1)
    await asyncio.sleep(2)
    print(f"request {url} succes...")
    return url


def callback_func(task):
    # 任务对象函数的返回值对象
    print(task.result())


# 返回一个协程对象
c = spider("www.taobao.com")

event_loop = asyncio.get_event_loop()

# 注册协程对象到loop中并且开始执行
# event_loop.run_until_complete(c)
# event_loop.close()

# task 基于event_loop
# t = event_loop.create_task(c)
# print(t) # 没有执行的状态
# event_loop.run_until_complete(t)
# print(t) # 执行完成

# future 基于asyncio创建
# t = asyncio.ensure_future(c)
# print(t)
# event_loop.run_until_complete(t)
# print(t)


# 绑定回调函数
# t = asyncio.ensure_future(c)
# # 绑定回调函数
# t.add_done_callback(callback_func)
# event_loop.run_until_complete(t)
# event_loop.close()


# -----------封装多任务对象-------

urls = ['www.sougou.com','www.baidu.com','www.zhihu.com']
tasks = []
for url in urls:
    t = asyncio.ensure_future(spider(url))
    tasks.append(t)

loop = asyncio.get_event_loop()
start_time = time.time()
loop.run_until_complete(asyncio.wait(tasks))
end_time = time.time()
print("耗时： ",end_time - start_time)



