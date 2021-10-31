Q: Many databases use a log internally, which is an ____ data file.
A. read-only
B. write-only
C. append-only
D. create-only

Q: An ____ is a type of data structure for efficiently finding a particular key in a database. It is a ____ structure derived from ____.
A. dictionary; additional; primary data
B. index; additional; primary data
C. index; primary; additional data
D. log; additional; primary data

Q: Storage engines fall into which two broad categories?
1.OLPP
2.OLTP
3.OLAP
4.OLOP

Q: True or False. OLTP systems are typically user-facing.

Q: The bottleneck for OLTP systems is usually ____. The bottleneck for OLAP is usually ____.
A. disk seek time; disk bandwidth
B. disk seek time; disk seek time
C. disk bandwidth; disk bandwidth
D. disk bandwidth; disk seek time

Q: True or False. Both OLTP and OLAP systems typically must serve a huge volume of requsets.

Exercise: Implement column compression.
See bitmap_encoding.go/bitmap_encoding_test.
