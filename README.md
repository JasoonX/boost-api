# BOOST-2021 Back v2

Code generation tool: [OpenAPITools](https://github.com/OpenAPITools/openapi-generator#overview).

Handlers and Requests keep naming convention:
- [add,get,update,delete]_[handler,request]
Add stands for POST;
All the rest match their HTTP methods.
#TODO:
- Refactor accroding to this spec:
- https://jsonapi.org/format

How to run locally:
# Running server with db locally
```
cd docker
docker-compose -f docker-compose.local.yaml up --build
```
(hit localhost:8080/v1/news or localhost:8080/healthcheck for test)

# Running docs preview locally
```
cd docker
docker-compose -f docker-compose.docs.yaml up
```
(hit localhost:9001 for swagger preview)