import time
import tweepy
import os
import grpcconn
import datetime

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
    time.sleep(1000)  # FOR TEST
    today = datetime.datetime.now() + datetime.timedelta(hours=9)
    yesterday = today - datetime.timedelta(1)
    yesterdayString = yesterday.strftime("%Y-%m-%d")
    text = ""
    text += "@azuki774s\n"
    text += yesterdayString + " の電力消費量は\n"
    text += "昼間:" + str(daytime) + " Wh\n"
    text += "夜間:" + str(nighttime) + " Wh\n"
    text += "合計:" + str(total) + " Wh\n"
    return text


def get_start_time():
    return int(os.environ.get("start_wait", "0"))


if __name__ == "__main__":
    wait_time = get_start_time()
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
