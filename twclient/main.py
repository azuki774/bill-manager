import time
import tweepy
import os
import grpcconn
from datetime import datetime, timedelta

wait_time = 0


def getClient():
    client = tweepy.Client(
        consumer_key=os.environ["consumer_key"],
        consumer_secret=os.environ["consumer_secret"],
        access_token=os.environ["access_token"],
        access_token_secret=os.environ["access_token_secret"],
    )
    return client


def makeTweetText(daytime, nighttime, total):
    today = datetime.now() + datetime.timedelta(hours=9)
    yesterday = today - timedelta(1)
    yesterdayString = yesterday.strftime("%Y-%m-%d")
    text = ""
    text += "@azuki774s\n"
    text += yesterdayString + " の電力消費量は\n"
    text += "昼間:" + str(daytime) + " Wh\n"
    text += "夜間:" + str(nighttime) + " Wh\n"
    text += "合計:" + str(total) + " Wh\n"
    return text


def get_start_time():
    if os.getenv("start_wait") == "":
        return 0
    else:
        return int(os.getenv("start_wait"))


if __name__ == "__main__":
    get_start_time()
    print("wait for " + str(wait_time) + "sec")
    time.sleep(wait_time)  # wait for other components

    print("get target date")
    targetDate = grpcconn.get_targetDay()
    print(targetDate)

    conn = grpcconn.grpcClient()
    conn.open()
    print("grpc connected")
    res = conn.ElectConsumeGet(targetDate)
    conn.close()

    if res.total == 0:
        print("fetch data error")
        exit(1)

    print("=== make tweetText ===")
    tweetText = makeTweetText(res.daytime, res.nighttime, res.total)
    print(tweetText)

    if os.getenv("twitter_stub") == "0":
        client = getClient()
        try:
            client.create_tweet(text=tweetText)
            print("tweet in Twitter")
        except Exception as e:
            print(e)
    else:
        print("tweet in stub")

    print("the program will end after 10 minutes")
    time.sleep(60 * 10)  # 10min sleep for blocking
    print("the program end")
