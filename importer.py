import collections
from datetime import datetime
import csv

ImportTransaction = collections.namedtuple("ImportTransaction", "date summary amount")


def import_csv(file_):
    reader = csv.DictReader(file_)
    for row in reader:
        date_str = row["Trans Date"]
        date = datetime.strptime(date_str, "%m/%d/%Y").date()
        summary = row["Description"].strip()
        amount = row["Amount"].strip()
        yield ImportTransaction(date, summary, amount)
