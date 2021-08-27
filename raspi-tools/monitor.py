# -*- coding=utf-8 -*-
import os
from typing import Generator


class PiMonitor:

    @classmethod
    def getCPUTemp(cls)->float:
        try:
            with open("/sys/class/thermal/thermal_zone0/temp") as temp_file:
                temp=temp_file.readline()
            return float(temp)/1e3  # e.g. temp='23849' -> 23.849
        except (ValueError,FileNotFoundError,OSError) as error:
            raise error


    @classmethod
    def getCPUFreq(cls)->float:
        try:
            freq=os.popen(
                    "vcgencmd measure_clock arm | awk -F '=' '{print $2}'"
                    ).readline()
            return float(freq)/1e9  # e.g. freq='600117184' -> 0.600117184
        except (ValueError,OSError) as error:
            raise error


    @classmethod
    def getCPUUsage(cls)->float:
        try:
            free=os.popen(
                    "top -n1 | awk '/Cpu\(s\):/ {print $8}'"
                    ).readline().strip()
            return 1e2-float(free)  # e.g. free='96.5' -> 3.5
        except (ValueError,OSError) as error:
            raise error


    @classmethod
    def getCPULoads(cls)->Generator[str,float,None]:
        try:
            loads=os.popen("uptime").readline()
            # e.g. loads="0.04 0.04 0.04"
            gen_loads=(
                    float(load.strip(',')) for load in loads.split()[-3:]
                    )
            return gen_loads
        except (ValueError,OSError) as error:
            raise error


    @classmethod
    def getMemInfos(cls)->Generator[str,float,None]:
        try:
            infos=os.popen(
                    "free | awk '/Mem:/ {print $2,$3,$4}'"
                    ).readline()
            # e.g. infos="1859364 279256 262252"
            gen_infos=(float(info)/1024 for info in infos.split())
            return gen_infos
        except (ValueError,OSError) as error:
            raise error


    @classmethod
    def getSwapInfos(cls)->Generator[str,float,None]:
        try:
            infos=os.popen(
                    "free | tail -n 1 | awk '{print $(NF-2),$(NF-1),$NF}'"
                    ).readline()
            # e.g. infos="102396 11776 90620"
            gen_infos=(
                    float(info)/1024 for info in infos.split()
                    )
            return gen_infos
        except (ValueError,OSError) as error:
            raise error


    @classmethod
    def getSDCardInfos(cls)->Generator[str,float,None]:
        try:
            infos=os.popen(
                    "df / | tail -n 1 | awk '{print $2,$3,$4}'"
                    ).readline()
            # e.g. infos="30353628 7267516 21753529"
            gen_infos=(
                    float(info)/1024**2 for info in infos.split()
                    )
            return gen_infos
        except (ValueError,OSError) as error:
            raise error


def print_CPU_temp():
    try:
        print(
                "CPU temperature: {0:.2f} 'C".format(
                    PiMonitor.getCPUTemp()
                    )
                )
    except Exception as e:
        print(e)

    return


def print_CPU_freq():
    try:
        print(
                "CPU frequency: {0:.2f} GHz".format(
                    PiMonitor.getCPUFreq()
                    )
                )
    except Exception as e:
        print(e)

    return


def print_CPU_usage():
    try:
        print(
                "CPU usage: {0:.2f} %".format(
                    PiMonitor.getCPUUsage()
                    )
                )
    except Exception as e:
        print(e)

    return


def print_CPU_loads():
    try:
        gen_loads=PiMonitor.getCPULoads()
        print("Loads:\t1min\t5min\t15min")
        for load in gen_loads:
            print("\t{}".format(load),end='')
        print("")
    except Exception as e:
        print(e)

    return


def print_mem_infos():
    try:
        total,used,free=PiMonitor.getMemInfos()
        print("Mem:\tTotal\tUsed\tFree\tUsage")
        print(
                "\t{0:.1f}M\t{1:.1f}M\t{2:.1f}M\t{3:.2%}".format(
                    total,used,free,used/total
                    )
                )
    except Exception as e:
        print(e)

    return


def print_swap_infos():
    try:
        total,used,free=PiMonitor.getSwapInfos()
        print("Swap:\tTotal\tUsed\tFree\tUsage")
        print(
                "\t{0:.1f}M\t{1:.1f}M\t{2:.1f}M\t{3:.2%}".format(
                    total,used,free,used/total
                    )
                )
    except Exception as e:
        print(e)

    return


def print_sdcard_infos():
    try:
        total,used,free=PiMonitor.getSDCardInfos()
        print("SDCard:\tTotal\tUsed\tFree\tUsage")
        print(
                "\t{0:.2f}G\t{1:.2f}G\t{2:.2f}G\t{3:.2%}".format(
                    total,used,free,used/total
                    )
                )
    except Exception as e:
        print(e)

    return


if __name__=="__main__":
    split_line="="*40

    print(split_line)
    # Print current statement of CPU
    print_CPU_temp()
    print_CPU_freq()
    print_CPU_usage()
    print_CPU_loads()
    print(split_line)
    # Print current statement of memory
    print_mem_infos()
    print(split_line)
    # Print current statement of swapspace
    print_swap_infos()
    print(split_line)
    # Print current statement of sdcard
    print_sdcard_infos()
    print(split_line)

