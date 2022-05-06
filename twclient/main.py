import time
import tweepy
import os
import grpcconn
from datetime import datetime, timedelta


def getClient():
    client = tweepy.Client(
        consumer_key=os.environ["consumer_key"],
        consumer_secret=os.environ["consumer_secret"],
        access_token=os.environ["access_token"],
        access_token_secret=os.environ["access_token_secret"],
    )
    return client


def makeTweetText(daytime, nighttime, total):
    today = datetime.now()
    yesterday = today - timedelta(1)
    yesterdayString = yesterday.strftime("%Y-%d-%m")
    text = ""
    text += "@azuki774s\n"
    text += yesterdayString + " の電力消費量は\n"
    text += "昼間:" + str(daytime) + " Wh\n"
    text += "夜間:" + str(nighttime) + " Wh\n"
    text += "合計:" + str(total) + " Wh\n"
    return text


if __name__ == "__main__":
    time.sleep(10)

    print("get target date")
    targetDate = grpcconn.getTargetDay()
    print(targetDate)

    conn = grpcconn.grpcClient()
    conn.open()
    print("grpc connected")
    res = conn.ElectConsumeGet()
    conn.close()

    print("get enviroment value")
    # print(os.environ["consumer_key"])
    # print(os.environ["consumer_secret"])
    # print(os.environ["access_token"])
    # print(os.environ["access_token_secret"])
    # print(os.getenv("twitter_stub"))

    print("make tweetText")
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
