namespace Checkpoint.API.Responses;

public class ResponseDto
{
    public bool Success { get; set; }
    public int StatusCode { get; set; }
    public object? Data { get; set; }
    public List<string> Messages { get; set; }
    public ResponseDto(bool success, int statusCode, object? data, List<string> messages)
    {
        Success = success;
        StatusCode = statusCode;
        Data = data;
        Messages = messages;
    }
    public ResponseDto(bool success, int statusCode, object? data)
    {
        Success = success;
        StatusCode = statusCode;
        Data = data;
        Messages = [];
    }
    public static ResponseDto CreateSuccess(object data, int statusCode = 200)
    {
        return new ResponseDto(true, statusCode, data);
    }
    public static ResponseDto CreateError(List<string> messages, int statusCode = 400)
    {
        return new ResponseDto(false, statusCode, null, messages);
    }
}
