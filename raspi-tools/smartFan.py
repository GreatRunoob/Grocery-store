# -*- coding=utf-8 -*-
import time
import RPi.GPIO as gpio

# monitor模块是我个人编写用于检测树莓派各项硬件信息的模块
# monitor模块包含一个用于检测硬件信息的PiMonitor类
# 内含多项通用静态类用于获取树莓派常用硬件信息
# 使用本脚本之前，请确保monitor模块被包含在同级目录下
from monitor import PiMonitor

# date模块时我个人编写用于获取不同格式的本地时间信息的模块
# date模块尚处于开发过程中，目前仅包含Date静态类
# 使用本脚本之前，请确保date模块被包含在同级目录下
from date import Date


class SmartFan:
    # 启用单例模式
    _instance=None
    _BOARDPIN=(3,5,7,8,10,11,12,13,15,16,18,19,21,22,23,24,26,)

    def __init__(self,pin=8):
        # 设定默认使用的gpio针脚为8号物理针脚
        if self.setGPIOPin(pin) is False:
            self._pin=8
        # CPU内核达到设定的最高温度后开启风扇散热
        self._htemp=50
        # CPU内核降至设定的最低温度后关闭散热风扇
        self._ltemp=35
        # 脚本监测CPU内核温度的时间监测，单位秒
        self._interval=30
        # 设定默认风扇状态为关闭状态
        self._fanoff=True
        # 设定gpio针脚编程模式
        gpio.setmode(gpio.BOARD)
        # 忽略因自定义针脚编程引起的警告信息
        gpio.setwarnings(False)
        # 初始化默认针脚为输出模式且默认为低电平
        gpio.setup(self._pin,gpio.OUT,initial=gpio.LOW)


    def __new__(cls,*args,**kw):
        if cls._instance is None:
            cls._instance=object.__new__(cls,*args,**kw)
        return cls._instance


    def fanControl(self):
        try:
            print("Pin {0} is currently in use.".format(self._pin))
            while True:
                # 获取当前CPU内核温度
                temp=PiMonitor.getCPUTemp()
                ftime=Date.getCurrentFormatDatetime()
                # 显示当前CPU内核温度
                print("{0} Current temperature of the CPU is {1:.2f}'C".format(ftime,temp))
                # 当内核温度达到50'C及以上时，启动风扇散热
                if temp>=self._htemp and self._fanoff is True:
                    gpio.output(self._pin,gpio.HIGH)
                    self._fanoff=False
                    print("The cooling fan is activated and running...")
                # 当内核温度降至35'C及以下时，关闭散热风扇
                elif temp<=self._ltemp and self._fanoff is False:
                    gpio.output(self._pin,gpio.LOW)
                    self._fanoff=True
                    print("The cooling fan is deactivated.")

                # 每隔30s监测决定风扇状态
                print("{0} Check after {1}s...".format(ftime,self._interval))
                time.sleep(self._interval)
        except KeyboardInterrupt:
            print("Keyboard interrupt!")
            print("User has terminated the script!")
        except Exception as error:
            print("Unknown exception!")
            print(error)
        finally:
            # 退出脚本前，清理并释放本次占用的gpio资源
            gpio.cleanup()
            print("Clean up successfully!")


    def setGPIOPin(self,pin:int)->bool:
        if pin not in self._BOARDPIN:
            return False
        else:
            self._pin=pin
            return True


    def setHighTemp(self,htemp:int)->bool:
        if htemp<=self._ltemp:
            return False
        else:
            self._htemp=htemp
            return True


    def setLowTemp(self,ltemp:int)->bool:
        if ltemp>=self._htemp:
            return False
        else:
            self._ltemp=ltemp
            return True

    def setInterval(self,interval:int)->bool:
        if interval<=0:
            return False
        else:
            self._interval=interval
            return True


    def getPin(self)->int:
        return self._pin


    def getHighTemp(self)->int:
        return self._htemp


    def getLowTemp(self)->int:
        return self._ltemp


    def getInterval(self)->int:
        return self._interval


    def __str__(self):
        return "{0}: a cooling fan intelligent control class.".format(
                    self.__class__.__name__
                )


if __name__=="__main__":
    smartFan=SmartFan()
    smartFan.fanControl()
