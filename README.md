# Clerk
Clerk is is [Ledger](http://www.ledger-cli.org/)-inspired personal finance application.

## Philosophy
Storing financial data in text form is great; editing it by hand is not. My financial data consists of a huge backlog of existing transcations, and a steady, high volume of incoming transactions. And tool that I want to use regularly must be compatible with that scenario. The impact this has on the storage as compared to Ledger:

- **Single-line transactions** -- Ledger files may look nice, but they don't lend themselves to easy text processing. Clerk's files are CSV, and all data for a given transaction is one line. Suddenly a lot more querying and processing and be done even with unix tools. The files are still hand-editable, though they're quite a bit more dense. But if you're regularly editing by hand, then something is wrong.
- **Strict format** -- Ledger files are quite free form, Clerk's are not. The arrangement of transactions and fields is explicitly defined, and operations will always export in this form. Parsers are somewhat tolerant, but the output will always be forced by to a normalized form.
- **Mutable files** -- A key principle of Ledger it is a reporting tool. Clerk has operations that operate directly on data files. Since said files are already in source control (right?), this only simplifies working with them. Even better, Clerk provides a sort of mini-reflog so you do need to constantly lean on your SCM and you mutate your files this way and that. Run experiments, break things, undo things, fixing thing up, and then run `git commit`.

# Amount inference
All transactions must balance. If 'Food' increases by $5, there must be $5 worth of decreases elsewhere.

### Commodity handling in Ledger

Ledger will examine the first use of any commodity to determine how that commodity should be printed on reports. It pays attention to whether the name of commodity was separated from the amount, whether it came before or after, the precision used in specifying the amount, whether thousand marks were used, etc. This is done so that printing the commodity looks the same as the way you use it.
