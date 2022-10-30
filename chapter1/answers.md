Q:In the book DDIA, Martin (the author) states that Data Intensive systems have which 3 concerns?
=> B: Reliability, Scalability, Maintainability

Q:What are some examples of an application being reliable?
=> 1. The system returns the correct result.
=> 3. Under normal circumstances, the performance is good.
=> 4. The system can tolerate users' mistakes.
Reliable means the services can work as normal even if there are some failures.

Q:Which of the following are likely to be load parameters in a software system?
=> 1. requests per second to a web server
=> 3. the ratio of reads to writes in a database
=> 4. the hit rate on a cache

Q:In a batch processing system, we usually care about ____. In online systems, we usually care about the service's ____. 
=> D: throughput; response time

Q:It's best to think of response time as a single number.
=> False. We should think of the distribution.

Q:The average is not a very good metric for knowing the "typical" response time.
=> True. It is better to use percentiles.

Q:The median is also known as the 95th percentile.
=> False (50!)

Q:When a slow request holds up subsequent requests, we call this ____.
=> head-of-line blocking

Q:If a service is dependent on several backend calls, all backend calls must be fast in order for the service to be fast.
=> True

Q:Vertical scaling is when we add more servers on top of each other.
=> False. Vertical scaling means you replace your current working machine to a faster one!

Q:It is always better to have an elastic system.
=> False. Depends on your needs *(Non-elastic systems must be manually scaled but this may be easier to operate.)*

Q:Systems should be designed such that operations teams can easily keep the system running.
=> True

Q:True or False. What are some examples of things we can do to make a system have good operability.
=> (O) provide monitoring tools to provide visibility into the runtime behavior of the system
=> (O) provide good default behavior
=> (O) allow administrators the freedom to override defaults
=> (X) use test-driven development =>  for evolvability
=> (X) try to stick to the same individual machine

Q:Your system's requirements will probably never change, so it is not important to think about how to change the system.
=> False
