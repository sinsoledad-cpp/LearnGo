namespace go api

struct Request {
    1: string message
}

struct Response {
    1: string message
}

service Echo {
    Response Echo(1: Request req)
}


// cwgo server -I ../../idl --type RPC --module github.com/CloudWeGo/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift