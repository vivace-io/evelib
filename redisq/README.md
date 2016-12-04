# evelib/redisq

This package implements a client to ZKillboard for retrieving killmails in near-real time from the service's RedisQ endpoint(http://redisq.zkillboard.com/listen.php). **It is still largely untested and considered experimental.**

## **PLEASE NOTE:**
*Due to he nature of RedisQ, having more than one RedisQ client running simultaneously (in the same program or another) will ultimately disrupt one another.* This means that for ever killmail consumed by one client, another client will **not** receive the same killmail.

*To avoid this issue:* consider routing the different applications through separate external IP addresses
