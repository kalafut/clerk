import block
import importer
import Levenshtein


def match(query, pool):
    best_score = -1
    best_block = None
    for block in pool:
        #print query, block.summary
        score = Levenshtein.ratio(query, block.summary)
        if score > best_score:
            best_block = block
            best_score = score
            if "22478" in query:
                pass
                #print score

    return best_block, best_score


with open("chase_freedom.ldg") as f:
    pool = block.parse(f)

with open("chase_sapphire.ldg") as f:
    pool2 = block.parse(f)

pool = pool2

with open("activity.csv") as f:
    for query in importer.import_csv(f):
        block, score = match(query.summary, pool)
        print query.summary
        print block, score
        print
