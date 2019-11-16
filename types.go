package main

type Project struct {
    Id string
}

type NewKey struct {
    Note string
    Project Project
}

type Response struct {
    StatusCode int
    Message string
    ExecutionTime int
    ResponseData interface{}
}

type TrafiklabKey struct {
    Key string
    Note string
    Api string
    Profile string
    Project string
    CreatedDate string
    UpdatedDate string
    Active bool
}

type TrafiklabKeys []struct {
    Key string
    Note string
    Api string
    Profile string
    Project string
    CreatedDate string
    UpdatedDate string
    Active bool
}

type RateLimit struct {
    Month int
    Minute int
}

type TrafiklabProfiles []struct {
    Id string
    Name string
    Api string
    RateLimit RateLimit
    Default bool
    CreatedDate string
    UpdatedDate string
}
