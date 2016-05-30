import collections
import csv
import hashlib
import os

from datetime import datetime

ImportTransaction = collections.namedtuple("ImportTransaction", "date summary amount")

IMPORTS_DIR = "imports"
ARCHIVE_DIR = 'archive'

fields_1 = [
    'status',  # N-New, A-Added, I-Ignored
    'qty',
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
        if f not in ['status', 'qty']:
            h.update(row.get(f, ""))
    return h.hexdigest()


def save_csv(log, additions, field_cfg, date_format):
    """Add any unique entries from new to current"""
    qtys = collections.Counter()
    pool = {}
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
        conv_row = {}
        for k, v in row.items():
            if k in field_cfg:
                conv_row[field_cfg[k]] = v.strip()

        conv_row['status'] = 'N'
        date = datetime.strptime(conv_row['transaction_date'], date_format).date()
        conv_row['transaction_date'] = date.strftime('%Y/%m/%d')

        hash_ = hash_row(conv_row)
        if hash_ in log_hashes:
            continue
        pool[hash_] = conv_row
        qtys[hash_] += 1

    for key, val in qtys.items():
        row = pool[key]
        if val > 1:
            row['qty'] = val
        log_rows.append(row)

    log_rows.sort(key=lambda row: (row['transaction_date'], row['description']))

    log.seek(0)
    writer = csv.DictWriter(log, fields_1, lineterminator='\n')
    writer.writeheader()
    for row in log_rows:
        writer.writerow(row)
    log.truncate()

def save(archive_name, input_):
    if not os.path.exists(ARCHIVE_DIR):
        os.mkdir(ARCHIVE_DIR)
    with open(os.path.join(ARCHIVE_DIR, archive_name), 'r+') as archive:
        mapping = {
            'Trans Date': 'transaction_date',
            'Description': 'description',
            'Amount': 'amount'
        }

        save_csv(archive, input_, mapping, '%m/%d/%Y')

