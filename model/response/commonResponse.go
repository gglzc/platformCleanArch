package response

//use to  general response

type CommonResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}


type ErrorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}