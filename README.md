# how to start locally
```bash
# spin up traefik at localhost:8080
docker-compose -f traefik.yml up -d 

# spin up helloserver with auth at helloserver.localhost
docker-compose -f whoami.yml up -d

# spin up authorization server at auth.localhost
docker-compose -f auth.yml up -d

# expect 200 with htlm content echo-ing request
curl whoami.localhost
```