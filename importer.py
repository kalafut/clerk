import collections
import hashlib
from datetime import datetime
import csv

ImportTransaction = collections.namedtuple("ImportTransaction", "date summary amount")

IMPORTS_DIR = "imports"

fields_1 = [
    'status',  # N-New, A-Added, I-Ignored
    'type',
    'transaction_date',
    'posting_date',
    'description',
    'amount'
    ]

#StdImport = collections.namedtuple('StdImport', fields_1)

def hash_row(row):
    h = hashlib.sha1()
    for f in fields_1:
        if f != 'status':
            h.update(row.get(f, ""))
    return h.hexdigest()


def save_csv(log, additions, field_cfg, date_format):
    """Add any unique entries from new to current"""
    log_rows = []
    log_hashes = set()
    for row in csv.DictReader(log):
        log_rows.append(row)
        hash_ = hash_row(row)
        assert hash_ not in log_hashes
        log_hashes.add(hash_)

    dialect = csv.Sniffer().sniff(additions.read(1024))
    additions.seek(0)
    for row in csv.DictReader(additions, dialect=dialect):
        conv_row = {field_cfg[k]: v.strip() for k, v in row.items()}

        conv_row['status'] = 'N'
        date = datetime.strptime(conv_row['transaction_date'], date_format).date()
        conv_row['transaction_date'] = date.strftime('%Y/%m/%d')

        hash_ = hash_row(conv_row)
        if hash_ in log_hashes:
            continue
        log_rows.append(conv_row)

    log_rows.sort(key=lambda row: row['transaction_date'])

    log.seek(0)
    writer = csv.DictWriter(log, fields_1)
    writer.writeheader()
    for row in log_rows:
        writer.writerow(row)
    log.truncate()

