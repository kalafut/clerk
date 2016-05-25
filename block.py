import collections
import re

from datetime import datetime

"""
State machine concept...

States:
	BeforeBlock
	LeadingComments
	Summary
	Postings

Line classifications:
	Empty
	GlobalComment
	Summary
	TransactionComment
	Posting
"""

BLANK = 0
COMMENT = 1
SUMMARY = 2
POSTING = 3
TXNCOMMENT = 4
INVALID = 5

COMMENT_CHARS = ";#|*%"
BLANK_CHARS = " \t"

reBlank = re.compile(r'^\s*$')
reComment = re.compile(r'^[;#|\*%].*$')
reSummary = re.compile(r'^(?P<date>\d{4}/\d\d/\d\d)(?: +(?P<cleared>[!\*]))?(?: +\((?P<code>.*?)\))? +(?P<summary>.*?) *$')
rePosting = re.compile(r'^\s+(?P<account>[^;#|\*%].*?)(?:\s{2,}(?P<amount>.*))?$')
reTxnComment = re.compile(r'^\s+[;#|\*%].*$')

"""
Blocks are a literal encapsulation of a Ledger transaction. They
are not called transcactions because the actual ledger file strings
and comments are preserved. A ledger file is a sequence of blocks.

Textually, a block is defined as:
   <0+ comment lines>
   <0 or 1 summary line: a) left justified  b) starting with a yyyy/mm/dd date
   <0+ acccount lines or comments: a) indented at least one space>

Whitespace between blocks is ignored.
"""

Posting = collections.namedtuple("Posting", "account amount")

class Block:
    def __init__(self):
        self.lines = []
        self.postings = []
        self.date = datetime.now()
        self.valid = False
        self.summary = ""
        self.cleared = None

    def __repr__(self):
        return repr(self.lines)


def st_before_block(line, block):
    assert block is None

    if reBlank.match(line) or reComment.match(line):
        return None, st_before_block

    if reSummary.match(line):
        match = reSummary.match(line)
        block = Block()
        matches = match.groupdict()
        block.date = datetime.strptime(matches["date"], "%Y/%m/%d").date()
        block.cleared = matches["cleared"]
        block.summary = matches["summary"]
        block.lines.append(line)

        return block, st_in_block

    raise Exception(line)

def st_in_block(line, block):
    assert block is not None

    if rePosting.match(line):
        match = rePosting.match(line)
        posting = match.groupdict()
        block.postings.append(Posting(posting["account"], posting["amount"]))

        block.lines.append(line)
        return block, st_in_block

    if reComment.match(line):
        return block, st_in_block

    if reBlank.match(line):
        return None, st_before_block

    raise Exception(line)


def parse(f):
    blocks = []
    state = st_before_block

    block = None
    for line in f:
        next_block, state = state(line.rstrip(), block)
        if block is not next_block:
            block = next_block
            if next_block is not None:
                blocks.append(block)

    return blocks

#with open("sample.dat") as f:
#    blocks = parse(f)
#    print blocks
#
#exit()

"""
	BeforeBlock
	LeadingComments
	Summary
	Postings
        """
r"""

// Note: values will not have '-', intentionally
var acctAmtRegex = regexp.MustCompile(`^\s+(.*?\S)(?:\s{2,}.*?([\d,\.]+))?\s*$`)


// ParseLines turns a chunk of text into a group of Blocks.
func ParseLines(data io.Reader) []Block {
	const (
		stBeforeBlock = iota
		stLeadingComments
		stSummary
		stPostings
	)
	var block Block
	var blocks []Block
	var state = stBeforeBlock

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		//_ = "breakpoint"
		line := scanner.Text()
		switch state {
		//case stBeforeBlock:
		//	if len(strings.TrimSpace(line)) > 0 {
		//		block.lines = append(block.lines, line)
		//	}
		}

		if len(strings.TrimSpace(line)) == 0 {
			if !block.Empty() {
				blocks = append(blocks, block)
				block = Block{}
			}
		} else {
			t, err := time.Parse("2006/01/02", line[0:10])
			if err == nil {
				// Start a new block
				if !block.Empty() {
					blocks = append(blocks, block)
					block = Block{}
				}
				block.date = t
			}
			block.lines = append(block.lines, line)
		}
	}

	if !block.Empty() {
		blocks = append(blocks, block)
	}

	return blocks
}
func (b *Block) Empty() bool {
	return len(b.lines) == 0
}

func (b Block) Accounts() []string {
	var ret []string
	for _, l := range b.lines {
		m := acctAmtRegex.FindStringSubmatch(l)
		if len(m) > 0 {
			ret = append(ret, m[1])
		}
	}
	sort.Strings(ret)
	return ret
}

func (b Block) Amounts() []string {
	var ret []string
	for _, l := range b.lines {
		m := acctAmtRegex.FindStringSubmatch(l)
		if len(m) > 0 {
			ret = append(ret, m[2])
		}
	}
	sort.Strings(ret)
	return ret
}

// IsDupe returns true if other is a likely duplicate based on:
//   date
//   affected accounts
//   amounts
func (b Block) IsDupe(other Block, tolerance time.Duration) bool {
	// Check time
	timeDiff := b.date.Sub(other.date)
	if timeDiff < 0 {
		timeDiff = -timeDiff
	}

	if timeDiff > tolerance {
		return false
	}

	// Check affected accounts
	accts := b.Accounts()
	acctsOther := other.Accounts()

	if len(accts) != len(acctsOther) {
		return false
	}

	for i := range accts {
		if accts[i] != acctsOther[i] {
			return false
		}
	}

	// Check affected accounts
	amts := b.Amounts()
	amtsOther := other.Amounts()

	if len(amts) != len(amtsOther) {
		return false
	}

	for i := range amts {
		if amts[i] != amtsOther[i] {
			return false
		}
	}

	return true
}

// prepareBlock processes []lines data, checking for errors and
// populating internal fields
func (b *Block) prepareBlock() {
	b.valid = true
}

func classifyLine(line string) (int, map[string]string) {
	var cls = clsInvalid
	var data map[string]string
	var captures []string
	var matchingRe *regexp.Regexp

	if reBlank.MatchString(line) {
		cls = clsBlank
	} else if reComment.MatchString(line) {
		cls = clsComment
	} else if rePosting.MatchString(line) {
		cls = clsPosting
	} else if reTxnComment.MatchString(line) {
		cls = clsTxnComment
	} else if captures = reSummary.FindStringSubmatch(line); len(captures) > 0 {
		cls = clsSummary
		matchingRe = reSummary
	}

	if captures != nil {
		data = make(map[string]string)
		for i, key := range matchingRe.SubexpNames() {
			if i > 0 {
				data[key] = captures[i]
			}
		}
	}

	return cls, data
}

// FindDupes returns a list of likely duplicate blocks. Duplicates
// are block with the same date and transaction structure. The same
// accounts and amounts must be present in both for it to be dupe.
func FindDupes(ledger Ledger) {
	blocks := ledger.blocks
	for i := range blocks {
		for j := i + 1; j < len(blocks); j++ {
			if blocks[i].IsDupe(blocks[j], 0) {
				fmt.Printf("%v,%v:%v\n", i, j, blocks[i].lines[0])
			}
		}
	}
}

func NewBlock(t transaction, config AccountConfig) Block {
	lines := fmt.Sprintf("%s   %s\n", t.date, t.description)
	lines += fmt.Sprintf("    %s          %s\n", importAcct, t.amount)
	lines += fmt.Sprintf("    %s", config.TargetAccount)

	blocks := ParseLines(strings.NewReader(lines))
	if len(blocks) != 1 {
		log.Fatalf("Expected 1 block, got %+v", blocks)
	}

	return blocks[0]
}
*/
"""
