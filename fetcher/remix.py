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
from datetime import datetime, timedelta
import datetime
import grpc
from proto.api_pb2 import *
from proto import api_pb2_grpc


def login(driver):
    driver.get("https://remix-denki.com/Account/Login")

    login_id = driver.find_element(by=By.XPATH, value='//*[@id="LoginId"]')

    login_id.send_keys(os.getenv("remix_id"))

    password_form = driver.find_element(by=By.XPATH, value='//*[@id="Password"]')
    password_form.send_keys(os.getenv("remix_pass"))

    login_button = driver.find_element(by=By.XPATH, value='//*[@id="LoginSubmit"]')
    login_button.click()
    driver.implicitly_wait(10)
    return


def fetch_now_month(driver):
    driver.get("https://remix-denki.com/Demand")
    month_button = driver.find_element(
        by=By.XPATH, value="/html/body/div[2]/div/div/div[3]/div[2]"
    )
    month_button.click()

    result_data = []  # [day, daytime_consume, nighttime_consume]

    try:
        for i in range(1, 40):
            e1text = (
                "/html/body/div[2]/div/div/div[5]/div[2]/div[2]/div[2]/div["
                + str(i)
                + "]/div[1]"
            )
            e2text = (
                "/html/body/div[2]/div/div/div[5]/div[2]/div[2]/div[2]/div["
                + str(i)
                + "]/div[2]/span"
            )
            e3text = (
                "/html/body/div[2]/div/div/div[5]/div[2]/div[2]/div[2]/div["
                + str(i)
                + "]/div[3]/span"
            )
            e1 = driver.find_element(
                by=By.XPATH,
                value=e1text,
            )
            e2 = driver.find_element(
                by=By.XPATH,
                value=e2text,
            )
            e3 = driver.find_element(
                by=By.XPATH,
                value=e3text,
            )
            result_data.append(
                [
                    int(e1.get_attribute("textContent")),
                    float(e2.get_attribute("textContent")),
                    float(e3.get_attribute("textContent")),
                ]
            )
    except Exception as e:
        print("fetch end")
    finally:
        print(result_data)
        return result_data


def fetch_now_month_dummy():
    result_data = []
    ds = get_targetDay()
    count = ds.day

    for num in range(count):
        result_data.append([(num + 1), 1.0 * (num + 1), 1.5 * (num + 1)])

    print(result_data)
    return result_data


def get_targetDay():
    # ?????????????????????????????????????????????
    nowadays = datetime.datetime.now() + datetime.timedelta(hours=9)
    yesterday = nowadays - datetime.timedelta(1)
    return DateStruct(year=yesterday.year, month=yesterday.month, day=yesterday.day)


def make_postdata(fetch_data):
    ret_data = []

    for data in fetch_data:
        targetday = get_targetDay()
        # [1, 2.0, 3.0] -> [DataStruct(year=2022,month=1,day=1), 2000, 3000, 5000] kWh -> Wh
        targetday.day = data[0]  # ?????????????????????????????????????????????
        if data[1] + data[2] == 0.0:
            break
        ret_data.append(
            [
                targetday,
                int(round(data[1] * 1000)),
                int(round(data[2] * 1000)),
                int(round(data[1] * 1000)) + int(round(data[2] * 1000)),
            ]
        )

    return ret_data
