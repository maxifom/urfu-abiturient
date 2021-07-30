import argparse
import os
import time
from functools import partial
from typing import Generator, Tuple

import numpy as np
import pandas as pd
import psycopg
import requests
import schedule
from scrapy import Selector

UA = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36"


def get_numbers() -> [int]:
    r = requests.get("https://urfu.ru/api/ratings/alphabetical/", headers={"User-Agent": UA})
    r.raise_for_status()
    s = Selector(text=r.text)
    numbers = list(map(int, s.css("span.letter.item a::attr(data-letter-num)").extract()))
    return numbers


def download_all() -> [pd.DataFrame]:
    dfs: [pd.DataFrame] = []
    number = get_numbers()
    for n in number:
        u = f"https://urfu.ru/api/ratings/alphabetical/27/{n}/"
        print(u)
        r = requests.get(u, headers={
            "User-Agent": UA
        })
        r.raise_for_status()

        u2 = "https://urfu.ru" + r.json()["url"]
        r2 = requests.get(u2, headers={"User-Agent": UA})
        r2.raise_for_status()
        html = r2.content
        if not r2.content:
            continue
        df = pd.read_html(html, encoding='utf-8')[0]
        df.columns = ["name", "number", "status", "type", "statement", "original", "speciality", "program", "form",
                      "budget", "vst", "vst1", "vst2", "vst3", "dost", "sum"]
        df = df.drop(['vst', 'vst1', 'vst2', 'vst3', 'dost'], axis=1)
        df = df[df["sum"] != "Выбыл из конкурса"]
        df = df[df["sum"] != "Забрал документы"]
        df['sum'] = df['sum'].replace(np.nan, -1)
        df['sum'] = df['sum'].astype('int32')
        dfs.append(df)
    return dfs


def concat_dfs(dfs: [pd.DataFrame]) -> pd.DataFrame:
    return pd.concat(dfs, ignore_index=True)


def generate_records(df: pd.DataFrame) -> Generator[Tuple, None, None]:
    for _, row in df.iterrows():
        yield (
            row["name"],
            int(row["number"]),
            row["status"],
            row["type"],
            bool(row["statement"] == "Да"),
            bool(row["original"] == "Да"),
            row["speciality"],
            row["program"],
            row["form"],
            row["budget"],
            int(row["sum"])
        )


def download():
    return concat_dfs(download_all())


def job(db_url: str):
    print("RUN JOB")
    df = download()
    with psycopg.connect(db_url) as conn:
        with conn.cursor() as cursor:
            cursor.execute("TRUNCATE TABLE abiturient_entries RESTART IDENTITY;")
            conn.commit()
            with cursor.copy(
                    "COPY abiturient_entries (name, number, status, type, statement_given, original_given, speciality, program, form, basis, sum) FROM STDIN") as copy:
                for r in generate_records(df):
                    copy.write_row(r)
            cursor.execute("""
            INSERT INTO last_updateds (id, last_updated)
VALUES (1, NOW())
ON CONFLICT (id) DO UPDATE SET last_updated = NOW()
""")
            conn.commit()
    print("FINISHED RUN JOB")


def main():
    argparser = argparse.ArgumentParser()
    argparser.add_argument("--db-url",
                           default=os.getenv("DB_URL", "postgres://postgres:postgres@localhost:5432/postgres"),
                           type=str)
    argparser.add_argument('--at-minutes', default=os.getenv("AT_MINUTES", ":00"), type=str)
    args = argparser.parse_args()
    print(args)
    schedule_job = partial(job, db_url=args.db_url)
    schedule.every().hour.at(args.at_minutes).do(schedule_job)
    print("JOB:", schedule.jobs[0])
    while True:
        schedule.run_pending()
        time.sleep(1)


if __name__ == '__main__':
    main()
