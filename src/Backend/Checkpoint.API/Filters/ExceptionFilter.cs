using System.Net;
using Checkpoint.API.Responses;
using Checkpoint.Exceptions.ExceptionBase;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;

namespace Checkpoint.API.Filters;

public class ExceptionFilter : IExceptionFilter
{
    private static readonly Dictionary<Type, Action<ExceptionContext>> ExceptionHandlers = new()
    {
        { typeof(ErrorOnValidationException), HandleValidationException },
        { typeof(NotFoundException), HandleNotFoundException },
        { typeof(UserAlreadyExistsException), HandleUserExistsException }
    };

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
        var exceptionType = context.Exception.GetType();

        if (ExceptionHandlers.TryGetValue(exceptionType, out var handler))
        {
            handler(context);
        }
    }

    private static void HandleValidationException(ExceptionContext context)
    {
        var exception = (ErrorOnValidationException)context.Exception;
        context.HttpContext.Response.StatusCode = (int)HttpStatusCode.BadRequest;
        context.Result = new BadRequestObjectResult(ResponseDto.CreateError(context));
    }

    private static void HandleNotFoundException(ExceptionContext context)
    {
        context.HttpContext.Response.StatusCode = (int)HttpStatusCode.NotFound;
        context.Result = new NotFoundObjectResult(ResponseDto.CreateError(context));
    }

    private static void HandleUserExistsException(ExceptionContext context)
    {
        context.HttpContext.Response.StatusCode = (int)HttpStatusCode.Conflict;
        context.Result = new ConflictObjectResult(ResponseDto.CreateError(context));
    }

    private static void HandleUnknownException(ExceptionContext context)
    {
        context.HttpContext.Response.StatusCode = (int)HttpStatusCode.InternalServerError;
        context.Result = new ObjectResult(ResponseDto.CreateError(["An unexpected error occurred"]));
    }
}
