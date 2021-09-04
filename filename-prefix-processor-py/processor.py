#!/usr/bin/python3

import os
from typing import List


def PathCheck(path:str)->bool:
    '''
    路径检查
    '''
    # 判断给定路径是否存在且是否为文件目录
    return os.path.exists(path) and os.path.isdir(path)


def GetFiles(path:str, prefix:str)->List[str]:
    '''
    根据名称前缀，获取指定路径下的所有直属文件名信息
    '''
    filenames = list()

    # 列出指定路径中所有直属文件信息
    # 根据前缀切分原文件名并生成新文件名
    # 如果想获取指定路径下所有文件、子目录及其子文件信息
    # 使用递归或者更高效的os.walk()实现
    for filename in os.listdir(path):
        if os.path.isfile(
                os.path.join(path,filename)
                ) and filename.startswith(prefix):
            filenames.append(filename)

    return filenames


def Rename(path:str, prefix:str, filenames:List[str]):
    '''
    重命名文件
    '''
    # 获取原文件名
    for filename in filenames:
        # 根据前缀切分原文件名并生成新文件名
        new_filename = filename[len(prefix):]

        old_path = os.path.join(path,filename)
        new_path = os.path.join(path,new_filename)

        os.rename(old_path,new_path)

    # Done!


def main():
    try:
        path = input("请输入待处理文件的所在路径:\n")
        if PathCheck(path) is False:
            print("指定文件路径不存在!")
            return

        prefix = input("请输入待处理文件的名称前缀:\n")
        filenames = GetFiles(path,prefix)
        if len(filenames) == 0:
            print("{0}路径下暂无以'{1}'为前缀的文件!".format(path,prefix))
            return

        print("已为您匹配到以下文件:")
        for filename in filenames:
            print(filename)

        verification = input("确认是否执行操作[y/n]:")
        if verification != "y":
            print("已终止操作!")
            return

        Rename(path,prefix,filenames)
        print("Done!")

    except Exception as e:
        print(e)


if __name__ == "__main__":
    main()

