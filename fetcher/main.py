from curses import raw
import re
import time
import json
import datetime
import os
import driver
from venv import create
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.chrome.service import Service
from bs4 import BeautifulSoup

import remix

if __name__ == '__main__':
    print('Program start')
    # Run Driver
    driver = driver.get_driver()
    driver.implicitly_wait(10)
    print('Get driver')
    
    print("------------------")

    remix.login(driver)
    remix.fetch_now_month(driver)

    # Close browser
    driver.quit()
    print("------------------")
    
    print('Program end')
