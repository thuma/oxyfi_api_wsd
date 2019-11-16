package main

var default_profiles = TrafiklabProfiles {
    {
    Id: "One",
    Name: "Standard",
    Api: "oxygps",
    RateLimit: RateLimit {
        Month: 24000,
        Minute: 6,
        },
    Default: true,
    CreatedDate: "2017-01-01T15:00:00.000Z",
    UpdatedDate: "2017-01-01T15:00:00.000Z",
    },
}

var key_example = TrafiklabKey{
    Key:"FFFF",
    Note:"",
    Api:"oxygps",
    Profile:"One",
    Project:"saiads982",
    CreatedDate:"2019-01-01T12:12:12.0000Z",
    UpdatedDate:"2019-01-01T12:12:12.0000Z",
    Active:true,
}

var keys_example = TrafiklabKeys{
    {
        Key:"FFFF",
        Note:"",
        Api:"oxygps",
        Profile:"One",
        Project:"saiads982",
        CreatedDate:"2019-01-01T12:12:12.000Z",
        UpdatedDate:"2019-01-01T12:12:12.000Z",
        Active:true,
    },
    {
        Key:"FFFF",
        Note:"",
        Api:"oxygps",
        Profile:"One",
        Project:"saiads982",
        CreatedDate:"2019-01-01T12:12:12.000Z",
        UpdatedDate:"2019-01-01T12:12:12.000Z",
        Active:true,
    },
    }

var default_response = Response{
    StatusCode : 200,
    Message: "",
    ExecutionTime: 72,
    ResponseData: nil,
    }
