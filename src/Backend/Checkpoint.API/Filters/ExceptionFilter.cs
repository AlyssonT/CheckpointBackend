using System.Net;
using Checkpoint.API.Responses;
using Checkpoint.Exceptions.ExceptionBase;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;

namespace Checkpoint.API.Filters;

public class ExceptionFilter : IExceptionFilter
{
    public void OnException(ExceptionContext context)
    {
        if (context.Exception is CheckpointException)
        {
            HandleCheckpointException(context);
        }
        else
        {
            HandleUnknownException(context);
        }
    }
    private static void HandleCheckpointException(ExceptionContext context)
    {
        if (context.Exception is ErrorOnValidationException exception) {
            context.HttpContext.Response.StatusCode = (int)HttpStatusCode.BadRequest;
            context.Result = new BadRequestObjectResult(ResponseDto.CreateError(exception.ErrorMessages.ToList()));
        }
        else if (context.Exception is NotFoundException)
        {
            context.HttpContext.Response.StatusCode = (int)HttpStatusCode.NotFound;
            context.Result = new NotFoundObjectResult(ResponseDto.CreateError([context.Exception.Message]));
        }
    }

    private static void HandleUnknownException(ExceptionContext context)
    {
        context.HttpContext.Response.StatusCode = (int)HttpStatusCode.InternalServerError;
        context.Result = new ObjectResult(ResponseDto.CreateError(["An error occurred"]));
    }
}
