import glob
import sys
import click
import Levenshtein
import importer
import block

def match(query, pool):
    query = query.lower()
    best_score = -1
    best_block = None
    for block in pool:
        score = Levenshtein.ratio(query, block.summary.lower())
        if score > best_score:
            best_block = block
            best_score = score

    return best_block, best_score


def test(import_file):
    ldg_files = glob.glob("*.ldg")
    corpus = []
    for file_ in ldg_files:
        with open(file_) as f:
            corpus.extend(block.parse(f))

    for query in importer.import_csv(import_file):
        ablock, score = match(query.summary, corpus)
        print query.date, query.summary
        print ablock, score
        print


@click.group()
def cli():
    pass


@cli.command()
@click.argument('target', type=click.File('r'))
def format(target):
    for t in block.parse(target):
        t.write(sys.stdout)


@cli.command(name="import")
@click.argument('import_file', type=click.File('r'), default='-')
@click.option('--target')
def import_(import_file, target):
    test(import_file)

if __name__ == "__main__":
    cli()
