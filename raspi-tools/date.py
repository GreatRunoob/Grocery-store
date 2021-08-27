import time
from datetime import datetime as dt


class Date:
    @classmethod
    def getCurrentFormatDatetime(
            cls,fmt:str="%Y-%m-%d %H:%M:%S")->str:
        return dt.now().strftime(fmt)

    @classmethod
    def getCurrentFormatTime(
            cls,fmt:str="%Y-%m-%d %H:%M:%S")->str:
        return time.strftime(fmt)


if __name__=="__main__":
    # fmt="%Y/%m/%d-%H:%M:%S"
    print(Date.getCurrentFormatDatetime())
    print(Date.getCurrentFormatTime())
