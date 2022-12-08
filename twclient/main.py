import time
import tweepy
import os
import sys
import datetime
import mysql.connector


def getClient():
    client = tweepy.Client(
        consumer_key=os.environ["consumer_key"],
        consumer_secret=os.environ["consumer_secret"],
        access_token=os.environ["access_token"],
        access_token_secret=os.environ["access_token_secret"],
    )
    return client

def fetch_from_db(yyyymmdd):
    db_host=os.environ["db_host"]
    db_user=os.environ["db_user"]
    db_pass=os.environ["db_pass"]
    conn = mysql.connector.connect(host=db_host ,user=db_user, password=db_pass, database='billmanager', use_unicode=True)
    cur = conn.cursor(buffered=True)
    sql = 'SELECT * FROM elect_consumption WHERE record_date = "{}" LIMIT 1'.format(yyyymmdd)
    print(sql)
    cur.execute(sql)
 
    # 全てのデータを取得
    rows = cur.fetchall()
    cur.close()
    conn.close()
    return rows[0]

def makeTweetText(total, daytime, nighttime):
    today = datetime.datetime.now() + datetime.timedelta(hours=9)
    yesterday = today - datetime.timedelta(1)
    yesterdayString = yesterday.strftime("%Y-%m-%d")
    text = ""
    text += "@azuki774s\n"
    text += yesterdayString + " の電力消費量は\n"
    text += "合計:" + str(total) + " Wh\n"
    text += "昼間:" + str(daytime) + " Wh\n"
    text += "夜間:" + str(nighttime) + " Wh\n"
    return text

if __name__ == "__main__":
    today = datetime.datetime.now() + datetime.timedelta(hours=9)
    yesterday = today - datetime.timedelta(1)
    yyyymmdd = yesterday.strftime("%Y-%m-%d")
    res = fetch_from_db(yyyymmdd=yyyymmdd)

    print("make tweetText")
    tweetText = makeTweetText(res[2], res[3], res[4])
    print(tweetText)

    if os.getenv("only_print_debug") != "":
        print("the program debug end")
        sys.exit(0)

    client = getClient()
    try:
        client.create_tweet(text=tweetText)
        print("tweet in Twitter")
    except Exception as e:
        print(e)

    print("the program end")
