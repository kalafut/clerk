from StringIO import StringIO
import textwrap
import importer


def test_import_csv():
    current = StringIO(textwrap.dedent('''\
            status,type,transaction_date,posting_date,description,amount
            R,,2016/11/02,,This is a test,$4.53
            '''))

    new = StringIO(textwrap.dedent('''\
            "Trans Date", "Summary", "Amount"
            5/2/2007, Regal Theaters, $15.99
            11/2/2016,  This is a test  ,  $4.53
            '''))

    mapping = {
        'Trans Date': 'transaction_date',
        'Summary': 'description',
        'Amount': 'amount'
    }

    importer.save_csv(current, new, mapping, '%m/%d/%Y')
    lines = current.getvalue().splitlines()

    assert lines[0].rstrip() == 'status,type,transaction_date,posting_date,description,amount'
    assert lines[1].rstrip() == 'N,,2007/05/02,,Regal Theaters,$15.99'
    assert lines[2].rstrip() == 'R,,2016/11/02,,This is a test,$4.53'
    assert len(lines) == 3



