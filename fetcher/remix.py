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

def login(driver):
    driver.get("https://remix-denki.com/Account/Login")

    login_id = driver.find_element(by=By.XPATH, value='//*[@id="LoginId"]')
     
    login_id.send_keys("")
    
    password_form = driver.find_element(by=By.XPATH, value='//*[@id="Password"]')
    password_form.send_keys("")

    login_button = driver.find_element(by=By.XPATH, value='//*[@id="LoginSubmit"]')
    login_button.click()

    return

def fetch_now_month(driver):
    driver.get("https://remix-denki.com/Demand")
    month_button = driver.find_element(by=By.XPATH, value='//*[@id="month"]')
    month_button.click()

    daytime_month_consume=driver.find_element(by=By.XPATH, value='//*[@id="daytime_total"]')
    night_month_consume = driver.find_element(by=By.XPATH, value='//*[@id="night_total"]')
    total_month_consume = driver.find_element(by=By.XPATH, value='//*[@id="total"]')
    print(daytime_month_consume.get_attribute("textContent"))
    print(night_month_consume.get_attribute("textContent"))
    print(total_month_consume.get_attribute("textContent"))
