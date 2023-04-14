# -*- coding: utf-8 -*-
# @File    : pkg_ddddocr.py
# @desc    : 基于ddddocr库的验证码识别

from PIL import Image
import ddddocr
import os


class DdddocrHandler:

    def __init__(self):
        self.ORR = ddddocr.DdddOcr(show_ad=False)

    def image_code(self, image_path):
        img = Image.open(image_path)
        return self.ORR.classification(img)

# test
def ddddocr_test():
    dir_list = os.listdir("./验证码")
    OCR = ddddocr.DdddOcr(show_ad=False)
    for dir in dir_list:
        if not dir.endswith(".jpg"):
            continue
        img = Image.open(f'./验证码/{dir}')
        img_code = OCR.classification(img)
        print(f"{dir}: {img_code}")


if __name__ == "__main__":
    ddddocr_test()
