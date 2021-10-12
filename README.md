# KrakenD for GoLang Servers

A test of KrakenD aggregator capabilities.

## Usage

```bash
# Validate the krakend configuration
krakend check --config krakend.json --test-gin-routes --debug

# Start the backend servers
go run .

# Start the api gateway server
krakend run --config krakend.json --debug
```

## Frontend for Backend

```md
curl -i http://localhost:8080/frontpage

#=> {
#=>   "comments":[
#=>     {"id":"1","message":"First comment!!!","postId":"1"},
#=>     {"id":"2","message":"Beat me by 1ns.","postId":"1"}
#=>   ],
#=>   "posts":[
#=>     {"id":"1","title":"Lorem Ipsum1"},
#=>     {"id":"2","title":"Lorem Ipsum2"}
#=>   ]
#=> }
```

## License

This project is __FREE__ to use, reuse, remix, and resell.
This is made possible by the [MIT license](/LICENSE).
