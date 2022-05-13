from curses import raw
import time
import os
import driver
from venv import create
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.chrome.service import Service
from bs4 import BeautifulSoup
import grpcconn
import remix

wait_time = 0


def get_start_time():
    if os.getenv("start_wait") == "":
        return 0
    else:
        return int(os.getenv("start_wait"))


if __name__ == "__main__":
    get_start_time()
    print("wait for " + str(wait_time) + "sec")
    time.sleep(wait_time)  # wait for other components
    print("fetcher start")

    remix_fetch_data = []
    if os.getenv("fetcher_stub") == "0":
        # Run Driver
        driver = driver.get_driver()
        driver.implicitly_wait(10)

        print("Get driver")

        remix.login(driver)
        remix_fetch_data = remix.fetch_now_month(driver)
    else:
        print("Use dummy remix")
        remix_fetch_data = remix.fetch_now_month_dummy()

    post_data = remix.make_postdata(remix_fetch_data)
    print(post_data)
    conn = grpcconn.grpcClient()
    conn.open()
    print("grpc connected")

    for data in post_data:
        conn.ElectConsumePost(data[0], data[1], data[2], data[3])

    conn.close()
    print("grpc closed")

    print("the program will end after 1 minutes")
    time.sleep(60)  # 1min sleep for blocking
    print("the program end")
