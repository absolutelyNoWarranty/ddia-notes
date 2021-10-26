# Chapter 1
Q:In the book DDIA, Martin (the author) states that Data Intensive systems have which 3 concerns?

A: Readability, Scalability, Maintainability
B: Reliability, Scalability, Maintainability
C: Speed, Correctness, Accuracy
D: Throughput, Performance, Reliability


Q:What are some examples of an application being reliable?
1. The system returns the correct result.
2. The system is easy to deploy on many servers.
3. Under normal circumstances, the performance is good.
4. The system can tolerate users' mistakes.

Q:Which of the following are likely to be load parameters in a software system?
1. requests per second to a web server
2. the number of CPUs in a server
3. the ratio of reads to writes in a database
4. the hit rate on a cache

Q:In a batch processing system, we usually care about ____. In online systems, we
usually care about the service's ____.
A: speed; throughput
B: response time; throughput
C: throughput; response time
D: latency; response time


Q:True or False. It's best to think of response time as a single number.

Q:True or False. The average is not a very good metric for knowing the "typical"
response time.

Q:True or False. The median is also known as the 95th percentile.

Q:When a slow request holds up subsequent requests, we call this ____.

Q:True or False. If a service is dependent on several backend calls, all backend calls
must be fast in order for the service to be fast.

Q:True or False. Vertical scaling is when we add more servers on top of each other.

Q:True or False. It is always better to have an elastic system.

Q:True or False. Systems should be designed such that operations teams can easily keep
the system running.

Q:True or False. What are some examples of things we can do to make a system have good operability.
1. provide monitoring tools to provide visibility into the runtime behavior of the system
2. provide good default behavior
3. allow administrators the freedom to override defaults
4. use test-driven development
5. try to stick to the same individual machine

Q:True or False. Your system's requirements will probably never change, so it is not important
to think about how to change the system.
