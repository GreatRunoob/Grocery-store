import time

from monitor import PiMonitor


if __name__=="__main__":
    try:
        while True:
            print("Current CPU temp: {0:.2f}'C".format(
                    PiMonitor.getCPUTemp()
                )
            )

            time.sleep(3)
    except KeyboardInterrupt:
        print("Exit!")
