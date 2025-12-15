# how to start locally
```bash
# spin up traefik at localhost:8080
docker-compose -f traefik.yml up -d 

# spin up helloserver with auth at helloserver.localhost
docker-compose -f helloserver.yml up -d

# expect 200 with htlm content echo-ing request
curl helloserver.localhost
```