# evelib/redisq

This package implements a client to [ZKillboard's RedisQ API](http://redisq.zkillboard.com/listen.php) for retrieving EVE Online killmails in near-real time as they are posted to ZKillboard.

## **PLEASE NOTE**

*Due to he nature of RedisQ, having more than one RedisQ client running simultaneously (in the same program or another) will ultimately disrupt one another.* This means that for ever killmail consumed by one client, another client will **not** receive the same killmail.

*Avoiding this issue:* You can specify a Queue ID in `redisq.Options` that uniquely identifies your RedisQ client.
