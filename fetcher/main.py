from curses import raw
import time
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

if __name__ == "__main__":
    print("fetcher start")
    time.sleep(10)
    conn = grpcconn.grpcClient()
    conn.open()
    print("grpc connected")
    conn.ElectConsumePost()
    conn.close()
    print("grpc closed")
    # print('Program start')
    # # Run Driver
    # driver = driver.get_driver()
    # driver.implicitly_wait(10)
    # print('Get driver')

    # print("------------------")

    # remix.login(driver)
    # remix.fetch_now_month(driver)

    # # Close browser
    # driver.quit()
    # print("------------------")

    # print('Program end')
